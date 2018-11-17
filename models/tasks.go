package models

import (
	"time"

	"github.com/vjeantet/jodaTime"

	"github.com/astaxie/beego/orm"
)

//Tasks Model
type Tasks struct {
	ID        int       `orm:"column(id);pk" json:"id"`
	Name      string    `orm:"column(name);" json:"name,omitempty" valid:"Required"`
	Status    string    `orm:"column(status);" json:"status" valid:"Required"`
	Date      string    `orm:"column(date);auto_now_add;type(datetime);" json:"date,omitempty"`
	DateEnd   string    `orm:"column(date_end);" json:"date_end,omitempty"`
	Worker    *Workers  `orm:"column(worker_id);rel(fk);" json:"worker,omitempty"`
	Goals     []*Goals  `orm:"reverse(many);" json:"goals,omitempty"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt time.Time `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName =
func (t *Tasks) TableName() string {
	return "tasks"
}

func (t *Tasks) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"Goals"}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddTasks insert a new Tasks into database and returns
// last inserted Id on success.
func AddTasks(m *Tasks) (id int64, err error) {
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

// GetTasksByID retrieves Points by Id. Returns error if
// Id doesn't exist
func GetTasksByID(id int) (v *Tasks, err error) {
	v = &Tasks{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// UpdateTasksByID updates Points by Id and returns error if
// the record to be updated doesn't exist
func UpdateTasksByID(m *Tasks) (err error) {
	o := orm.NewOrm()
	v := Tasks{ID: m.ID}
	// ascertain id exists in the database
	err = o.Read(&v)
	if err != nil {
		return
	}

	m.Date = v.Date
	m.DateEnd = v.DateEnd
	m.Status = v.Status

	_, err = o.Update(m)

	if err != nil {
		return
	}

	return
}

// DeleteTasks deletes Tasks by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTasks(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Tasks{ID: id}
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

//GetTasksFromTrash return Points soft Deleted
func GetTasksFromTrash() (tasks []*Tasks, err error) {

	o := orm.NewOrm()

	var v []*Tasks

	_, err = o.QueryTable("tasks").Filter("deleted_at__isnull", false).All(&v)

	if err != nil {
		return
	}

	tasks = v

	return

}

//GetTasksByWorkersID ...
func GetTasksByWorkersID(workerID int) (tasks []*Tasks, err error) {

	o := orm.NewOrm()

	var v []*Tasks

	_, err = o.QueryTable("tasks").Filter("worker_id", workerID).Filter("deleted_at__isnull", true).RelatedSel().All(&v)

	if err != nil {
		return
	}

	for _, task := range v {
		task.loadRelations()
	}

	tasks = v

	return

}
