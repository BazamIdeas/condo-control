package payments

import (
	"errors"
	"strings"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

//CreditCardStripe Generate a charge to a Credit Card
func (ccp *CreditCardPayment) CreditCardStripe() (err error) {

	//value on minimal unit, ej: "USD to cents"
	paymentAmount := ccp.Price * 100

	tax := ccp.Tax

	if tax > 0 {
		tax = (tax / 100) + 1
		paymentAmount += paymentAmount * tax
	}

	currency := strings.ToLower(ccp.Currency)

	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(int64(paymentAmount)),
		Currency:    stripe.String(currency),
		Description: stripe.String("XXXXXXXXXX"), //TODO: get services description
	}

	params.SetSource(ccp.Token)

	ch, err := charge.New(params)

	if stripeErr, ok := err.(*stripe.Error); ok {

		err = errors.New(stripeErr.Msg)
		return
	}

	ccp.ID = ch.ID

	return
}
