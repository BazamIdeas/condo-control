package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/vjeantet/jodaTime"
)

//GoalsComments Model
type GoalsComments struct {
	ID          int       `orm:"column(id);pk" json:"id"`
	Description string    `orm:"column(description);null" json:"description,omitempty" valid:"Required"`
	Date        string    `orm:"column(date);type(datetime);" json:"date,omitempty"`
	Goal        *Goals    `orm:"rel(fk);column(goal_id)" json:"goal,omitempty"`
	Worker      *Workers  `orm:"rel(fk);column(worker_id)" json:"worker,omitempty"`
	Attachment  string    `orm:"attachment;null" json:"attachment,omitempty" `
	CreatedAt   time.Time `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt   time.Time `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt   time.Time `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName =
func (t *GoalsComments) TableName() string {
	return "goals_comments"
}

func (t *GoalsComments) loadRelations() {

	o := orm.NewOrm()

	relations := []string{}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddGoalsComments insert a new GoalsComments into database and returns
// last inserted Id on success.
func AddGoalsComments(m *GoalsComments) (id int64, err error) {

	o := orm.NewOrm()

	now := jodaTime.Format("Y-M-d HH:mm:ss", time.Now())
	m.Date = now

	id, err = o.Insert(m)

	if err != nil {
		return
	}

	m.ID = int(id)

	return
}

// GetGoalsCommentsByID retrieves Points by Id. Returns error if
// Id doesn't exist
func GetGoalsCommentsByID(id int) (v *GoalsComments, err error) {
	v = &GoalsComments{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// UpdateGoalsCommentsByID updates Points by Id and returns error if
// the record to be updated doesn't exist
func UpdateGoalsCommentsByID(m *GoalsComments) (err error) {
	o := orm.NewOrm()
	v := GoalsComments{ID: m.ID}
	// ascertain id exists in the database
	err = o.Read(&v)
	if err != nil {
		return
	}

	_, err = o.Update(m)

	if err != nil {
		return
	}

	return
}

// DeleteGoalsComments deletes GoalsComments by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGoalsComments(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := GoalsComments{ID: id}
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

//GetGoalsCommentsFromTrash return Points soft Deleted
func GetGoalsCommentsFromTrash() (goalsComments []*GoalsComments, err error) {

	o := orm.NewOrm()

	var v []*GoalsComments

	_, err = o.QueryTable("goals_comments").Filter("deleted_at__isnull", false).All(&v)

	if err != nil {
		return
	}

	goalsComments = v

	return

}
