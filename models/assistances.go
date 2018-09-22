package models

import "time"

//Assitances Model
type Assitances struct {
	ID   int `orm:"column(id);pk" json:"id"`
	Type string
	//TODO:	worker_id
	//TODO:	watchers_id
	CreatedAt time.Time `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt time.Time `orm:"column(deleted_at);type(datetime);null" json:"-"`
}
