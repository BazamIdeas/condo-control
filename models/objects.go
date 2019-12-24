package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Objects ...
type Objects struct {
	ID          int            `orm:"column(id);pk" json:"id"`
	Name        string         `orm:"column(name);" json:"name,omitempty" valid:"Required"`
	Code        string         `orm:"column(code);" json:"code,omitempty" valid:"Required"`
	Quantity    int            `orm:"column(quantity);" json:"quantity,omitempty" valid:"Required"`
	Condo       *Condos        `orm:"rel(fk);column(condos_id)" json:"condos,omitempty"`
	Occurrences []*Occurrences `orm:"reverse(many);" json:"occurrences,omitempty"`
	CreatedAt   time.Time      `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt   time.Time      `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt   time.Time      `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName =
func (t *Objects) TableName() string {
	return "objects"
}

func (t *Objects) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"Occurrences"}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddObjects insert a new Tasks into database and returns
// last inserted Id on success.
func AddObjects(m *Objects) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(m)

	if err != nil {
		return
	}

	m.ID = int(id)

	return
}

// GetObjectsByID retrieves Objects by Id. Returns error if
// Id doesn't exist
func GetObjectsByID(id int) (v *Objects, err error) {
	v = &Objects{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// UpdateObjectsByID updates Objects by Id and returns error if
// the record to be updated doesn't exist
func UpdateObjectsByID(m *Objects, ignoreStatus bool) (err error) {
	o := orm.NewOrm()
	v := Objects{ID: m.ID}
	// ascertain id exists in the database
	err = o.Read(&v)
	if err != nil {
		return
	}

	_, err = o.Update(m)

	if err != nil {
		return
	}

	return
}

// DeleteObjects deletes Objects by Id and returns error if
// the record to be deleted doesn't exist
func DeleteObjects(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Objects{ID: id}
	// ascertain id exists in the database
	err = o.Read(&v)

	if err != nil {
		return
	}

	if trash {
		_, err = o.Delete(&v)
	} else {
		v.DeletedAt = time.Now().In(orm.DefaultTimeLoc)
		_, err = o.Update(&v)
	}

	if err != nil {
		return
	}

	return
}

//GetObjectsFromTrash return Points soft Deleted
func GetObjectsFromTrash() (objects []*Objects, err error) {

	o := orm.NewOrm()

	var v []*Objects

	_, err = o.QueryTable("objects").Filter("deleted_at__isnull", false).All(&v)

	if err != nil {
		return
	}

	objects = v

	return

}
