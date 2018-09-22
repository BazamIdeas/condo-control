package models

import (
	"github.com/globalsign/mgo/bson"
	"github.com/gofrs/uuid"
)

// Briefs model definiton.
type Briefs struct {
	ID     bson.ObjectId `orm:"-" bson:"_id,omitempty"      json:"id,omitempty"`
	Cookie string        `orm:"-" bson:"cookie"     json:"cookie,omitempty"`
	//ServiceForm *ServiceForms `orm:"-" bson:"service_form"     json:"service_form,omitempty" valid:"Required"`
	Client *Clients `orm:"-" bson:"client"     json:"client,omitempty"`
	Data   map[string]interface{}
}

//TableName define Name
func (m *Briefs) TableName() string {
	return "briefs"
}

// Insert a document to collection.
func (m *Briefs) Insert() (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	/* if m.Cookie != "" {
		err = m.GetBriefsByCookie(m.Cookie)

		if err == nil {
			return errors.New("El elemento ya existe")
		}
	} */

	UUID, err := uuid.NewV4()

	if err != nil {
		return err
	}

	m.Cookie = UUID.String()

	err = c.Insert(m)

	if err != nil {
		return err
	}

	return
}

// GetBriefsByID =
func (m *Briefs) GetBriefsByID(id string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	err = c.FindId(bson.ObjectIdHex(id)).One(m)

	if err != nil {
		return
	}

	return
}

// GetBriefsByCookie =
func (m *Briefs) GetBriefsByCookie(cookie string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	err = c.Find(bson.M{"cookie": cookie}).One(m)

	if err != nil {
		return
	}

	return
}

// GetAllBriefs =
func GetAllBriefs() (Briefs []*Briefs, err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C("briefs")

	err = c.Find(nil).All(&Briefs)

	if err != nil {
		return nil, err
	}

	return
}

// Update =
func (m *Briefs) Update() (err error) {
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
func (m *Briefs) Delete(id string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	err = c.RemoveId(bson.ObjectIdHex(id))

	if err != nil {
		return err
	}

	return
}
