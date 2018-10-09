package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//Verifications Model
type Verifications struct {
	ID        int       `orm:"column(id);pk" json:"id,omitempty"`
	Date      string    `orm:"column(date);type(datetime);" json:"date,omitempty" valid:"Required"`
	Latitude  float32   `orm:"column(latitude)" json:"latitude" valid:"Required"`
	Longitude float32   `orm:"column(longitude)" json:"longitude" valid:"Required"`
	Watcher   *Watchers `orm:"rel(fk);column(watchers_id)" json:"watchers,omitempty"`
	Point     *Points   `orm:"rel(fk);column(points_id)" json:"point,omitempty"`

	CreatedAt time.Time `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt time.Time `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName =
func (t *Verifications) TableName() string {
	return "verifications"
}

func (t *Verifications) loadRelations() {

	o := orm.NewOrm()

	relations := []string{}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddVerifications insert a new Verifications into database and returns
// last inserted Id on success.
func AddVerifications(m *Verifications) (id int64, err error) {
	o := orm.NewOrm()
	//m.Slug = GenerateSlug(m.TableName(), m.Name)
	id, err = o.Insert(m)
	return
}

// GetVerificationsByID retrieves Verifications by Id. Returns error if
// Id doesn't exist
func GetVerificationsByID(id int) (v *Verifications, err error) {
	v = &Verifications{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// GetAllVerifications retrieves all Verifications matches certain condition. Returns empty list if
// no records exist
func GetAllVerifications(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Verifications))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Verifications
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).Filter("deleted_at__isnull", true).RelatedSel().All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				v.loadRelations()
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				v.loadRelations()
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateVerificationsByID updates Verifications by Id and returns error if
// the record to be updated doesn't exist
func UpdateVerificationsByID(m *Verifications) (err error) {
	o := orm.NewOrm()
	v := Verifications{ID: m.ID}
	// ascertain id exists in the database
	err = o.Read(&v)
	if err != nil {
		return
	}

	var num int64

	num, err = o.Update(m)

	if err != nil {
		return
	}

	beego.Debug("Number of records updated in database:", num)

	return
}

// DeleteVerifications deletes Verifications by Id and returns error if
// the record to be deleted doesn't exist
func DeleteVerifications(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Verifications{ID: id}
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

//GetVerificationsFromTrash return Verifications soft Deleted
func GetVerificationsFromTrash() (verifications []*Verifications, err error) {

	o := orm.NewOrm()

	var v []*Verifications

	_, err = o.QueryTable("verifications").Filter("deleted_at__isnull", false).All(&v)

	if err != nil {
		return
	}

	verifications = v

	return

}

//AddManyVerifications ...
func AddManyVerifications(verifications []*Verifications) (err error) {

	o := orm.NewOrm()

	_, err = o.InsertMulti(200, &verifications)

	return

}
