package models

import "time"

//Verifications Model
type Verifications struct {
	ID        int       `orm:"column(id);pk" json:"id"`
	Date      time.Time `orm:"column(date);type(datetime);null;auto_now_add" json:"date"`
	Watcher   *Watchers `orm:"rel(fk)" json:"watchers"`
	Zone      *Zones    `orm:"rel(fk)" json:"zones"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt time.Time `orm:"column(deleted_at);type(datetime);null" json:"-"`
}
