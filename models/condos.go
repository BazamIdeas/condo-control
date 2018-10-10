package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//Condos Model
type Condos struct {
	ID                int         `orm:"column(id);pk" json:"id"`
	Name              string      `orm:"column(name);size(255)" json:"name,omitempty" valid:"Required"`
	UserLimit         int         `orm:"column(user_limit);size(255)" json:"user_limit,omitempty" valid:"Required"`
	ZoneLimit         int         `orm:"column(zone_limit);size(255)" json:"zone_limit,omitempty" valid:"Required"`
	HourValue         float32     `orm:"colum(hour_value);size(20)" json:"hour_value,omitempty" valid:"Required"`
	ExtraHourIncrease float32     `orm:"colum(extra_hour_increase);size(20)" json:"extra_hour_increase,omitempty" valid:"Required"`
	WorkingHours      int         `orm:"column(working_hours)" json:"working_hours,omitempty" valid:"Required"`
	AssistancesMod    bool        `orm:"column(assistances_mod)" json:"assistances_mod"`
	RoutesMod         bool        `orm:"column(routes_mod)" json:"routes_mod"`
	Zones             []*Zones    `orm:"reverse(many)" json:"zone,omitempty"`
	Workers           []*Workers  `orm:"reverse(many)" json:"workers,omitempty"`
	Holidays          []*Holidays `orm:"reverse(many)" json:"holidays,omitempty"`
	CreatedAt         time.Time   `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt         time.Time   `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt         time.Time   `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName =
func (t *Condos) TableName() string {
	return "condos"
}

func (t *Condos) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"Zones", "Workers", "Holidays"}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddCondos insert a new Condos into database and returns
// last inserted Id on success.
func AddCondos(m *Condos) (id int64, err error) {
	o := orm.NewOrm()
	//m.Slug = GenerateSlug(m.TableName(), m.Name)
	id, err = o.Insert(m)
	return
}

// GetCondosByID retrieves Condos by Id. Returns error if
// Id doesn't exist
func GetCondosByID(id int) (v *Condos, err error) {
	v = &Condos{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	for _, zone := range v.Zones {
		zone.loadRelations()
	}

	return
}

// GetAllCondos retrieves all Condos matches certain condition. Returns empty list if
// no records exist
func GetAllCondos(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Condos))
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

	var l []Condos
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

// UpdateCondosByID updates Condos by Id and returns error if
// the record to be updated doesn't exist
func UpdateCondosByID(m *Condos) (err error) {
	o := orm.NewOrm()
	v := Condos{ID: m.ID}
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

// DeleteCondos deletes Condos by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCondos(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Condos{ID: id}
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

//GetCondosFromTrash return Condos soft Deleted
func GetCondosFromTrash() (condos []*Condos, err error) {

	o := orm.NewOrm()

	var v []*Condos

	_, err = o.QueryTable("condos").Filter("deleted_at__isnull", false).RelatedSel().All(&v)

	if err != nil {
		return
	}

	condos = v

	return

}
