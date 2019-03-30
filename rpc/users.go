package rpc

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"

	assets "github.com/YaleOpenLab/openx/assets"
	consts "github.com/YaleOpenLab/openx/consts"
	database "github.com/YaleOpenLab/openx/database"
	ipfs "github.com/YaleOpenLab/openx/ipfs"
	notif "github.com/YaleOpenLab/openx/notif"
	opensolar "github.com/YaleOpenLab/openx/platforms/opensolar"
	recovery "github.com/YaleOpenLab/openx/recovery"
	utils "github.com/YaleOpenLab/openx/utils"
	wallet "github.com/YaleOpenLab/openx/wallet"
	xlm "github.com/YaleOpenLab/openx/xlm"
)

func setupUserRpcs() {
	registerUser()
	ValidateUser()
	getBalances()
	getXLMBalance()
	getAssetBalance()
	getIpfsHash()
	authKyc()
	sendXLM()
	notKycView()
	kycView()
	askForCoins()
	trustAsset()
	uploadFile()
	platformEmail()
	sendTellerShutdownEmail()
	sendTellerFailedPaybackEmail()
	tellerPing()
	increaseTrustLimit()
	addContractHash()
	sendSecrets()
	mergeSecrets()
	generateNewSecrets()
	generateResetPwdCode()
	resetPassword()
	sweepFunds()
	sweepAsset()
}

const (
	// TellerUrl defines the teller URL to check. In future, would be an array
	TellerUrl = "https://localhost"
)

// we want to pass to the caller whether the user is a recipient or an investor.
// For this, we have an additional param called Role which we can use to classify
// this information and return to the caller

// ValidateParams is a struct used fro validating user params
type ValidateParams struct {
	Role   string
	Entity interface{}
}

// removeSeedRecp removes the encrypted seed from the recipient structure
func removeSeedRecp(recipient database.Recipient) database.Recipient {
	// any field that is private needs to be set to null here. A person using the API
	// knows the username and password anyway, so the route must return all routes
	// that are accessible by a single login (uname + pwhash)
	var dummy []byte
	recipient.U.EncryptedSeed = dummy
	return recipient
}

// removeSeedInv removes the encrypted seed from the investor structure
func removeSeedInv(investor database.Investor) database.Investor {
	var dummy []byte
	investor.U.EncryptedSeed = dummy
	return investor
}

// removeSeedEntity removes the encrypted seed from the entity structure
func removeSeedEntity(entity opensolar.Entity) opensolar.Entity {
	var dummy []byte
	entity.U.EncryptedSeed = dummy
	return entity
}

// UserValidateHelper is a helper that validates a user on the platform
func UserValidateHelper(w http.ResponseWriter, r *http.Request) (database.User, error) {
	var prepUser database.User
	var err error
	// need to pass the pwhash param here
	if r.URL.Query() == nil || r.URL.Query()["username"] == nil || r.URL.Query()["pwhash"] == nil || len(r.URL.Query()["pwhash"][0]) != 128 {
		return prepUser, errors.New("invalid params passed")
	}

	prepUser, err = database.ValidateUser(r.URL.Query()["username"][0], r.URL.Query()["pwhash"][0])
	if err != nil {
		log.Println("did not validate user", err)
		return prepUser, err
	}

	return prepUser, nil
}

func registerUser() {
	http.HandleFunc("/user/register", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)

		// to register, we need the name, username and pwhash
		if r.URL.Query()["name"] == nil || r.URL.Query()["username"] == nil || r.URL.Query()["pwd"] == nil || r.URL.Query()["seedpwd"] == nil {
			log.Println("missing basic set of params that can be used to validate a user")
			responseHandler(w, r, StatusBadRequest)
			return
		}

		name := r.URL.Query()["name"][0]
		username := r.URL.Query()["username"][0]
		pwd := r.URL.Query()["pwd"][0]
		seedpwd := r.URL.Query()["seedpwd"][0]

		user, err := database.NewUser(username, pwd, seedpwd, name)
		if err != nil {
			log.Println(err)
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		MarshalSend(w, r, user)
	})
}

