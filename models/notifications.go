package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/vjeantet/jodaTime"
)

// Notifications ...
type Notifications struct {
	ID          int       `orm:"column(id);pk" json:"id"`
	Name        string    `orm:"column(name);" json:"name,omitempty" valid:"Required"`
	Description string    `orm:"column(description);" json:"description,omitempty" valid:"Required"`
	Approved    bool      `orm:"column(approved);" json:"approved"`
	View        bool      `orm:"column(view);" json:"view"`
	Date        string    `orm:"column(date);type(string);" json:"date,omitempty"`
	ImageUUID   string    `orm:"column(image_uuid);null" json:"image_uuid,omitempty"`
	ImageMime   string    `orm:"column(image_mime);null" json:"image_mime,omitempty"`
	Worker      *Workers  `orm:"rel(fk);column(workers_id)" json:"worker,omitempty"`
	CreatedAt   time.Time `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt   time.Time `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt   time.Time `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName =
func (t *Notifications) TableName() string {
	return "notifications"
}

func (t *Notifications) loadRelations() {

	o := orm.NewOrm()

	relations := []string{}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddNotifications insert a new Notifications into database and returns
// last inserted Id on success.
func AddNotifications(m *Notifications) (id int64, err error) {
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

// GetNotificationsByID retrieves Notifications by Id. Returns error if
// Id doesn't exist
func GetNotificationsByID(id int) (v *Notifications, err error) {
	v = &Notifications{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// UpdateNotificationsByID updates Notifications by Id and returns error if
// the record to be updated doesn't exist
func UpdateNotificationsByID(m *Notifications) (err error) {
	o := orm.NewOrm()
	v := Notifications{ID: m.ID}
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

// DeleteNotifications deletes Notifications by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNotifications(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Notifications{ID: id}
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

//GetNotificationsFromTrash return Notifications soft Deleted
func GetNotificationsFromTrash() (notifications []*Notifications, err error) {

	o := orm.NewOrm()

	var v []*Notifications

	_, err = o.QueryTable("notifications").Filter("deleted_at__isnull", false).All(&v)

	if err != nil {
		return
	}

	notifications = v

	return

}

//GetNotificationsByCondosID return Supervisors soft Deleted
func GetNotificationsByCondosID(condosID int) (notifications []*Notifications, err error) {

	o := orm.NewOrm()

	var (
		w []*Workers
		v []*Notifications
	)

	_, err = o.QueryTable("notifications").Filter("condos_id", condosID).All(&w)

	if err != nil {
		return
	}

	workersIDs := []interface{}{}

	for _, worker := range w {
		workersIDs = append(workersIDs, worker.ID)
	}

	_, err = o.QueryTable("notifications").Filter("worker_id__in", workersIDs...).Filter("deleted_at__isnull", true).RelatedSel().All(&v)

	if err != nil {
		return
	}

	notifications = v

	return

}
