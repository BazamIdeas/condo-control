package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//Watchers Model
type Watchers struct {
	ID            int              `orm:"column(id);pk" json:"id"`
	Email         string           `orm:"column(email);size(255)" json:"email,omitempty" valid:"Required"`
	Password      string           `orm:"column(password);" json:"password,omitempty" valid:"Required"`
	Phone         string           `orm:"column(phone);" json:"phone,omitempty" valid:"Required"`
	Token         string           `orm:"-" json:"token,omitempty"`
	Worker        *Workers         `orm:"rel(fk);column(workers_id)" json:"worker,omitempty"`
	Verifications []*Verifications `orm:"reverse(many);" json:"verifications,omitempty"`
	CreatedAt     time.Time        `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt     time.Time        `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt     time.Time        `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName =
func (t *Watchers) TableName() string {
	return "watchers"
}

func (t *Watchers) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"Verifications"}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddWatchers insert a new Watchers into database and returns
// last inserted Id on success.
func AddWatchers(m *Watchers) (id int64, err error) {

	o := orm.NewOrm()

	m.Password = GetMD5Hash(m.Password)

	id, err = o.Insert(m)

	if err != nil {
		return
	}
	m.ID = int(id)
	m.Password = ""

	return
}

// GetWatchersByID retrieves Watchers by Id. Returns error if
// Id doesn't exist
func GetWatchersByID(id int) (v *Watchers, err error) {
	v = &Watchers{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// LoginWatchers login a Watchers, returns
// if Exists.
func LoginWatchers(m *Watchers) (id int, err error) {
	o := orm.NewOrm()

	m.Password = GetMD5Hash(m.Password)

	err = o.QueryTable(m.TableName()).Filter("deleted_at__isnull", true).Filter("email", m.Email).Filter("password", m.Password).RelatedSel().One(m)

	if err != nil {
		return 0, err
	}

	m.Password = ""

	return m.ID, err
}

// GetAllWatchers retrieves all Watchers matches certain condition. Returns empty list if
// no records exist
func GetAllWatchers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Watchers))
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

	var l []Watchers
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

/*
func GetAllByCondosID (condosID int) (watchers []*Watchers, err error){

	GetCondosByID()

} */

// UpdateWatchersByID updates Watchers by Id and returns error if
// the record to be updated doesn't exist
func UpdateWatchersByID(m *Watchers) (err error) {
	o := orm.NewOrm()
	v := Watchers{ID: m.ID}
	// ascertain id exists in the database
	err = o.Read(&v)
	if err != nil {
		return
	}

	var num int64

	if m.Password != "" {
		m.Password = GetMD5Hash(m.Password)
	} else {
		m.Password = v.Password
	}

	num, err = o.Update(m)

	if err != nil {
		return
	}

	beego.Debug("Number of records updated in database:", num)

	return
}

// DeleteWatchers deletes Watchers by Id and returns error if
// the record to be deleted doesn't exist
func DeleteWatchers(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Watchers{ID: id}
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

//GetWatchersFromTrash return Watchers soft Deleted
func GetWatchersFromTrash() (watchers []*Watchers, err error) {

	o := orm.NewOrm()

	var v []*Watchers

	_, err = o.QueryTable("watchers").Filter("deleted_at__isnull", false).All(&v)

	if err != nil {
		return
	}

	watchers = v

	return

}

//GetWatchersByCondosID ...
func GetWatchersByCondosID(condosID int) (watchers []*Watchers, err error) {

	//condos, err := GetCondosByID(condosID)

	o := orm.NewOrm()

	v := []*Watchers{}

	_, err = o.QueryTable("watchers").Filter("deleted_at__isnull", true).RelatedSel().All(&v)

	if err != nil {
		return
	}

	for _, watcher := range v {

		if watcher.Worker.Condo.ID != condosID {
			continue
		}
		watchers = append(watchers, watcher)
	}

	return

}

//GetVerificationsByDate ...
func (t *Watchers) GetVerificationsByDate(date time.Time) (err error) {

	qb, _ := orm.NewQueryBuilder("mysql")

	var verifications []*Verifications

	qb.Select("verifications.id, verifications.date").From("verifications").Where("DATEDIFF(verifications.date, ?) = 0").And("verifications.watchers_id = ?").OrderBy("verifications.date").Asc()

	sql := qb.String()

	o := orm.NewOrm()

	_, err = o.Raw(sql, date.String(), t.ID).QueryRows(&verifications)

	if err != nil {
		return
	}

	if len(verifications) == 0 {
		err = orm.ErrNoRows
		return
	}

	for _, verification := range verifications {
		verification.loadRelations()

		searchFK(verification.TableName(), verification.ID).One(verification)
	}

	t.Verifications = verifications

	return

}