// ValidateUser is a route that helps validate users on the platform
func ValidateUser() {
	http.HandleFunc("/user/validate", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)
		// need to pass the pwhash param here
		prepUser, err := UserValidateHelper(w, r)
		if err != nil {
			log.Println("did not validate user", err)
			responseHandler(w, r, StatusUnauthorized)
			return
		}
		// no we need to see whether this guy is an investor or a recipient.
		var prepInvestor database.Investor
		var prepRecipient database.Recipient
		var prepEntity opensolar.Entity
		rec := false
		entity := false
		prepInvestor, err = database.RetrieveInvestor(prepUser.Index)
		if err != nil {
			log.Println("did not retrieve investor", err)
			// means the user is a recipient, retrieve recipient credentials
			prepRecipient, err = database.ValidateRecipient(r.URL.Query()["username"][0], r.URL.Query()["pwhash"][0])
			if err != nil {
				log.Println("did not validate recipient", err)
				// it is not a recipient either
				prepEntity, err = opensolar.ValidateEntity(r.URL.Query()["username"][0], r.URL.Query()["pwhash"][0])
				if err != nil {
					log.Println("did not validate entity", err)
					// not an investor, recipient or entity, must be a normal user
					MarshalSend(w, r, prepUser)
					return
				} else {
					entity = true
				}
			} else {
				rec = true
			}
		}

		// the frontend should read the received response and figure out the role of the person
		var x ValidateParams
		if rec {
			x.Role = "Recipient"
			x.Entity = removeSeedRecp(prepRecipient)
			MarshalSend(w, r, x)
		} else if entity {
			x.Role = "Entity"
			x.Entity = removeSeedEntity(prepEntity)
			MarshalSend(w, r, x)
		} else {
			x.Role = "Investor"
			x.Entity = removeSeedInv(prepInvestor)
			MarshalSend(w, r, x)
		}
	})
}

// getBalances returns a list of all balances (assets and coins) held by the user
func getBalances() {
	http.HandleFunc("/user/balances", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)
		prepUser, err := UserValidateHelper(w, r)
		if err != nil {
			log.Println("did not validate user", err)
			responseHandler(w, r, StatusUnauthorized)
			return
		}

		pubkey := prepUser.PublicKey
		balances, err := xlm.GetAllBalances(pubkey)
		if err != nil {
			log.Println("did not get all balances", err)
			responseHandler(w, r, StatusNotFound)
			return
		}
		MarshalSend(w, r, balances)
	})
}

// getXLMBalance gets the XLM balance of a user's account
func getXLMBalance() {
	http.HandleFunc("/user/balance/xlm", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)
		prepUser, err := UserValidateHelper(w, r)
		if err != nil {
			log.Println("did not validate user", err)
			responseHandler(w, r, StatusUnauthorized)
			return
		}

		pubkey := prepUser.PublicKey
		balance, err := xlm.GetNativeBalance(pubkey)
		if err != nil {
			log.Println("did not get native balance", err)
			responseHandler(w, r, StatusNotFound)
			return
		}
		MarshalSend(w, r, balance)
	})
}

// getAssetBalance gets the balance of a specific asset
func getAssetBalance() {
	http.HandleFunc("/user/balance/asset", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)
		prepUser, err := UserValidateHelper(w, r)
		if err != nil {
			responseHandler(w, r, StatusUnauthorized)
			return
		}
		if r.URL.Query()["asset"] == nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		pubkey := prepUser.PublicKey
		asset := r.URL.Query()["asset"][0]
		balance, err := xlm.GetAssetBalance(pubkey, asset)
		if err != nil {
			log.Println("did not get assset balance", err)
			responseHandler(w, r, StatusNotFound)
			return
		}
		MarshalSend(w, r, balance)
	})
}

// getIpfsHash gets the ipfs hash of the passed string
func getIpfsHash() {
	http.HandleFunc("/ipfs/hash", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)
		_, err := UserValidateHelper(w, r)
		if err != nil {
			responseHandler(w, r, StatusUnauthorized)
			return
		}
		if r.URL.Query()["string"] == nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		hashString := r.URL.Query()["string"][0]
		hash, err := ipfs.AddStringToIpfs(hashString)
		if err != nil {
			log.Println("did not add string to ipfs", err)
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		hashCheck, err := ipfs.GetStringFromIpfs(hash)
		if err != nil || hashCheck != hashString {
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		MarshalSend(w, r, hash)
	})
}

