package models

import (
	"errors"

	"github.com/globalsign/mgo/bson"
)

// User model definiton.
type User struct {
	ID       bson.ObjectId `orm:"-" bson:"_id,omitempty"      json:"id,omitempty"`
	Name     string        `orm:"-" bson:"name"     json:"name,omitempty" valid:"Required"`
	Email    string        `orm:"-" bson:"email"     json:"email,omitempty" valid:"Required"`
	Password string        `orm:"-" bson:"password" json:"password,omitempty" valid:"Required"`
	Token    string        `orm:"-" json:"token,omitempty"`
}

//TableName define Name
func (u *User) TableName() string {
	return "users"
}

// Insert a document to collection.
func (u *User) Insert() (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(u.TableName())

	u.Password = GetMD5Hash(u.Password)

	err = u.GetUsersByEmail()

	if err == nil {
		return errors.New("El usuario ya existe")
	}

	err = c.Insert(u)

	if err != nil {
		return err
	}

	u.Password = ""

	return
}

// GetUsersByID =
func (u *User) GetUsersByID(id string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C("users")

	err = c.FindId(bson.ObjectIdHex(id)).One(u)

	if err != nil {
		return
	}

	u.Password = ""

	return
}

//GetUsersByEmail ...
func (u *User) GetUsersByEmail() (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C("users")

	err = c.Find(bson.M{"email": u.Email}).One(u)

	if err != nil {
		return
	}

	u.Password = ""

	return
}

// ChangePassword  =
func (u *User) ChangePassword() (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(u.TableName())

	u.Password = GetMD5Hash(u.Password)

	_, err = c.UpsertId(bson.M{"email": u.Email}, u)

	if err != nil {
		return
	}

	u.Password = ""

	return
}

// GetAllUsers =
func (u *User) GetAllUsers() (users []*User, err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(u.TableName())

	u.Password = GetMD5Hash(u.Password)

	err = c.Find(nil).All(&users)

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		user.Password = ""
	}

	return
}

// Update =
func (u *User) Update() (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(u.TableName())

	u.Password = GetMD5Hash(u.Password)

	err = c.Update(bson.M{"_id": u.ID}, u)

	if err != nil {
		return err
	}

	return
}

// Delete =
func (u *User) Delete(id string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(u.TableName())

	u.Password = GetMD5Hash(u.Password)

	err = c.RemoveId(bson.ObjectIdHex(id))

	if err != nil {
		return err
	}

	return
}

// LoginUsers =
func (u *User) LoginUsers() (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("liderlogo_site").C(u.TableName())

	u.Password = GetMD5Hash(u.Password)

	err = c.Find(bson.M{"email": u.Email, "password": u.Password}).One(u)

	if err != nil {
		return
	}

	u.Password = ""

	return
}

//AddDefaultDataUsers ...
func AddDefaultDataUsers() (id *bson.ObjectId, err error) {
	u := User{ID: bson.NewObjectId(), Email: "admin@logo.pro", Password: "123456"}

	err = u.Insert()

	if err != nil {
		return nil, err
	}

	return &u.ID, nil
}
