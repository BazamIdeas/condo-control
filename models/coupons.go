package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//Coupons Model
type Coupons struct {
	ID         int       `orm:"column(id);auto" json:"id"`
	Percentage float32   `orm:"column(percentage)" json:"percentage,omitempty" valid:"Required"`
	Code       string    `orm:"column(code);size(45)" json:"code,omitempty" valid:"Required; AlphaNumeric"`
	Status     int8      `orm:"column(status);null" json:"status,omitempty" valid:"Required"`
	Orders     []*Orders `orm:"reverse(many)" json:"orders,omitempty"`
	CreatedAt  time.Time `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt  time.Time `orm:"column(updated_at);type(datetime);null;auto_now_add" json:"-"`
	DeletedAt  time.Time `orm:"column(deleted_at);type(datetime);null"  json:"-"`
}

//TableName =
func (t *Coupons) TableName() string {
	return "coupons"
}

func (t *Coupons) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"Orders"}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddCoupons insert a new Coupons into database and returns
// last inserted Id on success.
func AddCoupons(m *Coupons) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCouponsByID retrieves Coupons by Id. Returns error if
// Id doesn't exist
func GetCouponsByID(id int) (v *Coupons, err error) {
	v = &Coupons{ID: id}
	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// GetCouponByCode retrieves Coupon by Code. Returns error if Id doesn't exist
func GetCouponByCode(Code string) (v *Coupons, err error) {
	o := orm.NewOrm()

	v = &Coupons{Code: Code}
	err = o.Read(v, "code")

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// GetAllCoupons retrieves all Coupons matches certain condition. Returns empty list if
// no records exist
func GetAllCoupons(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Coupons))
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

	var l []Coupons
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

// UpdateCouponsByID updates Coupons by Id and returns error if
// the record to be updated doesn't exist
func UpdateCouponsByID(m *Coupons) (err error) {
	o := orm.NewOrm()
	v := Coupons{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCoupons deletes Coupons by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCoupons(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Coupons{ID: id}
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
