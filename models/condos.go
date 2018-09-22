package models

import "time"

//Condos Model
type Condos struct {
	ID                int         `orm:"column(id);pk" json:"id"`
	Name              string      `orm:"column(name);size(255)" json:"name,omitempty" valid:"Required"`
	UsersLimit        int         `orm:"column(users_limit);size(255)" json:"users_limit,omitempty" valid:"Required"`
	ZonesLimit        int         `orm:"column(zones_limit);size(255)" json:"zones_limit,omitempty" valid:"Required"`
	HourValue         float32     `orm:"colum(hour_value);size(20)" valid:"Required"`
	ExtraHourIncrease float32     `orm:"colum(extra_hour_increase);size(20)" valid:"Required"`
	WorkingHours      int         `orm:"column(working_hours)" json:"working_hours" valid:"Required"`
	Zones             []*Zones    `orm:"reverse(many)" json:"zones,omitempty"`
	Workers           []*Workers  `orm:"reverse(many)" json:"workers,omitempty"`
	Holidays          []*Holidays `orm:"reverse(many)" json:"holidays,omitempty"`
	CreatedAt         time.Time   `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt         time.Time   `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt         time.Time   `orm:"column(deleted_at);type(datetime);null" json:"-"`
}
