package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//Residents Model
type Residents struct {
	ID          int       `orm:"column(id);pk" json:"id"`
	Name        string    `orm:"column(name);size(255)" json:"name,omitempty" valid:"Required"`
	Condo       *Condos   `orm:"rel(fk);column(condos_id)" json:"condos,omitempty"`
	Email       string    `orm:"column(email);size(255)" json:"email,omitempty" valid:"Required"`
	Phone       string    `orm:"column(phone);size(255)" json:"phone,omitempty"`
	Committee   bool      `orm:"column(committee);" json:"committee,omitempty"`
	RUT         string    `orm:"column(rut);" json:"rut,omitempty"`
	Password    string    `orm:"column(password);size(255)" json:"password,omitempty" valid:"Required"`
	ImageUUID   string    `orm:"column(image_uuid)" json:"image_uuid,omitempty"`
	ImageMime   string    `orm:"column(image_mime)" json:"image_mime,omitempty"`
	Departament string    `orm:"column(departament)" json:"departament,omitempty"`
	Percentage  float32   `orm:"column(percentage)" json:"percentage,omitempty"`
	Approved    bool      `orm:"column(approved)" json:"approved"`
	Votes       []*Votes  `orm:"reverse(many)" json:"votes,omitempty"`
	Token       string    `orm:"-" json:"token,omitempty"`
	CreatedAt   time.Time `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt   time.Time `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt   time.Time `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName =
func (t *Residents) TableName() string {
	return "residents"
}

func (t *Residents) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"Votes"}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// EmailExist verify if email exist into database and returns
// last inserted Id on success.
func EmailExist(m *Residents) (res string, err error) {
	o := orm.NewOrm()
	v := Residents{Email: m.Email}

	// verifiy if email exists in the database
	err = o.Read(&v, "email")

	if err == orm.ErrNoRows {
		res = "email not exist"
		return
	}
	if err != nil && err != orm.ErrNoRows {
		res = err.Error()
		return
	}
	res = "Email already exist"
	return
}

// AddResidents insert a new Residents into database and returns
// last inserted Id on success.
func AddResidents(m *Residents) (id int64, err error) {
	o := orm.NewOrm()

	m.Password = GetMD5Hash(m.Password)

	id, err = o.Insert(m)
	if err != nil {
		return
	}

	m.ID = int(id)
	return
}

// GetResidentsByID retrieves Residents by Id. Returns error if
// Id doesn't exist
func GetResidentsByID(id int) (v *Residents, err error) {
	v = &Residents{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// GetAllResidents retrieves all Residents matches certain condition. Returns empty list if
// no records exist
func GetAllResidents(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Residents))
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

	var l []Residents
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

// UpdateResidentsByID updates Residents by Id and returns error if
// the record to be updated doesn't exist
func UpdateResidentsByID(m *Residents) (err error) {
	o := orm.NewOrm()
	v := Residents{ID: m.ID}
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

// DeleteResidents deletes Residents by Id and returns error if
// the record to be deleted doesn't exist
func DeleteResidents(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Residents{ID: id}
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

//GetResidentsFromTrash return Residents soft Deleted
func GetResidentsFromTrash() (residents []*Residents, err error) {

	o := orm.NewOrm()

	var v []*Residents

	_, err = o.QueryTable("residents").Filter("deleted_at__isnull", false).All(&v)

	if err != nil {
		return
	}

	residents = v

	return

}

// LoginResidents login a Residents, returns
// if Exists.
func LoginResidents(m *Residents) (id int, err error) {
	o := orm.NewOrm()

	m.Password = GetMD5Hash(m.Password)

	err = o.QueryTable(m.TableName()).Filter("deleted_at__isnull", true).Filter("email", m.Email).Filter("password", m.Password).RelatedSel().One(m)

	if err != nil {
		return 0, err
	}

	m.Password = ""

	return m.ID, err
}

//GetResidentsByCondosID ...
func GetResidentsByCondosID(condosID int) (residents []*Residents, err error) {

	//condos, err := GetCondosByID(condosID)

	o := orm.NewOrm()

	v := []*Residents{}

	_, err = o.QueryTable("residents").Filter("deleted_at__isnull", true).Filter("condos_id", condosID).RelatedSel().All(&v)

	if err != nil {
		return
	}

	residents = v

	return

}

// GetResidentsByEmail retrieves Residents by email. Returns error if
// Id doesn't exist
func GetResidentsByEmail(email string) (v *Residents, err error) {
	v = &Residents{Email: email}

	o := orm.NewOrm()

	err = o.Read(v, "Email")

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}