// authKyc authenticates a user. Should ideally be part of a callback from the third
// party service that we choose
func authKyc() {
	http.HandleFunc("/user/kyc", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)
		prepUser, err := UserValidateHelper(w, r)
		if err != nil {
			responseHandler(w, r, StatusUnauthorized)
			return
		}
		if r.URL.Query()["userIndex"] == nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		uInput := utils.StoI(r.URL.Query()["userIndex"][0])
		err = prepUser.Authorize(uInput)
		if err != nil {
			log.Println("did not authorize user", err)
			responseHandler(w, r, StatusInternalServerError)
			return
		}
		responseHandler(w, r, StatusOK)
	})
}

// sendXLM sends a given amount of XLM to the destination address specified.
func sendXLM() {
	http.HandleFunc("/user/sendxlm", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)
		prepUser, err := UserValidateHelper(w, r)
		if err != nil {
			responseHandler(w, r, StatusUnauthorized)
			return
		}
		if r.URL.Query()["destination"] == nil || r.URL.Query()["amount"] == nil ||
			r.URL.Query()["seedpwd"] == nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		destination := r.URL.Query()["destination"][0]
		amount := r.URL.Query()["amount"][0]

		seedpwd := r.URL.Query()["seedpwd"][0]
		seed, err := wallet.DecryptSeed(prepUser.EncryptedSeed, seedpwd)
		if err != nil {
			log.Println("did not decrypt seed", err)
			responseHandler(w, r, StatusBadRequest)
			return
		}

		var memo string
		if r.URL.Query()["memo"] != nil {
			memo = r.URL.Query()["memo"][0]
		}

		_, txhash, err := xlm.SendXLM(destination, amount, seed, memo)
		if err != nil {
			log.Println("did not send xlm", err)
			responseHandler(w, r, StatusBadRequest)
			return
		}
		MarshalSend(w, r, txhash)
	})
}

// notKycView returns a list of all the users who have not yet been verified through KYC. Called by KYC Inspectors
func notKycView() {
	http.HandleFunc("/user/notkycview", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)
		prepUser, err := UserValidateHelper(w, r)
		if err != nil {
			log.Println("did not validate user", err)
			responseHandler(w, r, StatusUnauthorized)
			return
		}
		if !prepUser.Inspector {
			responseHandler(w, r, StatusUnauthorized)
			return
		}

		users, err := database.RetrieveAllUsersWithoutKyc()
		if err != nil {
			log.Println("did not retrieve all users wihtout kyc", err)
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		MarshalSend(w, r, users)
	})
}

// kycView returns a list of all the users who have been verified through KYC. Called by KYC Inspectors
func kycView() {
	http.HandleFunc("/user/kycview", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)
		prepUser, err := UserValidateHelper(w, r)
		if err != nil {
			log.Println("did not validate user", err)
			responseHandler(w, r, StatusUnauthorized)
			return
		}
		if !prepUser.Inspector {
			responseHandler(w, r, StatusUnauthorized)
			return
		}

		users, err := database.RetrieveAllUsersWithKyc()
		if err != nil {
			log.Println("did not retrieve users", err)
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		MarshalSend(w, r, users)
	})
}

// askForCoins asks for coins from the testnet faucet. Will be disabled once we move to testnet
func askForCoins() {
	http.HandleFunc("/user/askxlm", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)
		prepUser, err := UserValidateHelper(w, r)
		if err != nil {
			log.Println("did not validate user", err)
			responseHandler(w, r, StatusUnauthorized)
			return
		}

		err = xlm.GetXLM(prepUser.PublicKey)
		if err != nil {
			log.Println("did not get xlm from friendbot", err)
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		responseHandler(w, r, StatusOK)
	})
}

