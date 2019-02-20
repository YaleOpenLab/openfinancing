package rpc

import (
	"encoding/json"
	"log"
	"net/http"
	//utils "github.com/YaleOpenLab/openx/utils"
	"io"
	"strings"
)

// we need to call the endpoitns and display the stuff returned from that endpoint.
// TODO: what do we do with the returned event streams? We could analyse it and provide a live feed
// of sorts but people who need to verify it have access to the portal anwyay. A more efficient way
// would be to write those details to a separate file and then parse that to retrieve results.

// function to setup all the particle related endpoints
func setupParticleHandlers() {
	listAllDevices()
	listProductInfo()
	getDeviceInfo()
	pingDevice()
	signalDevice()
	serialNumberInfo()
	getDiagnosticsLast()
	getAllDiagnostics()
	getParticleUserInfo()
	getAllSims()
}

// ParticleDevice is a structure to parse the returned particle.io data
type ParticleDevice struct {
	Id                    string `json:"id"`
	Name                  string `json:"name"`
	LastApp               string `json:"last_app"`
	LastIPAddress         string `json:"last_ip_address"`
	ProductID             int    `json:"product_id"`
	Connected             bool   `json:"connected"`
	PlatformID            int    `json:"platform_id"`
	Cellular              bool   `json:"cellular"`
	Notes                 string `json:"notes"`
	Status                string `json:"status"`
	SerialNumber          string `json:"serial_number"`
	CurrentBuildTarget    string `json:"current_build_target"`
	SystemFirmwareVersion string `json:"system_firmware_version"`
	DefaultBuildTarget    string `json:"default_build_target"`
}

// ParticleProductDevice is a structure to parse returned particle.io data
type ParticleProductDevice struct {
	Id                             string   `json:"id"`
	ProductID                      int      `json:"product_id"`
	LastIPAddress                  string   `json:"last_ip_address"`
	LastHandshakeAt                string   `json:"last_handshake_at"`
	UserID                         string   `json:"user_id"`
	Online                         bool     `json:"online"`
	Name                           string   `json:"name"`
	PlatformID                     int      `json:"platform_id"`
	FirmwareProductID              int      `json:"firmware_product_id"`
	Quarantined                    bool     `json:"quarantined"`
	Denied                         bool     `json:"denied"`
	Development                    bool     `json:"development"`
	Groups                         []string `json:"groups"`
	TargetedFirmwareReleaseVersion string   `json:"targeted_firmware_release_version"`
	SystemFirmwareVersion          string   `json:"system_firmware_version"`
	SerialNumber                   string   `json:"serial_number"`
	Owner                          string   `json:"owner"`
}

// ParticleProductInfo is a structure to parse returned particle.io data
type ParticleProductInfo struct {
	Devices []ParticleProductDevice
}

// ParticlePingResponse is a structure to parse returned particle.io data
type ParticlePingResponse struct {
	Online bool `json:"online"`
	Ok     bool `json:"ok"`
}

// SignalResponse is a structure to parse returned particle.io data
type SignalResponse struct {
	Id        string `json:"id"`
	Connected bool   `json:"connected"`
	Signaling bool   `json:"signaling"`
}

// SerialNumberResponse is a structure to parse returned particle.io data
type SerialNumberResponse struct {
	Ok         bool   `json:"ok"`
	DeviceID   string `json:"device_id"`
	PlatformID int    `json:"platform_id"`
}

// ParticleUser is a structure to parse returned particle.io data
type ParticleUser struct {
	Username        string   `json:"username"`
	SubscriptionIds []string `json:"subscription_ids"`
	AccountInfo     struct {
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		CompanyName     string `json:"company_name"`
		BusinessAccount bool   `json:"business_account"`
	} `json:"account_info"`
	TeamInvites         []string `json:"team_invites"`
	WifiDeviceCount     int      `json:"wifi_device_count"`
	CellularDeviceCount int      `json:"cellular_device_count"`
}

// ParticleEventStream is a structure to parse returned particle.io data
type ParticleEventStream struct {
	Data        string `json:"data"`
	Ttl         string `json:"ttl"`
	PublishedAt string `json:"published_at"`
	Coreid      string `json:"coreid"`
}

