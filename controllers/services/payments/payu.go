package payments

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

//PAYU API DATA

//payuCredentials define API credentials
type payuCredentials struct {
	APIKey   string `json:"apiKey"`
	APILogin string `json:"apiLogin"`
}

//payuCredentials define a Card (credit or debit)
type payuCard struct {
	Number         string `json:"number"`
	SecurityCode   string `json:"securityCode"`
	ExpirationDate string `json:"expirationDate"`
	Name           string `json:"name"`
}

//billingAddress define a Billing address
type payuBillingAddress struct {
	Street1 string `json:"street1"`
	Street2 string `json:"street2"`
	City    string `json:"city"`
	/* State string */
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"` //MEXICO
	/* Phone string */
}

//payuPayer define a Payer
type payuPayer struct {
	/* MerchantBuyerID string */
	FullName       string             `json:"fullName"`
	EmailAddress   string             `json:"emailAddress"`
	Contactphone   string             `json:"contactPhone"`
	DniNumber      string             `json:"dniNumber"`
	Birthdate      string             `json:"birthDate,omitempty"` //MEXICO
	BillingAddress payuBillingAddress `json:"billingAddress"`
}

//payuShipping define a Shipping Address
type payuShippingAddress struct {
	Street1 string `json:"street1"`
	Street2 string `json:"street2"`
	City    string `json:"city"`
	/* State      string */
	Country string `json:"country"`
	/* PostalCode string */
	/* Phone      string */
}

//payuBuyer define a Buyer Data
type payuBuyer struct {
	/* MerchantBuyerID string */
	FullName        string              `json:"fullName"`
	EmailAddress    string              `json:"emailAddress"`
	ContactPhone    string              `json:"contactPhone"`
	DniNumber       string              `json:"dniNumber"`
	ShippingAddress payuShippingAddress `json:"shippingAddress"`
}

//payuTaxValue define a Transaction Value Total - Only Colombia
type payuTaxValue struct {
	Value    float32 `json:"value"`
	Currency string  `json:"currency"`
}

type payuAdditionalValues struct {
	TaxValue payuTaxValue `json:"TX_VALUE"`
}

type payuOrder struct {
	AccountID        string               `json:"accountId"`
	ReferenceCode    string               `json:"referenceCode"`
	Description      string               `json:"description"`
	Language         string               `json:"language"`
	Signature        string               `json:"signature"`
	NotifyURL        string               `json:"notifyUrl"`
	AdditionalValues payuAdditionalValues `json:"addtionalValues,omitempty"`
	Buyer            payuBuyer            `json:"buyer"`
	/* ShippingAddres struct {
		Street1 string
		Street2 string
		City string
		State string
		Country string
		PostalCode string
		Phone string
	} */
}

type payuExtraParameters struct {
	InstallamentsNumber int `json:"INSTALLMENTS_NUMBER"`
}

//payuTransaction define a payU transaction
type payuTransaction struct {
	Order             payuOrder           `json:"order"`
	Payer             payuPayer           `json:"payer"`
	CreditCardTokenID string              `json:"creditCardTokenId,omitempty"`
	ExtraParameters   payuExtraParameters `json:"extraParameters"`
	/* CreditCard struct {
		Number         string
		SecurityCode   string
		ExpirationDate string
		Name           string
	} */
	DebitCard       payuCard `json:"debitCard,omitempty"`
	Type            string   `json:"type"`
	PaymentMethod   string   `json:"paymentMethod"`
	PaymentCountry  string   `json:"paymentCountry"`
	DeviceSessionID string   `json:"deviceSessionId"`
	IPAddress       string   `json:"ipAddress"`
	Cookie          string   `json:"cookie"`
	UserAgent       string   `json:"userAgent"`
}

type payuRequestData struct {
	Language    string          `json:"language"`
	Command     string          `json:"command"`
	Merchant    payuCredentials `json:"merchant"`
	Transaction payuTransaction `json:"transaction"`
	Test        bool            `json:"test"`
}

var payu = payuCredentials{
	APIKey:   "XXXXXXXXXX",
	APILogin: "XXXXXXXXXX",
}

//PayuCreditCard define a payment with credit card
func (ccp *CreditCardPayment) PayuCreditCard() (err error) {

	defer (func() {
		r := recover()

		if r != nil {
			err = errors.New("Bad entity data")
		}
	})()

	languageCode := "ES"

	payerData := ccp.ExtraData["payer"].(map[string]interface{})

	orderData := ccp.ExtraData["order"].(map[string]interface{})

	billingAddress := payerData["billing_address"].(payuBillingAddress)
	billingAddress.Country = ccp.Country

	shippingAddress := payuShippingAddress{
		Street1: billingAddress.Street1,
		Street2: billingAddress.Street1,
		City:    billingAddress.Street1,
		Country: billingAddress.Country,
	}

	//Buyer
	buyer := payuBuyer{
		FullName:        payerData["full_name"].(string),
		EmailAddress:    payerData["email"].(string),
		ContactPhone:    payerData["contact_phone"].(string),
		DniNumber:       payerData["dni_number"].(string),
		ShippingAddress: shippingAddress,
	}

	taxValue := payuTaxValue{
		Currency: ccp.Currency,
		Value:    ccp.Price,
	}

	additionalValues := payuAdditionalValues{TaxValue: taxValue}

	order := payuOrder{
		AccountID:        orderData["account_id"].(string),
		ReferenceCode:    orderData["reference_code"].(string),
		Description:      "", //TODO: description order by products
		Language:         languageCode,
		Signature:        orderData["signature"].(string),
		NotifyURL:        orderData["notify_url"].(string),
		AdditionalValues: additionalValues,
		Buyer:            buyer, // orderData["billing_address"].(payuBillingAddress)
	}

	////
	//Payer
	payer := payuPayer{
		FullName:       payerData["full_name"].(string),
		EmailAddress:   payerData["email"].(string),
		Contactphone:   payerData["contact_phone"].(string),
		DniNumber:      payerData["dni_number"].(string),
		Birthdate:      payerData["birthdate"].(string),
		BillingAddress: billingAddress,
	}

	extraParameters := payuExtraParameters{InstallamentsNumber: 1}

	transaction := payuTransaction{
		Order:             order,
		Payer:             payer,
		ExtraParameters:   extraParameters,
		CreditCardTokenID: ccp.Token,
		Type:              "AUTHORIZATION",
		PaymentMethod:     payerData["payment_method"].(string),
		PaymentCountry:    ccp.Country,
		DeviceSessionID:   payerData["device_session_id"].(string),
		IPAddress:         payerData["ip_address"].(string),
		Cookie:            payerData["cookie"].(string),
		UserAgent:         payerData["user_agent"].(string),
	}

	payuRequest := &payuRequestData{
		Language:    languageCode,
		Command:     "SUBMIT_TRANSACTION",
		Merchant:    payu,
		Transaction: transaction,
		Test:        false,
	}

	payuRequestJSON, err := json.Marshal(payuRequest)

	if err != nil {
		return
	}

	client := &http.Client{}

	response, err := client.Post("", "", bytes.NewReader(payuRequestJSON)) //TODO: agregar urls de payu

	if err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {

		err = errors.New("Payment Failed")

		return
	}

	return
}