// trustAsset creates a trustline for the given limit with a remote peer for receiving that asset.
func trustAsset() {
	http.HandleFunc("/user/trustasset", func(w http.ResponseWriter, r *http.Request) {
		// since this is testnet, give caller coins from the testnet faucet
		checkGet(w, r)
		checkOrigin(w, r)
		prepUser, err := UserValidateHelper(w, r)
		if err != nil {
			log.Println("did not validate user", err)
			responseHandler(w, r, StatusUnauthorized)
			return
		}

		assetCode := r.URL.Query()["assetCode"][0]
		assetIssuer := r.URL.Query()["assetIssuer"][0]
		limit := r.URL.Query()["limit"][0]

		seedpwd := r.URL.Query()["seedpwd"][0]
		seed, err := wallet.DecryptSeed(prepUser.EncryptedSeed, seedpwd)
		if err != nil {
			log.Println("did not decrypt seed", err)
			responseHandler(w, r, StatusBadRequest)
			return
		}

		// func TrustAsset(assetCode string, assetIssuer string, limit string, PublicKey string, Seed string) (string, error) {
		txhash, err := assets.TrustAsset(assetCode, assetIssuer, limit, prepUser.PublicKey, seed)
		if err != nil {
			log.Println("did not trust asset", err)
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		MarshalSend(w, r, txhash)
	})
}

// uploadFile uploads a fil to ipfs and returns the ipfs hash of the uploaded file
// this is a POST request
func uploadFile() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		checkPost(w, r)
		checkOrigin(w, r)
		_, err := UserValidateHelper(w, r)
		if err != nil {
			log.Println("did not validate user", err)
			responseHandler(w, r, StatusUnauthorized)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			log.Println("did not parse form", err)
			responseHandler(w, r, StatusBadRequest)
			return
		}
		defer file.Close()

		supportedType := false
		header := fileHeader.Header.Get("Content-Type")

		switch header {
		case "image/jpeg":
			supportedType = true
		case "image/png":
			supportedType = true
		case "application/pdf":
			supportedType = true
		}

		// can't do anything with extensions, so while decrypting from ipfs, we can attach
		// all three types and return to the user.
		if !supportedType {
			responseHandler(w, r, StatusNotAcceptable)
			return
		}

		// file type is supported, store in ipfs
		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println("did not  read", err)
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		hashString, err := ipfs.IpfsHashData(data)
		if err != nil {
			log.Println("did not hash data to ipfs", err)
			responseHandler(w, r, StatusInternalServerError)
			return
		}
		MarshalSend(w, r, hashString)
	})
}

// PlatformEmailResponse is a structure used to contain the platform's email response
type PlatformEmailResponse struct {
	Email string
}

func platformEmail() {
	http.HandleFunc("/platformemail", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)
		_, err := UserValidateHelper(w, r)
		if err != nil {
			log.Println("did not validate user", err)
			responseHandler(w, r, StatusUnauthorized)
			return
		}

		var x PlatformEmailResponse
		x.Email = consts.PlatformEmail
		MarshalSend(w, r, x)
	})
}

func sendTellerShutdownEmail() {
	http.HandleFunc("/tellershutdown", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)
		prepUser, err := UserValidateHelper(w, r)
		if err != nil {
			responseHandler(w, r, StatusUnauthorized)
			return
		}
		if r.URL.Query()["projIndex"] == nil || r.URL.Query()["deviceId"] == nil ||
			r.URL.Query()["tx1"] == nil || r.URL.Query()["tx2"] == nil {
			log.Println("did not validate user", err)
			responseHandler(w, r, StatusBadRequest)
			return
		}

		projIndex := r.URL.Query()["projIndex"][0]
		deviceId := r.URL.Query()["deviceId"][0]
		tx1 := r.URL.Query()["tx1"][0]
		tx2 := r.URL.Query()["tx2"][0]
		notif.SendTellerShutdownEmail(prepUser.Email, projIndex, deviceId, tx1, tx2)
		responseHandler(w, r, StatusOK)
	})
}

