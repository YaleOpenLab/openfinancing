package rpc

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	database "github.com/OpenFinancing/openfinancing/database"
	solar "github.com/OpenFinancing/openfinancing/platforms/solar"
	utils "github.com/OpenFinancing/openfinancing/utils"
	wallet "github.com/OpenFinancing/openfinancing/wallet"
	xlm "github.com/OpenFinancing/openfinancing/xlm"
)

// setupInvestorRPCs sets up all RPCs related to the investor
func setupInvestorRPCs() {
	insertInvestor()
	validateInvestor()
	getAllInvestors()
	investInProject()
	changeReputationInv()
	voteTowardsProject()
}

func parseInvestor(r *http.Request) (database.Investor, error) {
	var prepInvestor database.Investor
	err := r.ParseForm()
	if err != nil || r.FormValue("username") == "" || r.FormValue("pwhash") == "" || r.FormValue("Name") == "" || r.FormValue("EPassword") == "" {
		return prepInvestor, fmt.Errorf("One of required fields missing: username, pwhash, Name, EPassword")
	}

	prepInvestor.AmountInvested = float64(0)
	prepInvestor.U, err = database.NewUser(r.FormValue("username"), r.FormValue("pwhash"), r.FormValue("Name"), r.FormValue("EPassword"))
	return prepInvestor, err
}

func insertInvestor() {
	// this should be a post method since you want to accetp an project and then insert
	// that into the database
	http.HandleFunc("/investor/insert", func(w http.ResponseWriter, r *http.Request) {
		checkPost(w, r)
		prepInvestor, err := parseInvestor(r)
		if err != nil {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		log.Println("Prepared Investor:", prepInvestor)
		err = prepInvestor.Save()
		if err != nil {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		Send200(w, r)
	})
}

// validateInvestor retrieves the investor after valdiating if such an ivnestor exists
// by checking the pwhash of the given investor with the stored one
func validateInvestor() {
	http.HandleFunc("/investor/validate", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		if r.URL.Query() == nil || r.URL.Query()["username"] == nil || r.URL.Query()["pwhash"] == nil ||
			len(r.URL.Query()["pwhash"][0]) != 128 { // sha 512 length
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		prepInvestor, err := database.ValidateInvestor(r.URL.Query()["username"][0], r.URL.Query()["pwhash"][0])
		if err != nil {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		log.Println("Prepared Investor:", prepInvestor)
		MarshalSend(w, r, prepInvestor)
	})
}

// getAllInvestors gets a list of all the investors in the system so that we can
// display it to some entity that is interested to view such stats
func getAllInvestors() {
	http.HandleFunc("/investor/all", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		investors, err := database.RetrieveAllInvestors()
		if err != nil {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		MarshalSend(w, r, investors)
	})
}

func investInProject() {
	http.HandleFunc("/investor/invest", func(w http.ResponseWriter, r *http.Request) {
		checkGet(w, r)
		// need the following params to invest in a project:
		// 1. Seed pwhash (for the investor)
		// 2. project index
		// 3. investment amount
		// 4. Login username (for the investor)
		// 5. Login pwhash (for the investor)
		if r.URL.Query() == nil || r.URL.Query()["seedpwd"] == nil || r.URL.Query()["projIndex"] == nil ||
			r.URL.Query()["amount"] == nil || r.URL.Query()["username"] == nil || r.URL.Query()["pwhash"] == nil ||
			len(r.URL.Query()["pwhash"][0]) != 128 { // sha 512 length
			errorHandler(w, r, http.StatusNotFound)
			return
		}

		investor, err := database.ValidateInvestor(r.URL.Query()["username"][0], r.URL.Query()["pwhash"][0])
		if err != nil {
			errorHandler(w, r, http.StatusNotFound)
			return
		}

		seedpwd := r.URL.Query()["seedpwd"][0]
		investorSeed, err := wallet.DecryptSeed(investor.U.EncryptedSeed, seedpwd)
		if err != nil {
			errorHandler(w, r, http.StatusNotFound)
			return
		}

		projIndex := utils.StoI(r.URL.Query()["projIndex"][0])
		amount := r.URL.Query()["amount"][0]
		investorPubkey, err := wallet.ReturnPubkey(investorSeed)
		if err != nil {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		// splitting the conditions into two since in the future we will be returning
		// error codes towards each type
		if !xlm.AccountExists(investorPubkey) {
			errorHandler(w, r, http.StatusNotFound)
			return
		}

		// note that while using this route, we can't send the investor assets (maybe)
		// make it so in the UI that only they can accept an investment so we can get their
		// seed and send them assets. By not accepting, they would forfeit their investment,
		// so incentive would be there to unlock the seed.
		_, err = solar.InvestInProject(projIndex, investor.U.Index, amount, investorSeed)
		if err != nil {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		Send200(w, r)
	})
}

func InvValidateHelper(w http.ResponseWriter, r *http.Request) (database.Investor, error) {
	// first validate the investor or anyone would be able to set device ids
	checkGet(w, r)
	var prepInvestor database.Investor
	// need to pass the pwhash param here
	if r.URL.Query() == nil || r.URL.Query()["username"] == nil ||
		len(r.URL.Query()["pwhash"][0]) != 128 {
		return prepInvestor, fmt.Errorf("Invalid params passed")
	}

	prepInvestor, err := database.ValidateInvestor(r.URL.Query()["username"][0], r.URL.Query()["pwhash"][0])
	if err != nil {
		return prepInvestor, err
	}

	return prepInvestor, nil
}

func changeReputationInv() {
	http.HandleFunc("/investor/reputation", func(w http.ResponseWriter, r *http.Request) {
		investor, err := InvValidateHelper(w, r)
		if err != nil || r.URL.Query()["reputation"] == nil {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		reputation, err := strconv.ParseFloat(r.URL.Query()["reputation"][0], 32) // same as StoI but we need to catch the error here
		if err != nil {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		err = database.ChangeInvReputation(investor.U.Index, reputation)
		if err != nil {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		Send200(w, r)
	})
}

func voteTowardsProject() {
	http.HandleFunc("/investor/vote", func(w http.ResponseWriter, r *http.Request) {
		investor, err := InvValidateHelper(w, r)
		if err != nil || r.URL.Query()["votes"] == nil || r.URL.Query()["projIndex"] == nil {
			errorHandler(w, r, http.StatusNotFound)
			return
		}

		votes := utils.StoI(r.URL.Query()["votes"][0])
		projIndex := utils.StoI(r.URL.Query()["projIndex"][0])
		err = solar.VoteTowardsProposedProject(investor.U.Index, votes, projIndex)
		if err != nil {
			log.Println("ERROR: ", err)
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		Send200(w, r)
	})
}
