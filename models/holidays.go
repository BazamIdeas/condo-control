package models

import "time"

//Holidays Model
type Holidays struct {
	ID        int       `orm:"column(id);pk" json:"id"`
	Date      time.Time `orm:"column(user_limit);type(date)" json:"date,omitempty" valid:"Required"`
	Condo     *Condos   `orm:"rel(fk)" json:"condos"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt time.Time `orm:"column(deleted_at);type(datetime);null" json:"-"`
}
