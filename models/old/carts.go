package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/gofrs/uuid"
)

//Carts Model
type Carts struct {
	ID           int         `orm:"column(id);auto" json:"id"`
	Cookie       string      `orm:"column(cookie);size(255)" json:"cookie,omitempty"`
	Services     []*Services `orm:"rel(m2m)" json:"services,omitempty"`
	InitialValue float32     `orm:"-" json:"initial_value,omitempty"`
	FinalValue   float32     `orm:"-" json:"final_value,omitempty"`
	Currency     *Currencies `orm:"-" json:"currency,omitempty"`
	CreatedAt    time.Time   `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt    time.Time   `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt    time.Time   `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//CartsServices ...
type CartsServices struct {
	CartsID    int `orm:"column(carts_id);auto" json:"carts_id"`
	ServicesID int `orm:"column(services_id);auto" json:"services_id"`
	Quantity   int `orm:"column(quantity);auto" json:"quantity"`
}

//TableName define Name
func (t *Carts) TableName() string {
	return "carts"
}

func (t *Carts) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"Services"}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddCarts insert a new Carts into database and returns last inserted Id on success.
func AddCarts(m *Carts) (id int64, err error) {
	o := orm.NewOrm()

	UUID, err := uuid.NewV4()

	if err != nil {
		return 0, err
	}

	m.Cookie = UUID.String()

	id, err = o.Insert(m)
	return
}

// GetOrCreateCartsByCookie retrieves Carts by Cookie. Returns error if Id doesn't exist
func GetOrCreateCartsByCookie(Cookie string, Iso string) (v *Carts, err error) {
	o := orm.NewOrm()

	v = &Carts{Cookie: Cookie}
	result, id, err := o.ReadOrCreate(v, "cookie")

	if err != nil {
		return nil, err
	}

	//beego.Debug(result, id, err)

	if result {
		UUID, err := uuid.NewV4()

		if err != nil {
			return nil, err
		}

		v.Cookie = UUID.String()
		v.ID = int(id)

		UpdateCartsByID(v)
	}

	v.loadRelations()

	country, err := GetCountriesByIso(Iso)

	if err != nil {
		return nil, err
	}

	v.Currency = country.Currency

	for i, service := range v.Services {

		// Get Quantity
		var result = &CartsServices{}

		err := o.Raw("SELECT * FROM carts_servicess WHERE carts_id = ? AND services_id = ?").SetArgs(v.ID, service.ID).QueryRow(result)

		if err != nil {
			return nil, err
		}

		//Get the price
		price := &Prices{Currency: country.Currency, Service: service}

		err = o.Read(price, "Currency", "Service")

		if err != nil {
			return nil, err
		}

		serviceQuantity := float32(result.Quantity)
		priceWithQuantity := (price.Value * serviceQuantity)
		v.InitialValue += (priceWithQuantity * service.Percertage) / 100
		v.FinalValue += priceWithQuantity

		price.Service = nil
		price.Currency = nil

		v.Services[i].Price = price
		v.Services[i].Quantity = result.Quantity
	}

	return
}

// GetCartsByID retrieves Carts by Id. Returns error if Id doesn't exist
func GetCartsByID(id int) (v *Carts, err error) {
	v = &Carts{ID: id}
	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// GetAllCarts retrieves all Carts matches certain condition. Returns empty list if
// no records exist
func GetAllCarts(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Carts))
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

	var l []Carts
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

//UpdateCartsByID updates Carts by Id and returns error if the record to be updated doesn't exist
func UpdateCartsByID(m *Carts) (err error) {
	o := orm.NewOrm()
	v := Carts{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCarts deletes Carts by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCarts(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Carts{ID: id}
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
