package models

import (
	"errors"

	"github.com/globalsign/mgo/bson"
)

//Methods ...
type Methods []map[string]interface{}

// PaymentsMethods model definiton.
type PaymentsMethods struct {
	ID      bson.ObjectId `orm:"-" bson:"_id,omitempty"      json:"id,omitempty"`
	Country *Countries    `orm:"-" bson:"country"     json:"country,omitempty" valid:"Required"`
	Gateway *Gateways     `orm:"-" bson:"gateway"     json:"gateway,omitempty" valid:"Required"`
	Methods *Methods      `orm:"-" bson:"methods"     json:"methods,omitempty" valid:"Required"`
}

//TableName define Name
func (m *PaymentsMethods) TableName() string {
	return "payments_methods"
}

// Insert a document to collection.
func (m *PaymentsMethods) Insert() (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	err = m.GetByIsoAndGateway(m.Country.Iso, m.Gateway.Name)

	if err == nil {
		return errors.New("La relacion ya existe")
	}

	err = c.Insert(m)

	if err != nil {
		return err
	}

	return
}

// GetPaymentsMethodsByID =
func (m *PaymentsMethods) GetPaymentsMethodsByID(id string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	err = c.FindId(bson.ObjectIdHex(id)).One(m)

	if err != nil {
		return
	}

	return
}

// GetByIsoAndGateway =
func (m *PaymentsMethods) GetByIsoAndGateway(CountryIso string, GatewayName string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	err = c.Find(bson.M{"country.iso": CountryIso, "gateway.name": GatewayName}).One(m)

	if err != nil {
		return
	}

	return
}

// GetAllPaymentsMethods =
func (m *PaymentsMethods) GetAllPaymentsMethods() (PaymentsMethods []*PaymentsMethods, err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	err = c.Find(nil).All(&PaymentsMethods)

	if err != nil {
		return nil, err
	}

	return
}

// Update =
func (m *PaymentsMethods) Update() (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	err = c.Update(bson.M{"_id": m.ID}, m)

	if err != nil {
		return err
	}

	return
}

// Delete =
func (m *PaymentsMethods) Delete(id string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	err = c.RemoveId(bson.ObjectIdHex(id))

	if err != nil {
		return err
	}

	return
}

/* func AddDefaultDataMethods() {
	var methods []*Methods

	var eur []

	for _, method := range methods {
		Insert()
	}
}
*/
