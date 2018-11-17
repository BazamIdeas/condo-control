package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/vjeantet/jodaTime"
)

// Deliveries Model
type Deliveries struct {
	ID        int       `orm:"column(id);pk" json:"id"`
	Name      string    `orm:"column(name);" json:"name,omitempty" valid:"Required"`
	Approved  bool      `orm:"column(approved);" json:"approved"`
	Date      string    `orm:"column(date);auto_now_add;type(datetime);" json:"date,omitempty"`
	Worker    *Workers  `orm:"column(worker_id);rel(fk);" json:"worker,omitempty"`
	Items     []*Items  `orm:"reverse(many);" json:"items,omitempty"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt time.Time `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName =
func (t *Deliveries) TableName() string {
	return "deliveries"
}

func (t *Deliveries) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"Items"}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddDeliveries insert a new Tasks into database and returns
// last inserted Id on success.
func AddDeliveries(m *Deliveries) (id int64, err error) {
	o := orm.NewOrm()

	now := jodaTime.Format("Y-M-d HH:mm:ss", time.Now())
	m.Date = now

	id, err = o.Insert(m)

	if err != nil {
		return
	}

	m.ID = int(id)

	return
}

// GetDeliveriesByID retrieves deliveries by Id. Returns error if
// Id doesn't exist
func GetDeliveriesByID(id int) (v *Deliveries, err error) {
	v = &Deliveries{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// UpdateDeliveriesByID updates deliveries by Id and returns error if
// the record to be updated doesn't exist
func UpdateDeliveriesByID(m *Deliveries) (err error) {
	o := orm.NewOrm()
	v := Deliveries{ID: m.ID}
	// ascertain id exists in the database
	err = o.Read(&v)
	if err != nil {
		return
	}

	m.Approved = v.Approved
	m.Date = v.Date

	_, err = o.Update(m)

	if err != nil {
		return
	}

	return
}

// DeleteDeliveries deletes Deliveries by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDeliveries(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Deliveries{ID: id}
	// ascertain id exists in the database
	err = o.Read(&v)

	if err != nil {
		return
	}

	if trash {
		_, err = o.Delete(&v)
	} else {
		v.DeletedAt = time.Now()
		_, err = o.Update(&v)
	}

	if err != nil {
		return
	}

	return
}

//GetDeliveriesFromTrash return Points soft Deleted
func GetDeliveriesFromTrash() (deliveries []*Deliveries, err error) {

	o := orm.NewOrm()

	var v []*Deliveries

	_, err = o.QueryTable("deliveries").Filter("deleted_at__isnull", false).All(&v)

	if err != nil {
		return
	}

	deliveries = v

	return

}

//GetDeliveriesByWorkersID ...
func GetDeliveriesByWorkersID(workerID int) (deliveries []*Deliveries, err error) {

	o := orm.NewOrm()

	var v []*Deliveries

	_, err = o.QueryTable("deliveries").Filter("worker_id", workerID).Filter("deleted_at__isnull", true).RelatedSel().All(&v)

	if err != nil {
		return
	}

	for _, delivery := range v {
		delivery.loadRelations()
	}

	deliveries = v

	return

}