// GetAndSendJson is a handler that makes a get request and returns json data
func GetAndSendJson(w http.ResponseWriter, r *http.Request, body string, x interface{}) {
	data, err := GetRequest(body)
	if err != nil {
		log.Println("did not get response", err)
		responseHandler(w, r, StatusBadRequest)
		return
	}
	// now data is in byte, we need the other structure now
	err = json.Unmarshal(data, &x)
	if err != nil {
		log.Println("did not unmarshal json", err)
		responseHandler(w, r, StatusInternalServerError)
		return
	}
	MarshalSend(w, r, x)
}

// GetAndSendByte is a handler that makes a get request and returns byte data. THis is used
// in cases for which we don;t know the format of the returned data, so we can't parse
// what stuff is in here.
func GetAndSendByte(w http.ResponseWriter, r *http.Request, body string) {
	data, err := GetRequest(body)
	if err != nil {
		log.Println("did not get response", err)
		responseHandler(w, r, StatusBadRequest)
		return
	}

	w.Write(data)
}

// PutAndSend is a handler that PUTs data and returns the response
func PutAndSend(w http.ResponseWriter, r *http.Request, body string, payload io.Reader) {
	data, err := PutRequest(body, payload)
	if err != nil {
		log.Println("did not receive success response", err)
		responseHandler(w, r, StatusBadRequest)
		return
	}
	var x ParticlePingResponse
	err = json.Unmarshal(data, &x)
	if err != nil {
		log.Println("did not unmarshal json", err)
		responseHandler(w, r, StatusInternalServerError)
		return
	}
	MarshalSend(w, r, x)
}

// listAllDevices lists all the devices registered to the user holding the specific access token
func listAllDevices() {
	// make a curl request out to lcoalhost and get the ping response
	http.HandleFunc("/particle/devices", func(w http.ResponseWriter, r *http.Request) {
		// validate if the person requesting this is a vlaid user on the platform
		_, err := UserValidateHelper(w, r)
		if err != nil || r.URL.Query()["accessToken"] == nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		accessToken := r.URL.Query()["accessToken"][0]
		body := "https://api.particle.io/v1/devices?access_token=" + accessToken
		var x []ParticleDevice
		GetAndSendJson(w, r, body, x)
	})
}

// listProductInfo liusts all the producsts belonging to the user with the access token
func listProductInfo() {
	http.HandleFunc("/particle/productinfo", func(w http.ResponseWriter, r *http.Request) {

		_, err := UserValidateHelper(w, r)
		if err != nil || r.URL.Query()["accessToken"] == nil || r.URL.Query()["productInfo"] == nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		accessToken := r.URL.Query()["accessToken"][0]
		productInfo := r.URL.Query()["productInfo"][0]

		body := "https://api.particle.io/v1/products/" + productInfo + "/devices?access_token=" + accessToken
		var x ParticleProductInfo
		GetAndSendJson(w, r, body, x)
	})
}

// getDeviceInfo returns the information of a specific device. REquires device id and the accesstoken
func getDeviceInfo() {
	http.HandleFunc("/particle/deviceinfo", func(w http.ResponseWriter, r *http.Request) {
		// validate if the person requesting this is a vlaid user on the platform
		_, err := UserValidateHelper(w, r)
		if err != nil || r.URL.Query()["accessToken"] == nil || r.URL.Query()["deviceId"] == nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		accessToken := r.URL.Query()["accessToken"][0]
		deviceId := r.URL.Query()["deviceId"][0]

		body := "https://api.particle.io/v1/devices/" + deviceId + "?access_token=" + accessToken
		var x ParticleDevice
		GetAndSendJson(w, r, body, x)
	})
}

// pingDevice pings a specific device and sees whether its up. Could be useful to create a monitoring
// dashboard of sorts where people can see if their devices are online or not
func pingDevice() {
	http.HandleFunc("/particle/deviceping", func(w http.ResponseWriter, r *http.Request) {

		_, err := UserValidateHelper(w, r)
		if err != nil || r.URL.Query()["accessToken"] == nil || r.URL.Query()["deviceId"] == nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		accessToken := r.URL.Query()["accessToken"][0]
		deviceId := r.URL.Query()["deviceId"][0]
		body := "https://api.particle.io/v1/devices/" + deviceId + "/ping"
		payload := strings.NewReader("access_token=" + accessToken)

		PutAndSend(w, r, body, payload)
	})
}

