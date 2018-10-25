package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//Supervisors Model
type Supervisors struct {
	ID        int       `orm:"column(id);pk" json:"id"`
	Email     string    `orm:"column(email);size(255)" json:"email,omitempty" valid:"Required,Email"`
	Password  string    `orm:"column(password);" json:"password,omitempty" valid:"Required"`
	Phone     string    `orm:"column(phone);null" json:"phone,omitempty" valid:"Required"`
	Token     string    `orm:"-" json:"token,omitempty"`
	Worker    *Workers  `orm:"rel(fk);column(workers_id)" json:"worker,omitempty"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt time.Time `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName =
func (t *Supervisors) TableName() string {
	return "supervisors"
}

func (t *Supervisors) loadRelations() {

	o := orm.NewOrm()

	relations := []string{}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddSupervisors insert a new Supervisors into database and returns
// last inserted Id on success.
func AddSupervisors(m *Supervisors) (id int64, err error) {
	o := orm.NewOrm()
	//m.Slug = GenerateSlug(m.TableName(), m.Name)
	id, err = o.Insert(m)

	if err != nil {
		return
	}

	m.ID = int(id)

	return
}

// GetSupervisorsByID retrieves Supervisors by Id. Returns error if
// Id doesn't exist
func GetSupervisorsByID(id int) (v *Supervisors, err error) {
	v = &Supervisors{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// LoginSupervisors login a Supervisors, returns
// if Exists.
func LoginSupervisors(m *Supervisors) (id int, err error) {
	o := orm.NewOrm()

	m.Password = GetMD5Hash(m.Password)

	err = o.QueryTable(m.TableName()).Filter("deleted_at__isnull", true).Filter("email", m.Email).Filter("password", m.Password).RelatedSel().One(m)

	if err != nil {
		return 0, err
	}

	m.Password = ""

	return m.ID, err
}

// GetAllSupervisors retrieves all Supervisors matches certain condition. Returns empty list if
// no records exist
func GetAllSupervisors(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Supervisors))
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

	var l []Supervisors
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

// UpdateSupervisorsByID updates Supervisors by Id and returns error if
// the record to be updated doesn't exist
func UpdateSupervisorsByID(m *Supervisors) (err error) {
	o := orm.NewOrm()
	v := Supervisors{ID: m.ID}
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

// DeleteSupervisors deletes Supervisors by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSupervisors(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Supervisors{ID: id}
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

//GetSupervisorsFromTrash return Supervisors soft Deleted
func GetSupervisorsFromTrash() (supervisors []*Supervisors, err error) {

	o := orm.NewOrm()

	var v []*Supervisors

	_, err = o.QueryTable("supervisors").Filter("deleted_at__isnull", false).All(&v)

	if err != nil {
		return
	}

	supervisors = v

	return

}

//GetSupervisorsByCondosID return Supervisors soft Deleted
func GetSupervisorsByCondosID(condosID int) (supervisors []*Supervisors, err error) {

	o := orm.NewOrm()

	var (
		w []*Workers
		v []*Supervisors
	)

	_, err = o.QueryTable("workers").Filter("condos_id", condosID).Filter("deleted_at__isnull", true).All(&w)

	if err != nil {
		return
	}

	workersIDs := []interface{}{}

	for _, worker := range w {
		workersIDs = append(workersIDs, worker.ID)
	}

	_, err = o.QueryTable("supervisors").Filter("workers_id__in", workersIDs...).Filter("deleted_at__isnull", true).RelatedSel().All(&v)

	if err != nil {
		return
	}

	supervisors = v

	return

}
