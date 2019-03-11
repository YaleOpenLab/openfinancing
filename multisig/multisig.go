package multisig

import (
	"fmt"
	"log"
	"net/http"

	xlm "github.com/YaleOpenLab/openx/xlm"
	"github.com/stellar/go/build"
	clients "github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/network"
)

/*
InflationDest("GCT7S5BA6ZC7SV7GGEMEYJTWOBYTBOA7SC4JEYP7IAEDG7HQNIWKRJ4G"),
	SetAuthRequired(),
	SetAuthRevocable(),
	SetAuthImmutable(),
	ClearAuthRequired(),
	ClearAuthRevocable(),
	ClearAuthImmutable(),
	MasterWeight(1),
	SetThresholds(2, 3, 4),
	HomeDomain("stellar.org"),
	AddSigner("GC6DDGPXVWXD5V6XOWJ7VUTDYI7VKPV2RAJWBVBHR47OPV5NASUNHTJW", 5),
*/
var TestNetClient = &clients.Client{
	// URL: "http://35.192.122.229:8080",
	URL:  "https://horizon-testnet.stellar.org",
	HTTP: http.DefaultClient,
}

// Convert2of2 converts the accoutn wiht pubeky myPubkey to a 2of2 multisig account
func Convert2of2(myPubkey string, mySeed string, cosignerPubkey string) error {
	// don't check if the account exists or not, hopefully it does
	memo := "testsign"
	amount := "1"

	tx, err := build.Transaction(
		build.SourceAccount{myPubkey},
		build.AutoSequence{TestNetClient},
		build.Network{network.TestNetworkPassphrase},
		build.MemoText{memo},
		build.Payment(
			build.Destination{myPubkey},
			build.NativeAmount{amount},
		),
		build.SetOptions(
			build.MasterWeight(1),
			build.AddSigner(cosignerPubkey, 1), // add x-1 signers here
			build.SetThresholds(2, 2, 2),       // set all thresholds to the threshold you want
		),
	)

	if err != nil {
		log.Println("error while constructing tx", err)
		return err
	}

	_, _, err = xlm.SendTx(mySeed, tx)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// AddSignerTx is used to add a signer to the account with Public Key pubkey
func AddSignerTx(seed string, pubkey string, cosignerPubkey string) error {
	memo := "addsigner"
	amount := "1" // some token amount, this can be any number though (even larger than the number of xlm in  xistence)

	tx, err := build.Transaction(
		build.SourceAccount{pubkey},
		build.AutoSequence{TestNetClient},
		build.Network{network.TestNetworkPassphrase},
		build.MemoText{memo},
		build.Payment(
			build.Destination{pubkey},
			build.NativeAmount{amount},
		),
		build.SetOptions(
			build.AddSigner(cosignerPubkey, 1), // add first signer
		),
	)

	if err != nil {
		log.Println("error while constructing tx", err)
		return err
	}

	_, _, err = xlm.SendTx(seed, tx)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

// when the number of tx's reaches x-1, call the threshold tx to set thresholds
func ConstructThresholdTx(seed string, pubkey string, cosignerPubkey string, y int) error {
	memo := "sealthreshold"
	amount := "1" // some token amount, this can be any number though (even larger than the number of xlm in  xistence)
	x := uint32(y)

	tx, err := build.Transaction(
		build.SourceAccount{pubkey},
		build.AutoSequence{TestNetClient},
		build.Network{network.TestNetworkPassphrase},
		build.MemoText{memo},
		build.Payment(
			build.Destination{pubkey},
			build.NativeAmount{amount},
		),
		build.SetOptions(
			build.MasterWeight(0),              // set the seed od account 2 to have zero weight
			build.AddSigner(cosignerPubkey, 1), // add second signer
			build.SetThresholds(x, x, x),       // set all thresholds to the threshold you want
		),
	)

	if err != nil {
		log.Println("error while constructing tx", err)
		return err
	}

	_, _, err = xlm.SendTx(seed, tx)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

func New1of2(cosigner1Pubkey string, cosigner2Pubkey string) (string, error) {

	tempSeed, pubkey, err := xlm.GetKeyPair()
	if err != nil {
		log.Println(err)
		return "", err
	}

	// setup both accounts
	err = xlm.GetXLM(pubkey)
	if err != nil {
		log.Println(err)
		return pubkey, err
	}

	err = AddSignerTx(tempSeed, pubkey, cosigner1Pubkey)
	if err != nil {
		log.Println(err)
		return pubkey, err
	}

	// we've reached x-1 = 1 signers, call threshold tx
	// threshold is 1 here since we need a 1 of 2 multisig account
	err = ConstructThresholdTx(tempSeed, pubkey, cosigner2Pubkey, 1)
	if err != nil {
		log.Println(err)
		return pubkey, err
	}

	return pubkey, nil
}

func New2of2(cosigner1Pubkey string, cosigner2Pubkey string) (string, error) {
	// don't check if the account exists or not, hopefully it does

	tempSeed, pubkey, err := xlm.GetKeyPair()
	if err != nil {
		log.Println(err)
		return "", err
	}

	// setup both accounts
	err = xlm.GetXLM(pubkey)
	if err != nil {
		log.Println(err)
		return pubkey, err
	}

	err = AddSignerTx(tempSeed, pubkey, cosigner1Pubkey)
	if err != nil {
		log.Println(err)
		return pubkey, err
	}

	// we've reached x-1 = 1 signers, call threshold tx
	err = ConstructThresholdTx(tempSeed, pubkey, cosigner2Pubkey, 2)
	if err != nil {
		log.Println(err)
		return pubkey, err
	}

	return pubkey, nil
}

// Construct2of2Tx constructs a tx where the source account pubkey1 is the 2of2 account
// we need 2 signers for this tx
func Tx2of2(pubkey1 string, destination string, signer1 string, signer2 string) error {

	memo := "testmultisig"

	// construct a tx sending coins from account 1 to account 1
	tx, err := build.Transaction(
		build.SourceAccount{pubkey1},
		build.AutoSequence{TestNetClient},
		build.Network{network.TestNetworkPassphrase},
		build.MemoText{memo},
		build.Payment(
			build.Destination{pubkey1},
			build.NativeAmount{"1"},
		),
	)

	txe, err := tx.Sign(signer1, signer2) // sign using party 2's seed
	if err != nil {
		log.Println("second party couldn't sign tx", err)
		return err
	}

	txeB64, err := txe.Base64()
	if err != nil {
		log.Println(err)
		return err
	}

	resp, err := TestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		log.Println("error while submitting tx", err)
		return err
	}

	log.Printf("Two party multisig tx: %s, sequence: %d\n", resp.Hash, resp.Ledger)
	return nil
}
