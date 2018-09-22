package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// Currencies Model
type Currencies struct {
	ID        int          `orm:"column(id);auto" json:"id"`
	Name      string       `orm:"column(name);size(255)" json:"name,omitempty" valid:"Required"`
	Iso       string       `orm:"column(iso);size(3)" json:"iso,omitempty" valid:"Required; Length(3); Alpha"`
	Symbol    string       `orm:"column(symbol);size(3)" json:"symbol,omitempty" valid:"Required"`
	Gateways  []*Gateways  `orm:"reverse(many)" json:"gateways,omitempty"`
	Countries []*Countries `orm:"reverse(many)" json:"countries,omitempty"`
	CreatedAt time.Time    `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt time.Time    `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt time.Time    `orm:"column(deleted_at);type(datetime);null"  json:"-"`
}

// TableName =
func (t *Currencies) TableName() string {
	return "currencies"
}

func (t *Currencies) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"Gateways", "Countries"}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddCurrencies insert a new Currencies into database and returns
// last inserted Id on success.
func AddCurrencies(m *Currencies) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCurrenciesByID retrieves Currencies by Id. Returns error if
// Id doesn't exist
func GetCurrenciesByID(id int) (v *Currencies, err error) {
	v = &Currencies{ID: id}
	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// GetAllCurrencies retrieves all Currencies matches certain condition. Returns empty list if
// no records exist
func GetAllCurrencies(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Currencies))
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

	var l []Currencies
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

// UpdateCurrenciesByID updates Currencies by Id and returns error if
// the record to be updated doesn't exist
func UpdateCurrenciesByID(m *Currencies) (err error) {
	o := orm.NewOrm()
	v := Currencies{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCurrencies deletes Currencies by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCurrencies(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Currencies{ID: id}
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

//AddDefaultDataCurrencies on init app
func AddDefaultDataCurrencies() (count int64, err error) {

	o := orm.NewOrm()

	dummyData := []*Currencies{
		{
			Symbol: "€",
			Name:   "Euro",
			Iso:    "EUR",
		},
		{
			Symbol: "$",
			Name:   "Dólar estadounidense",
			Iso:    "USD",
		},
	}

	count, err = o.InsertMulti(100, dummyData)

	return
}

func addRelationsGatewaysCurrencies() (count int64, errors []error) {

	o := orm.NewOrm()

	dummyData := map[string][]string{
		"01": /* Paypal */ {
			"USD",
		},
	}

	for key, dummyGateway := range dummyData {

		gateway := Gateways{Code: key}

		err := o.Read(&gateway, "code")

		if err != nil {
			continue
		}

		m2m := o.QueryM2M(&gateway, "Currencies")

		var InsertManyCurrencies []*Currencies

		for _, iso := range dummyGateway {

			currency := Currencies{Iso: iso}

			err := o.Read(&currency, "iso")

			if err != nil {
				continue
			}

			InsertManyCurrencies = append(InsertManyCurrencies, &currency)

		}

		result, err := m2m.Add(InsertManyCurrencies)

		if err != nil {
			errors = append(errors, err)
			continue
		}

		count += result

	}

	return
}
