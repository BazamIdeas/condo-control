package models

import (
	"errors"
	"fmt"

	"github.com/globalsign/mgo/bson"
)

//Sections ...
type Sections []map[string]interface{}

// ServiceForms model definiton.
type ServiceForms struct {
	ID          bson.ObjectId `orm:"-" bson:"_id,omitempty"      json:"id,omitempty"`
	Title       string        `orm:"-" bson:"title"     json:"title,omitempty" valid:"Required"`
	Subtitle    string        `orm:"-" bson:"subtitle"     json:"subtitle,omitempty" valid:"Required"`
	Description string        `orm:"-" bson:"description"     json:"description,omitempty" valid:"Required"`
	Service     *Services     `orm:"-" bson:"service"     json:"service,omitempty" valid:"Required"`
	Sections    *Sections     `orm:"-" bson:"sections"     json:"sections,omitempty" valid:"Required"`
}

//TableName define Name
func (m *ServiceForms) TableName() string {
	return "service_forms"
}

// Insert a document to collection.
func (m *ServiceForms) Insert() (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	fmt.Println(m.Service.Slug)

	err = m.GetServiceFormsByServiceSlug(m.Service.Slug)

	if err == nil {
		return errors.New("El elemento ya existe")
	}

	err = c.Insert(m)

	if err != nil {
		return err
	}

	return
}

// GetServiceFormsByID =
func (m *ServiceForms) GetServiceFormsByID(id string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	err = c.FindId(bson.ObjectIdHex(id)).One(m)

	if err != nil {
		return
	}

	return
}

// GetServiceFormsByServiceSlug =
func (m *ServiceForms) GetServiceFormsByServiceSlug(slug string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	err = c.Find(bson.M{"service.slug": slug}).One(m)

	if err != nil {
		return
	}

	return
}

// GetAllServiceForms =
func (m *ServiceForms) GetAllServiceForms() (ServiceForms []*ServiceForms, err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	err = c.Find(nil).All(&ServiceForms)

	if err != nil {
		return nil, err
	}

	return
}

// Update =
func (m *ServiceForms) Update() (err error) {
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
func (m *ServiceForms) Delete(id string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	err = c.RemoveId(bson.ObjectIdHex(id))

	if err != nil {
		return err
	}

	return
}
