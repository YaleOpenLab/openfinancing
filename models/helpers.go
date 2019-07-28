package models

import (
	"github.com/pkg/errors"
	"log"
	"time"

	xlm "github.com/Varunram/essentials/crypto/xlm"
	assets "github.com/Varunram/essentials/crypto/xlm/assets"
	utils "github.com/Varunram/essentials/utils"
	consts "github.com/YaleOpenLab/openx/consts"
)

// the models package won't be imported directly in any place but would be imported
// by all the investment models that exist

// SendUSDToPlatform sends STABLEUSD back to the platform for investment
func SendUSDToPlatform(invSeed string, invAmount float64, memo string) (string, error) {
	// send stableusd to the platform (not the issuer) since the issuer will be locked
	// and we can't use the funds. We also need ot be able to redeem the stablecoin for fiat
	// so we can't burn them

	var oldPlatformBalance string
	var err error
	oldPlatformBalance, err = xlm.GetAssetBalance(consts.PlatformPublicKey, consts.StablecoinCode)
	if err != nil {
		// platform does not have stablecoin, shouldn't arrive here ideally
		oldPlatformBalance = "0"
	}

	_, txhash, err := assets.SendAsset(consts.StablecoinCode, consts.StablecoinPublicKey, consts.PlatformPublicKey, invAmount, invSeed, memo)
	if err != nil {
		return txhash, errors.Wrap(err, "sending stableusd to platform failed")
	}

	log.Println("Sent STABLEUSD to platform, confirmation: ", txhash)
	time.Sleep(5 * time.Second) // wait for a block

	newPlatformBalance, err := xlm.GetAssetBalance(consts.PlatformPublicKey, consts.StablecoinCode)
	if err != nil {
		return txhash, errors.Wrap(err, "error while getting asset balance")
	}

	npBS, err := utils.ToFloat(newPlatformBalance)
	if err != nil {
		return txhash, err
	}

	opBS, err := utils.ToFloat(oldPlatformBalance)
	if err != nil {
		return txhash, err
	}

	if npBS-opBS < invAmount-1 {
		return txhash, errors.New("Sent amount doesn't match with investment amount")
	}
	return txhash, nil
}
