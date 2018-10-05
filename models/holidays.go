package models

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/vjeantet/jodaTime"
)

//Holidays Model
type Holidays struct {
	ID        int       `orm:"column(id);pk" json:"id"`
	Date      string    `orm:"column(date);type(date)" json:"date,omitempty" valid:"Required"`
	Condo     *Condos   `orm:"rel(fk);column(condos_id)" json:"condos,omitempty"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt time.Time `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName =
func (t *Holidays) TableName() string {
	return "holidays"
}

func (t *Holidays) loadRelations() {

	o := orm.NewOrm()

	relations := []string{}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddHolidays insert a new Holidays into database and returns
// last inserted Id on success.
func AddHolidays(m *Holidays) (id int64, err error) {
	o := orm.NewOrm()
	//m.Slug = GenerateSlug(m.TableName(), m.Name)
	id, err = o.Insert(m)

	if err != nil {
		return
	}
	m.ID = int(id)
	return
}

// GetHolidaysByID retrieves Holidays by Id. Returns error if
// Id doesn't exist
func GetHolidaysByID(id int) (v *Holidays, err error) {
	v = &Holidays{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// GetAllHolidays retrieves all Holidays matches certain condition. Returns empty list if
// no records exist
func GetAllHolidays(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Holidays))
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

	var l []Holidays
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

// UpdateHolidaysByID updates Holidays by Id and returns error if
// the record to be updated doesn't exist
func UpdateHolidaysByID(m *Holidays) (err error) {
	o := orm.NewOrm()
	v := Holidays{ID: m.ID}
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

// DeleteHolidays deletes Holidays by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHolidays(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Holidays{ID: id}
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

//GetHolidaysFromTrash return Holidays soft Deleted
func GetHolidaysFromTrash() (holidays []*Holidays, err error) {

	o := orm.NewOrm()

	var v []*Holidays

	_, err = o.QueryTable("holidays").Filter("deleted_at__isnull", false).All(&v)

	if err != nil {
		return
	}

	holidays = v

	return

}

func ExistHolidaysByCondoID(date string, condosID int) (ok bool, err error) {

	o := orm.NewOrm()

	count, err := o.QueryTable("holidays").Filter("condos_id", condosID).Filter("date", date).Count()

	if err != nil {
		return
	}

	if count > 0 {
		ok = true
	}

	return
}

func GetHolidaysByCondosID(year int, month time.Month, condosID int) (holidays map[string]*Holidays, err error) {

	holidays = map[string]*Holidays{}

	qb, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		return
	}

	monthTarget := time.Date(year, month, 1, 1, 1, 1, 1, time.UTC)
	monthTargetString := jodaTime.Format("Y-M-d", monthTarget)

	qb.Select("holidays.id", "holidays.date, DAY(holidays.date) as day").From("holidays").Where("holidays.condos_id = ?").And("YEAR(holidays.date) = YEAR(?)").And("MONTH(holidays.date) = MONTH(?)").OrderBy("holidays.date").Asc()

	sql := qb.String()

	o := orm.NewOrm()

	var paramMaps []orm.Params

	_, err = o.Raw(sql).SetArgs(condosID, monthTargetString, monthTargetString).Values(&paramMaps)

	if err != nil && err != orm.ErrNoRows {
		return
	}

	if err == orm.ErrNoRows {
		err = nil
		return
	}

	for _, holiday := range paramMaps {

		id, _ := strconv.Atoi(holiday["id"].(string))
		day := holiday["day"].(string)
		date := holiday["date"].(string)

		holidays[day] = &Holidays{ID: id, Date: date}
	}

	return
}