func sendTellerFailedPaybackEmail() {
	http.HandleFunc("/tellerpayback", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)
		prepUser, err := UserValidateHelper(w, r)
		if err != nil {
			responseHandler(w, r, StatusUnauthorized)
			return
		}
		if r.URL.Query()["projIndex"] == nil || r.URL.Query()["deviceId"] == nil {
			log.Println("did not validate user", err)
			responseHandler(w, r, StatusBadRequest)
			return
		}

		projIndex := r.URL.Query()["projIndex"][0]
		deviceId := r.URL.Query()["deviceId"][0]
		notif.SendTellerPaymentFailedEmail(prepUser.Email, projIndex, deviceId)
		responseHandler(w, r, StatusOK)
	})
}

func tellerPing() {
	http.HandleFunc("/tellerping", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)
		_, err := UserValidateHelper(w, r)
		if err != nil {
			log.Println("did not validate user", err)
			responseHandler(w, r, StatusUnauthorized)
			return
		}

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}

		req, err := http.NewRequest("GET", TellerUrl+"/ping", nil)
		if err != nil {
			log.Println("did not create new GET request", err)
			responseHandler(w, r, StatusBadRequest)
			return
		}

		req.Header.Set("Origin", "localhost")
		res, err := client.Do(req)
		if err != nil {
			log.Println("did not make request", err)
			responseHandler(w, r, StatusBadRequest)
			return
		}
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		var x StatusResponse

		err = x.UnmarshalJSON(data)
		if err != nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		MarshalSend(w, r, x)
	})
}

func increaseTrustLimit() {
	http.HandleFunc("/user/increasetrustlimit", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)
		prepUser, err := UserValidateHelper(w, r)
		if err != nil {
			responseHandler(w, r, StatusUnauthorized)
			return
		}
		if r.URL.Query()["trust"] == nil || r.URL.Query()["seedpwd"] == nil {
			log.Println("did not validate user", err)
			responseHandler(w, r, StatusBadRequest)
			return
		}

		// now the user is validated, we need to call the db function to increase the trust limit
		trust := r.URL.Query()["trust"][0]
		seedpwd := r.URL.Query()["seedpwd"][0]

		err = database.IncreaseTrustLimit(prepUser.Index, seedpwd, trust)
		if err != nil {
			log.Println(err)
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		responseHandler(w, r, StatusOK)
	})
}

// AddContractHash adds a specific contract hash to the database
func addContractHash() {
	http.HandleFunc("/utils/addhash", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)
		var err error
		_, err = UserValidateHelper(w, r)
		if err != nil {
			responseHandler(w, r, StatusUnauthorized)
			return
		}
		if r.URL.Query()["projIndex"] == nil {
			log.Println("couldn't validate investor", err)
			responseHandler(w, r, StatusBadRequest)
			return
		}
		if r.URL.Query()["choice"] == nil || r.URL.Query()["choicestr"] == nil {
			log.Println("choice of ipfs hash not given. quitting!")
			responseHandler(w, r, StatusBadRequest)
			return
		}
		choice := r.URL.Query()["choice"][0]
		hashString := r.URL.Query()["choicestr"][0]
		projIndex, err := utils.StoICheck(r.URL.Query()["projIndex"][0])
		if err != nil {
			log.Println("passed project index not int, quitting!")
			responseHandler(w, r, StatusBadRequest)
			return
		}

		project, err := opensolar.RetrieveProject(projIndex)
		if err != nil {
			log.Println("couldn't retrieve prject index from database")
			responseHandler(w, r, StatusInternalServerError)
			return
		}
		// there are in total 5 types of hashes: OriginatorMoUHash, ContractorContractHash, InvPlatformContractHash, RecPlatformContractHash, SpecSheetHash
		// lets have a fixed set of strings that we can map on here so we have a single endpoitn for storing all these hashes

		// TODO: right now any entity can add the required hashes but in the future we must restrict adding hashes
		// to entities that are associated with the particular hashes
		// TODO: change this based on different stages. right now static
		switch choice {
		case "omh":
			// update the originator mou hash
			project.StageData = append(project.StageData, hashString)
		case "cch":
			project.StageData = append(project.StageData, hashString)
		case "ipch":
			project.StageData = append(project.StageData, hashString)
		case "rpch":
			project.StageData = append(project.StageData, hashString)
		case "ssh":
			project.StageData = append(project.StageData, hashString)
		default:
			log.Println("invalid choice passed, quitting!")
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		err = project.Save()
		if err != nil {
			log.Println("error while saving project to db, quitting!")
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		responseHandler(w, r, StatusOK)
	})
}