// signalDevice sends a rainbow signal to the device and the device flashes in rainbow colors
// on receiving this signal. Can be set to on or off depending on whether we want the device to flash
// in rainbow colors or not
func signalDevice() {
	http.HandleFunc("/particle/devicesignal", func(w http.ResponseWriter, r *http.Request) {

		_, err := UserValidateHelper(w, r)
		if err != nil || r.URL.Query()["signal"] == nil || r.URL.Query()["accessToken"] == nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		accessToken := r.URL.Query()["accessToken"][0]
		deviceId := r.URL.Query()["deviceId"][0]
		signal := r.URL.Query()["signal"][0]
		if signal != "on" && signal != "off" {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		var body string
		var payload io.Reader
		body = "https://api.particle.io/v1/devices/" + deviceId
		if signal == "ok" {
			payload = strings.NewReader("signal=" + "1" + "&access_token=" + accessToken)
			body += "?signal=" + "1" + "&accessToken=" + accessToken
		} else {
			payload = strings.NewReader("signal=" + "0" + "&access_token=" + accessToken)
		}

		PutAndSend(w, r, body, payload)
	})
}

// serialNumberInfo gets the device id of a device on recipt of the serial number
func serialNumberInfo() {
	http.HandleFunc("/particle/getdeviceid", func(w http.ResponseWriter, r *http.Request) {

		_, err := UserValidateHelper(w, r)
		if err != nil || r.URL.Query()["serialNumber"] == nil || r.URL.Query()["accessToken"] == nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		serialNumber := r.URL.Query()["serialNumber"][0]
		accessToken := r.URL.Query()["accessToken"][0]

		body := "https://api.particle.io/v1/serial_numbers/" + serialNumber + "?access_token=" + accessToken
		var x SerialNumberResponse
		GetAndSendJson(w, r, body, x)
	})
}

// getDiagnosticsLast gets a list of the last diagnostic report that belongs to the specific device
func getDiagnosticsLast() {
	http.HandleFunc("/particle/diag/last", func(w http.ResponseWriter, r *http.Request) {

		_, err := UserValidateHelper(w, r)
		if err != nil || r.URL.Query()["accessToken"] == nil || r.URL.Query()["deviceId"] == nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		accessToken := r.URL.Query()["accessToken"][0]
		deviceId := r.URL.Query()["deviceId"][0]

		body := "https://api.particle.io/v1/diagnostics/" + deviceId + "/last?access_token=" + accessToken
		GetAndSendByte(w, r, body)
	})
}

// getAllDiagnostics gets all the past diagnostic reports of the associated device id. Requires
// accessToken for authentication
func getAllDiagnostics() {
	http.HandleFunc("/particle/diag/all", func(w http.ResponseWriter, r *http.Request) {

		_, err := UserValidateHelper(w, r)
		if err != nil || r.URL.Query()["accessToken"] == nil || r.URL.Query()["deviceId"] == nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		accessToken := r.URL.Query()["accessToken"][0]
		deviceId := r.URL.Query()["deviceId"][0]

		body := "https://api.particle.io/v1/diagnostics/" + deviceId + "?access_token=" + accessToken
		GetAndSendByte(w, r, body)
	})
}

// getParticleUserInfo gets the information of a particular user associated with an accessToken
func getParticleUserInfo() {
	http.HandleFunc("/particle/user/info", func(w http.ResponseWriter, r *http.Request) {

		_, err := UserValidateHelper(w, r)
		if err != nil || r.URL.Query()["accessToken"] == nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		accessToken := r.URL.Query()["accessToken"][0]
		body := "https://api.particle.io/v1/user?access_token=" + accessToken
		var x ParticleUser
		GetAndSendJson(w, r, body, x)
	})
}

// getAllSims gets the informatiomn of all sim card that areassociated with the particular accessToken
func getAllSims() {
	http.HandleFunc("/particle/sims", func(w http.ResponseWriter, r *http.Request) {

		_, err := UserValidateHelper(w, r)
		if err != nil || r.URL.Query()["accessToken"] == nil {
			responseHandler(w, r, StatusBadRequest)
			return
		}

		accessToken := r.URL.Query()["accessToken"][0]

		body := "https://api.particle.io/v1/sims?access_token=" + accessToken
		GetAndSendByte(w, r, body)
	})
}
