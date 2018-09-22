package payments

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

//PAYPAL API DATA
var paypal = struct {
	APIKey string
}{
	APIKey: "sfsdf",
}

type paypalResponseToken struct {
	Scope       string `json:"scope"`
	Nonce       string `json:"nonce"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	AppID       string `json:"app_id"`
	ExpiresIn   int    `json:"expires_in"`
}

////////////////////////////////////////////////
//////EXTERNAL CHECKOUT WITHOUT REDIRECTS///////
////////////////////////////////////////////////

func getPaypalBearerToken() (bearerToken string, err error) {
	req, err := http.NewRequest("GET", "https://api.sandbox.paypal.com/v1/oauth2/token", bytes.NewReader([]byte{}))

	if err != nil {
		return
	}

	req.Header.Add("client_id", paypal.APIKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en_US")
	req.Header.Add("grant_type", "client_credentials")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return
	}

	responseData := &paypalResponseToken{}

	err = json.NewDecoder(res.Body).Decode(responseData)

	if err != nil {
		return
	}

	bearerToken = responseData.AccessToken

	return
}

//ButtonCheckoutPaypal execute a approved paypal pay
func (wcp *WebCheckoutPayment) ButtonCheckoutPaypal() (err error) {

	bearerToken, err := getPaypalBearerToken()

	if err != nil {
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api.sandbox.paypal.com/v1/payments/payment/"+wcp.ID, bytes.NewReader([]byte{}))

	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+bearerToken)

	res, err := client.Do(req)

	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {

		err = errors.New("Payment execution failed")
		return
	}

	var executedData map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(executedData)

	if err != nil {
		return
	}

	/* 	if executedData["intent"] != "sale" || executedData["state"] != "approved" {

		err = errors.New("Payment execution failed")
		return
	} */

	wcp.ID = executedData["id"].(string)

	return

}