// sendSecrets sends secrets out to the email ids passed. This does not require the seedpwd since one can generate a new seed
// anyway using the username and password, so possessing the secrets does not require seed authentication
func sendSecrets() {
	http.HandleFunc("/user/sendrecovery", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)

		user, err := UserValidateHelper(w, r)
		if err != nil {
			responseHandler(w, r, StatusUnauthorized)
			return
		}
		if r.URL.Query()["email1"] == nil || r.URL.Query()["email2"] == nil || r.URL.Query()["email3"] == nil {
			log.Println("couldn't validate investor", err)
			responseHandler(w, r, StatusBadRequest)
			return
		}

		// we should distribute the shares and then set them to nil since a person who is in
		// control of the server c ould then reconstruct the seed
		// now send emails out to these three trusted entities with the share
		email1 := r.URL.Query()["email1"][0]
		email2 := r.URL.Query()["email2"][0]
		email3 := r.URL.Query()["email3"][0]

		err = notif.SendSecretsEmail(user.Email, email1, email2, email3, user.RecoveryShares[0], user.RecoveryShares[1], user.RecoveryShares[2])
		if err != nil {
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		// set the stored shares to nil since possessing them would enable an attacker to generate the secrets he needs by simply controlling the server
		user.RecoveryShares[0] = ""
		user.RecoveryShares[1] = ""
		user.RecoveryShares[2] = ""

		responseHandler(w, r, StatusOK)
	})
}

type SeedResponse struct {
	Seed string
}

// mergeSecrets takes in two shares in a 2 of 3 Shamir Secret Sharing Scheme and reconstructs the seed
func mergeSecrets() {
	http.HandleFunc("/user/seedrecovery", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)

		_, err := UserValidateHelper(w, r)
		if err != nil {
			responseHandler(w, r, StatusUnauthorized)
			return
		}
		if r.URL.Query()["secret1"] == nil || r.URL.Query()["secret2"] == nil {
			log.Println("couldn't validate investor", err)
			responseHandler(w, r, StatusBadRequest)
			return
		}

		var shares []string
		secret1 := r.URL.Query()["secret1"][0]
		secret2 := r.URL.Query()["secret2"][0]
		shares = append(shares, secret1, secret2)
		// now we have 2 out of the 3 secrets needed to reconstruct. Reconstruct the seed.
		secret, err := recovery.Combine(shares)
		if err != nil {
			log.Println("couldn't combine shares: ", err)
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		var x SeedResponse
		x.Seed = secret
		MarshalSend(w, r, x)
	})
}

// generateNewSecrets generates an ew set of secrets for the given function
func generateNewSecrets() {
	http.HandleFunc("/user/newsecrets", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)

		user, err := UserValidateHelper(w, r)
		if err != nil {
			responseHandler(w, r, StatusUnauthorized)
			return
		}
		if r.URL.Query()["seedpwd"] == nil || r.URL.Query()["email1"] == nil ||
			r.URL.Query()["email2"] == nil || r.URL.Query()["email3"] == nil {
			log.Println("couldn't validate investor", err)
			responseHandler(w, r, StatusBadRequest)
			return
		}

		seedpwd, err := ValidateSeedPwd(w, r, user.EncryptedSeed, user.PublicKey)
		if err != nil {
			log.Println(err)
			responseHandler(w, r, StatusBadRequest)
			return
		}

		seed, err := wallet.DecryptSeed(user.EncryptedSeed, seedpwd)
		if err != nil {
			log.Println(err)
			responseHandler(w, r, StatusBadRequest)
			return
		}
		// user has validated his seed and identity. Generate new shares and send them out
		shares, err := recovery.Create(2, 3, seed)
		if err != nil {
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		email1 := r.URL.Query()["email1"][0]
		email2 := r.URL.Query()["email2"][0]
		email3 := r.URL.Query()["email3"][0]

		err = notif.SendSecretsEmail(user.Email, email1, email2, email3, shares[0], shares[1], shares[2])
		if err != nil {
			responseHandler(w, r, StatusInternalServerError)
			return
		}
		responseHandler(w, r, StatusOK)
	})
}

