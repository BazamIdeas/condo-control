package models

import "time"

//Workers Model
type Workers struct {
	ID         int           `orm:"column(id);pk" json:"id"`
	FirstName  string        `orm:"column(first_name);size(255)" json:"first_name,omitempty" valid:"Required"`
	LastName   string        `orm:"column(last_name);" json:"last_name,omitempty" valid:"Required"`
	Condo      *Condos       `orm:"rel(fk)" json:"condos"`
	Assitances []*Assitances `orm:"reverse(many)" json:"assistances,omitempty"`
	CreatedAt  time.Time     `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt  time.Time     `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt  time.Time     `orm:"column(deleted_at);type(datetime);null" json:"-"`
}
