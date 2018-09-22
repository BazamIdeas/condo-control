package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/globalsign/mgo/bson"

	"github.com/gosimple/slug"
)

//Services Model
type Services struct {
	ID         int           `orm:"column(id);auto" json:"id"`
	Name       string        `orm:"column(name);size(255)" json:"name,omitempty" valid:"Required"`
	Percertage float32       `orm:"column(percertage)" json:"percentage,omitempty" valid:"Required"`
	Prices     []*Prices     `orm:"reverse(many)" json:"prices,omitempty"`
	Price      *Prices       `orm:"-" json:"price,omitempty"`
	Slug       string        `orm:"column(slug);size(255)" json:"slug,omitempty" valid:"AlphaDash"`
	Code       string        `orm:"column(code);size(255)" json:"-"`
	Portfolios []*Portfolios `orm:"reverse(many)" json:"portfolios,omitempty"`
	Quantity   int           `orm:"-" bson:"quantity"     json:"quantity,omitempty"`
	CreatedAt  time.Time     `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt  time.Time     `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt  time.Time     `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName define Name
func (m *Services) TableName() string {
	return "services"
}

func (m *Services) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"Portfolios", "Prices"}

	for _, relation := range relations {
		o.LoadRelated(m, relation)
	}

	return

}

//AddServices insert a new Services into database and returns last inserted Id on success.
func AddServices(m *Services) (id int64, err error) {
	o := orm.NewOrm()

	m.Code = bson.NewObjectId().String()

	m.Slug = GenerateSlug(m.TableName(), m.Name)

	id, err = o.Insert(m)
	return
}

//GetServicesByID retrieves Services by Id. Returns error if Id doesn't exist
func GetServicesByID(id int) (v *Services, err error) {
	v = &Services{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

//GetAllServices retrieves all Services matches certain condition. Returns empty list if no records exist
func GetAllServices(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Services))
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

	var l []Services
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).Filter("deleted_at__isnull", true).RelatedSel().All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				v.loadRelations()
				v.Portfolios = nil
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
				v.Portfolios = nil
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

//UpdateServicesByID updates Services by Id and returns error if the record to be updated doesn't exist
func UpdateServicesByID(m *Services) (err error) {
	o := orm.NewOrm()
	v := Services{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteServices deletes Services by Id and returns error if
// the record to be deleted doesn't exist
func DeleteServices(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Services{ID: id}
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

//AddDefaultDataServices on init app
func AddDefaultDataServices() (result int64, err error) {

	o := orm.NewOrm()

	dummyData := []*Services{
		{
			Name:       "Logo a Medida",
			Percertage: 10.0,
			Code:       "01",
		},
	}

	for _, dummyService := range dummyData {
		dummyService.Slug = slug.Make(dummyService.Name)
	}

	result, err = o.InsertMulti(100, dummyData)

	return
}

//GetPricesServicesByID retrieves Services by Id. Returns error if Id doesn't exist
func (m *Services) GetPricesServicesByID(iso string) (err error) {
	o := orm.NewOrm()

	err = o.Read(m)

	if err != nil {
		return
	}

	//Get countries by Iso
	country, err := GetCountriesByIso(iso)

	if err != nil {
		return
	}

	price := &Prices{Currency: country.Currency, Service: m}

	err = o.Read(price, "Currency", "Service")

	if err != nil {
		return
	}

	err = searchFK(price.TableName(), price.ID).One(price)

	price.Service = nil
	m.Price = price

	return
}
