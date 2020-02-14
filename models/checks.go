package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//Checks Model
type Checks struct {
	ID          int            `orm:"column(id);pk" json:"id,omitempty"`
	Date        string         `orm:"column(date);type(datetime);" json:"date,omitempty" valid:"Required"`
	Comment     string         `orm:"column(comment);null" json:"comment,omitempty"`
	Worker      *Workers       `orm:"rel(fk);column(workers_id)" json:"workers,omitempty"`
	Occurrences []*Occurrences `orm:"reverse(many);" json:"occurrences,omitempty"`
	CreatedAt   time.Time      `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt   time.Time      `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt   time.Time      `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName =
func (t *Checks) TableName() string {
	return "checks"
}

func (t *Checks) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"Occurrences"}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddChecks insert a new Checks into database and returns
// last inserted Id on success.
func AddChecks(m *Checks) (id int64, err error) {
	o := orm.NewOrm()
	//m.Slug = GenerateSlug(m.TableName(), m.Name)
	id, err = o.Insert(m)

	if err != nil {
		return
	}

	m.ID = int(id)

	return
}

// GetChecksByID retrieves Checks by Id. Returns error if
// Id doesn't exist
func GetChecksByID(id int) (v *Checks, err error) {
	v = &Checks{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// GetAllChecks retrieves all Checks matches certain condition. Returns empty list if
// no records exist
func GetAllChecks(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Checks))
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

	var l []Checks
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

// UpdateChecksByID updates Checks by Id and returns error if
// the record to be updated doesn't exist
func UpdateChecksByID(m *Checks) (err error) {
	o := orm.NewOrm()
	v := Checks{ID: m.ID}
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

// DeleteChecks deletes Checks by Id and returns error if
// the record to be deleted doesn't exist
func DeleteChecks(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Checks{ID: id}
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

//GetChecksFromTrash return Checks soft Deleted
func GetChecksFromTrash() (checks []*Checks, err error) {

	o := orm.NewOrm()

	var v []*Checks

	_, err = o.QueryTable("checks").Filter("deleted_at__isnull", false).All(&v)

	if err != nil {
		return
	}

	checks = v

	return

}