func generateResetPwdCode() {
	http.HandleFunc("/user/resetpwd", func(w http.ResponseWriter, r *http.Request) {
		log.Println("CALLING ENDPOINT")
		checkGet(w, r)
		checkOrigin(w, r)

		// the notion here si that the user must have his seedpwd in order to reset the password.
		// we retrieve the user using his email id and lookup his encrypted seed. If the
		// seed can be unlocked using hte seedpwd, we send a pwd reset email. One of two passwords
		// must be remembered
		if r.URL.Query()["email"] == nil || r.URL.Query()["seedpwd"] == nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}
		email := r.URL.Query()["email"][0]

		rUser, err := database.SearchWithEmailId(email)
		if err != nil {
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		_, err = ValidateSeedPwd(w, r, rUser.EncryptedSeed, rUser.PublicKey)
		if err != nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}
		// now we can verify that this is rellay the user. Now we need to cgenerate a verification code
		// and send it over to the user.
		verificationCode := utils.GetRandomString(16)
		log.Println("VERIFICATION CODE: ", verificationCode)
		rUser.PwdResetCode = verificationCode
		err = rUser.Save()
		if err != nil {
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		// now send this verification code to the email we have in the database
		err = notif.SendPasswordResetEmail(rUser.Email, verificationCode)
		if err != nil {
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		responseHandler(w, r, StatusOK)
	})
}

func resetPassword() {
	http.HandleFunc("/user/pwdreset", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)

		if r.URL.Query()["email"] == nil || r.URL.Query()["seedpwd"] == nil || r.URL.Query()["verificationCode"] == nil ||
			r.URL.Query()["pwhash"] == nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		email := r.URL.Query()["email"][0]
		vCode := r.URL.Query()["verificationCode"][0]
		pwhash := r.URL.Query()["pwhash"][0]

		rUser, err := database.SearchWithEmailId(email)
		if err != nil {
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		_, err = ValidateSeedPwd(w, r, rUser.EncryptedSeed, rUser.PublicKey)
		if err != nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		if vCode != rUser.PwdResetCode || vCode == "INVALID" {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		// reset the user's password
		rUser.Pwhash = pwhash
		rUser.PwdResetCode = "INVALID" // invalidate the pwd reset code to avoid replay attacks
		err = rUser.Save()
		if err != nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		responseHandler(w, r, StatusOK)
	})
}

// sweepFunds tries to sweep all funds that we have from one account to another. Requires
// the seedpwd. Can't transfre assets automatically since platform does not know the list
// of issuer publickeys
func sweepFunds() {
	http.HandleFunc("/user/sweep", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)

		prepUser, err := UserValidateHelper(w, r)
		if err != nil {
			responseHandler(w, r, StatusUnauthorized)
			return
		}
		if r.URL.Query()["seedpwd"] == nil || r.URL.Query()["destination"] == nil {
			log.Println("did not validate user", err)
			responseHandler(w, r, StatusBadRequest)
			return
		}

		transferAddress := r.URL.Query()["destination"][0]
		if !xlm.AccountExists(transferAddress) {
			log.Println("Can only transfer to existing accounts, quitting")
			responseHandler(w, r, StatusBadRequest)
			return
		}

		seedpwd, err := ValidateSeedPwd(w, r, prepUser.EncryptedSeed, prepUser.PublicKey)
		if err != nil {
			log.Println(err)
			responseHandler(w, r, StatusBadRequest)
			return
		}

		seed, err := wallet.DecryptSeed(prepUser.EncryptedSeed, seedpwd)
		if err != nil {
			log.Println(err)
			responseHandler(w, r, StatusBadRequest)
			return
		}
		// validated the user, so now proceed to sweep funds
		xlmBalance, err := xlm.GetNativeBalance(prepUser.PublicKey)
		if err != nil {
			log.Println(err)
			responseHandler(w, r, StatusInternalServerError)
			return
		}
		xlmBalanceF, err := utils.StoFWithCheck(xlmBalance)
		if err != nil {
			log.Println(err)
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		// reduce 0.05 xlm and then sweep funds
		if xlmBalanceF < 5 {
			log.Println("xlm balance for user too small to sweep funds, quitting!")
			responseHandler(w, r, StatusBadRequest)
			return
		}
		xlmBalanceF -= 5
		// now we have the xlm balance, shift funds to the other account as requested by the user.
		sweepAmt := math.Round(xlmBalanceF)
		sweepStr := utils.FtoS(sweepAmt)
		_, txhash, err := xlm.SendXLM(transferAddress, sweepStr, seed, "sweep funds")
		if err != nil {
			log.Println(err)
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		log.Println("sweep funds txhash: ", txhash)
		responseHandler(w, r, StatusOK)
	})
}

// sweepAsset sweeps a given asset from one account to another. Can't transfer multiple
// assets since we require the issuer pubkey
func sweepAsset() {
	http.HandleFunc("/user/sweepasset", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		checkOrigin(w, r)

		prepUser, err := UserValidateHelper(w, r)
		if err != nil {
			responseHandler(w, r, StatusUnauthorized)
			return
		}
		if r.URL.Query()["seedpwd"] == nil || r.URL.Query()["destination"] == nil ||
			r.URL.Query()["assetName"] == nil || r.URL.Query()["issuerPubkey"] == nil {
			log.Println("did not validate user", err)
			responseHandler(w, r, StatusBadRequest)
			return
		}

		assetName := r.URL.Query()["assetName"][0]
		destination := r.URL.Query()["destination"][0]
		issuerPubkey := r.URL.Query()["issuerPubkey"][0]

		seedpwd, err := ValidateSeedPwd(w, r, prepUser.EncryptedSeed, prepUser.PublicKey)
		if err != nil {
			log.Println(err)
			responseHandler(w, r, StatusBadRequest)
			return
		}

		seed, err := wallet.DecryptSeed(prepUser.EncryptedSeed, seedpwd)
		if err != nil {
			log.Println(err)
			responseHandler(w, r, StatusBadRequest)
			return
		}

		// validated the user, so now proceed to sweep funds
		assetBalance, err := xlm.GetAssetBalance(prepUser.PublicKey, assetName)
		if err != nil {
			log.Println(err)
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		assetBalanceF, err := utils.StoFWithCheck(assetBalance)
		if err != nil {
			log.Println(err)
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		// reduce 0.05 xlm and then sweep funds
		if assetBalanceF < 5 {
			log.Println("asset balance for user too smal lto sweep funds, quitting!")
			responseHandler(w, r, StatusBadRequest)
			return
		} else {
			assetBalanceF -= 5
		}

		sweepAmt := math.Round(assetBalanceF)
		sweepStr := utils.FtoS(sweepAmt)

		_, txhash, err := assets.SendAsset(assetName, issuerPubkey, destination, sweepStr, seed, prepUser.PublicKey, "sweeping funds")
		if err != nil {
			log.Println(err)
			responseHandler(w, r, StatusInternalServerError)
			return
		}

		log.Println("txhash: ", txhash)
		responseHandler(w, r, StatusOK)
	})
}

func ValidateSeedPwd(w http.ResponseWriter, r *http.Request, encryptedSeed []byte, userPublickey string) (string, error) {
	seedpwd := r.URL.Query()["seedpwd"][0]
	// we've validated the seedpwd, try decrypting the Encrypted Seed.
	seed, err := wallet.DecryptSeed(encryptedSeed, seedpwd)
	if err != nil {
		return seedpwd, fmt.Errorf("could not decrypt seed")
	}

	// now get the pubkey from this seed and match with original pubkey
	pubkey, err := wallet.ReturnPubkey(seed)
	if err != nil {
		return seedpwd, fmt.Errorf("could not retrieve pubkey")
	}

	if pubkey != userPublickey {
		return seedpwd, fmt.Errorf("pubkeys don't match, quitting")
	}

	return seedpwd, nil
}
