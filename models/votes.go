package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//Votes Model
type Votes struct {
	ID        int        `orm:"column(id);pk" json:"id"`
	Resident  *Residents `orm:"rel(fk);column(residents_id)" json:"residents,omitempty"`
	Question  *Questions `orm:"rel(fk);column(questions_id)" json:"questions,omitempty"`
	Accepted  bool       `orm:"column(accepted);" json:"accepted,omitempty" valid:"Required"`
	Comment   string     `orm:"column(comment);size(255)" json:"comment,omitempty"`
	CreatedAt time.Time  `orm:"column(created_at);type(datetime);null;auto_now_add" json:"created_at,omitempty"`
	UpdatedAt time.Time  `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt time.Time  `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName =
func (t *Votes) TableName() string {
	return "votes"
}

func (t *Votes) loadRelations() {

	o := orm.NewOrm()

	relations := []string{}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddVotes insert a new Votes into database and returns
// last inserted Id on success.
func AddVotes(m *Votes) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(m)
	if err != nil {
		return
	}

	m.ID = int(id)
	return
}

// GetVotesByID retrieves Votes by Id. Returns error if
// Id doesn't exist
func GetVotesByID(id int) (v *Votes, err error) {
	v = &Votes{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// GetAllVotes retrieves all Votes matches certain condition. Returns empty list if
// no records exist
func GetAllVotes(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Votes))
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

	var l []Votes
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

// UpdateVotesByID updates Votes by Id and returns error if
// the record to be updated doesn't exist
func UpdateVotesByID(m *Votes) (err error) {
	o := orm.NewOrm()
	v := Votes{ID: m.ID}
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

// DeleteVotes deletes Votes by Id and returns error if
// the record to be deleted doesn't exist
func DeleteVotes(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Votes{ID: id}
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

//GetVotesFromTrash return Votes soft Deleted
func GetVotesFromTrash() (votes []*Votes, err error) {

	o := orm.NewOrm()

	var v []*Votes

	_, err = o.QueryTable("votes").Filter("deleted_at__isnull", false).All(&v)

	if err != nil {
		return
	}

	votes = v

	return

}
