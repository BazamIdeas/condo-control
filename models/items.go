package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Items ...
type Items struct {
	ID          int         `orm:"column(id);pk" json:"id"`
	Address     string      `orm:"column(address);" json:"address,omitempty" valid:"Required"`
	Description string      `orm:"column(description);" json:"description,omitempty" valid:"Required"`
	Delivered   bool        `orm:"column(delivered);" json:"delivered" `
	DateEnd     string      `orm:"column(date_end);type(string);" json:"date_end,omitempty"`
	Comment     string      `orm:"column(comment);null" json:"comment,omitempty"`
	ImageUUID   string      `orm:"column(image_uuid);null" json:"image_uuid,omitempty"`
	ImageMime   string      `orm:"column(image_mime);null" json:"image_mime,omitempty"`
	Delivery    *Deliveries `orm:"rel(fk);column(delivery_id)" json:"delivery,omitempty"`
	CreatedAt   time.Time   `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt   time.Time   `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt   time.Time   `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName =
func (t *Items) TableName() string {
	return "items"
}

func (t *Items) loadRelations() {

	o := orm.NewOrm()

	relations := []string{}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddItems insert a new Tasks into database and returns
// last inserted Id on success.
func AddItems(m *Items) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(m)

	if err != nil {
		return
	}

	m.ID = int(id)

	return
}

// GetItemsByID retrieves Items by Id. Returns error if
// Id doesn't exist
func GetItemsByID(id int) (v *Items, err error) {
	v = &Items{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// UpdateItemsByID updates Items by Id and returns error if
// the record to be updated doesn't exist
func UpdateItemsByID(m *Items, ignoreStatus bool) (err error) {
	o := orm.NewOrm()
	v := Items{ID: m.ID}
	// ascertain id exists in the database
	err = o.Read(&v)
	if err != nil {
		return
	}

	m.DateEnd = v.DateEnd

	if ignoreStatus {
		m.Delivered = v.Delivered
	}

	_, err = o.Update(m)

	if err != nil {
		return
	}

	return
}

// DeleteItems deletes Items by Id and returns error if
// the record to be deleted doesn't exist
func DeleteItems(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Items{ID: id}
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

//GetItemsFromTrash return Points soft Deleted
func GetItemsFromTrash() (items []*Items, err error) {

	o := orm.NewOrm()

	var v []*Items

	_, err = o.QueryTable("items").Filter("deleted_at__isnull", false).All(&v)

	if err != nil {
		return
	}

	items = v

	return

}
