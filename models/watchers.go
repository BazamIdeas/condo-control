package models

import "time"

//Watchers Model
type Watchers struct {
	ID            int              `orm:"column(id);pk" json:"id"`
	Email         string           `orm:"column(email);size(255)" json:"email,omitempty" valid:"Required"`
	Password      string           `orm:"column(password);" json:"password,omitempty" valid:"Required"`
	Phone         string           `orm:"column(phone);" json:"phone,omitempty" valid:"Required"`
	Worker        *Workers         `orm:"rel(fk);" json:"worker,omitempty"`
	Verifications []*Verifications `orm:"reverse(many);" json:"verifications,omitempty"`
	CreatedAt     time.Time        `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt     time.Time        `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt     time.Time        `orm:"column(deleted_at);type(datetime);null" json:"-"`
}
