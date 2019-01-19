package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/vjeantet/jodaTime"
)

//Goals Model
type Goals struct {
	ID            int              `orm:"column(id);pk" json:"id"`
	Name          string           `orm:"column(name);" json:"name,omitempty" valid:"Required"`
	Description   string           `orm:"column(description);" json:"description,omitempty"`
	Completed     bool             `orm:"column(completed);" json:"completed" valid:"Required"`
	Date          string           `orm:"column(date);type(datetime);" json:"date,omitempty"`
	DateEnd       string           `orm:"column(date_end);" json:"date_end,omitempty"`
	Task          *Tasks           `orm:"column(task_id);rel(fk);" json:"task,omitempty"`
	GoalsComments []*GoalsComments `orm:"reverse(many);" json:"comments,omitempty"`
	Token         string           `orm:"-" json:"token,omitempty"`
	CreatedAt     time.Time        `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt     time.Time        `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt     time.Time        `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName =
func (t *Goals) TableName() string {
	return "goals"
}

func (t *Goals) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"GoalsComments"}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddGoals insert a new Goals into database and returns
// last inserted Id on success.
func AddGoals(m *Goals) (id int64, err error) {

	o := orm.NewOrm()

	now := jodaTime.Format("Y-M-d HH:mm:ss", time.Now().In(orm.DefaultTimeLoc))
	m.Date = now

	id, err = o.Insert(m)

	if err != nil {
		return
	}

	m.ID = int(id)

	return
}

// GetGoalsByID retrieves Points by Id. Returns error if
// Id doesn't exist
func GetGoalsByID(id int) (v *Goals, err error) {
	v = &Goals{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// UpdateGoalsByID updates Points by Id and returns error if
// the record to be updated doesn't exist
func UpdateGoalsByID(m *Goals, ignoreStatus bool) (err error) {
	o := orm.NewOrm()
	v := Goals{ID: m.ID}
	// ascertain id exists in the database
	err = o.Read(&v)
	if err != nil {
		return
	}

	m.Date = v.Date
	m.DateEnd = v.DateEnd

	if ignoreStatus {
		m.Completed = v.Completed
	}

	_, err = o.Update(m)

	if err != nil {
		return
	}

	return
}

// DeleteGoals deletes Goals by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGoals(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Goals{ID: id}
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

//GetGoalsFromTrash return Points soft Deleted
func GetGoalsFromTrash() (goals []*Goals, err error) {

	o := orm.NewOrm()

	var v []*Goals

	_, err = o.QueryTable("goals").Filter("deleted_at__isnull", false).RelatedSel().All(&v)

	if err != nil {
		return
	}

	goals = v

	return

}
