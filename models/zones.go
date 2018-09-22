package models

import "time"

//Zones Model
type Zones struct {
	ID            int              `orm:"column(id);pk" json:"id"`
	Name          string           `orm:"column(name);size(255)" json:"name,omitempty" valid:"Required"`
	Condo         *Condos          `orm:"rel(fk)" json:"condos"`
	Verifications []*Verifications `orm:"reverse(many);" json:"verifications,omitempty"`
	CreatedAt     time.Time        `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt     time.Time        `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt     time.Time        `orm:"column(deleted_at);type(datetime);null" json:"-"`
}
