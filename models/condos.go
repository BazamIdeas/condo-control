package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/vjeantet/jodaTime"
)

//Condos Model
type Condos struct {
	ID                int            `orm:"column(id);pk" json:"id"`
	Name              string         `orm:"column(name);size(255)" json:"name,omitempty" valid:"Required"`
	RUT               string         `orm:"column(rut);size(255)" json:"rut" valid:"Required"`
	UserLimit         int            `orm:"column(user_limit);size(255)" json:"user_limit" valid:"Required"`
	ZoneLimit         int            `orm:"column(zone_limit);size(255)" json:"zone_limit" valid:"Required"`
	HourValue         float32        `orm:"colum(hour_value);size(20)" json:"hour_value" valid:"Required"`
	ExtraHourIncrease float32        `orm:"colum(extra_hour_increase);size(20)" json:"extra_hour_increase" valid:"Required"`
	WorkingHours      int            `orm:"column(working_hours)" json:"working_hours" valid:"Required"`
	AssistancesMod    bool           `orm:"column(assistances_mod)" json:"assistances_mod" valid:"Required"`
	RoutesMod         bool           `orm:"column(routes_mod)" json:"routes_mod" valid:"Required"`
	DeliveryMod       bool           `orm:"column(delivery_mod)" json:"delivery_mod" valid:"Required"`
	TasksMod          bool           `orm:"column(tasks_mod)" json:"tasks_mod" valid:"Required"`
	ChecksMod         bool           `orm:"column(checks_mod)" json:"checks_mod" valid:"Required"`
	SurveysMod        bool           `orm:"column(surveys_mod)" json:"surveys_mod" valid:"Required"`
	AlertsTime        string         `orm:"column(alerts_time);type(alerts_time);null" json:"alerts_time,omitempty"`
	Zones             []*Zones       `orm:"reverse(many)" json:"zone,omitempty"`
	Workers           []*Workers     `orm:"reverse(many)" json:"workers,omitempty"`
	Holidays          []*Holidays    `orm:"reverse(many)" json:"holidays,omitempty"`
	Objects           []*Objects     `orm:"reverse(many)" json:"objects,omitempty"`
	Residents         []*Residents   `orm:"reverse(many)" json:"residents,omitempty"`
	Questions         []*Questions   `orm:"reverse(many)" json:"questions,omitempty"`
	Watchers          []*Watchers    `orm:"-" json:"watchers,omitempty"`
	Supervisors       []*Supervisors `orm:"-" json:"supervisors,omitempty"`
	CreatedAt         time.Time      `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt         time.Time      `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt         time.Time      `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

// EmptyAssistancesWorkers ...
type EmptyAssistancesWorkers struct {
	Entry       []*Workers `json:"entry,omitempty"`
	Break       []*Workers `json:"break,omitempty"`
	FinishBreak []*Workers `json:"finish_break,omitempty"`
	Exit        []*Workers `json:"exit,omitempty"`
}

//TableName =
func (t *Condos) TableName() string {
	return "condos"
}

func (t *Condos) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"Zones", "Workers", "Holidays", "Objects", "Residents", "Questions"}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddCondos insert a new Condos into database and returns
// last inserted Id on success.
func AddCondos(m *Condos) (id int64, err error) {
	o := orm.NewOrm()
	//m.Slug = GenerateSlug(m.TableName(), m.Name)
	id, err = o.Insert(m)

	if err != nil {
		return
	}

	m.ID = int(id)

	return
}

// GetCondosByID retrieves Condos by Id. Returns error if
// Id doesn't exist
func GetCondosByID(id int) (v *Condos, err error) {
	v = &Condos{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	for _, zone := range v.Zones {
		zone.loadRelations()
	}

	return
}

// GetAllCondos retrieves all Condos matches certain condition. Returns empty list if
// no records exist
func GetAllCondos(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Condos))
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

	var l []Condos
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

// UpdateCondosByID updates Condos by Id and returns error if
// the record to be updated doesn't exist
func UpdateCondosByID(m *Condos) (err error) {
	o := orm.NewOrm()
	v := Condos{ID: m.ID}
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

// DeleteCondos deletes Condos by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCondos(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Condos{ID: id}
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

//GetCondosFromTrash return Condos soft Deleted
func GetCondosFromTrash() (condos []*Condos, err error) {

	o := orm.NewOrm()

	var v []*Condos

	_, err = o.QueryTable("condos").Filter("deleted_at__isnull", false).RelatedSel().All(&v)

	if err != nil {
		return
	}

	condos = v

	return

}

//GetCondosByRUT ...
func GetCondosByRUT(RUT string) (condo *Condos, err error) {

	o := orm.NewOrm()

	v := Condos{RUT: RUT}

	err = o.Read(&v, "rut")

	//err = o.QueryTable("condos").Filter("RUT", RUT).Filter("deleted_at__isnull", false).RelatedSel().One(&v)

	if err != nil {
		return
	}

	condo = &v

	return

}

//GetCondosVerificationsByMonth ...
func GetCondosVerificationsByMonth(condosID int, year int, month time.Month) (Verifications []*Verifications, err error) {

	qb, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		return
	}

	monthTarget := time.Date(year, month, 1, 1, 1, 1, 1, time.UTC)
	monthTargetString := jodaTime.Format("Y-M-d", monthTarget)

	qb.Select("verifications.*").From("verifications, watchers, workers").Where("workers.condos_id = ?").And("watchers.workers_id = workers.id").And("verifications.watchers_id = watchers.id").And("YEAR(verifications.date) = YEAR(?)").And("MONTH(verifications.date) = MONTH(?)").OrderBy("verifications.date").Desc()

	sql := qb.String()

	o := orm.NewOrm()

	_, err = o.Raw(sql).SetArgs(condosID, monthTargetString, monthTargetString).QueryRows(&Verifications)

	if err != nil {
		return
	}

	for _, verification := range Verifications {

		watcher, errW := GetWatchersByID(verification.Watcher.ID)
		if errW == nil {
			verification.Watcher = watcher
		}

		point, errP := GetPointsByID(verification.Point.ID)
		if errP == nil {
			verification.Point = point
		}

	}

	return

}

// GetCondosWorkersEmptyAssistancesByDate ...
func GetCondosWorkersEmptyAssistancesByDate(condosID int, date time.Time) (emptyAssistancesWorkers *EmptyAssistancesWorkers, err error) {

	w := &EmptyAssistancesWorkers{}

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("workers.*").From("workers").Where("workers.condos_id = ?").And("workers.id NOT IN(SELECT assistances.workers_id FROM assistances WHERE YEAR(assistances.date) = YEAR(?) AND MONTH(assistances.date) = MONTH(?) AND DAY(assistances.date) = DAY(?) AND assistances.type = ?)").OrderBy("workers.id").Desc()

	sql := qb.String()

	assistanceTypes := []string{"entry", "break", "finish-break", "exit"}

	for _, assistanceType := range assistanceTypes {

		o := orm.NewOrm()

		workers := []*Workers{}

		dateStr := date.String()

		_, err = o.Raw(sql, condosID, dateStr, dateStr, dateStr, assistanceType).QueryRows(&workers)

		if err != nil && err != orm.ErrNoRows {
			return
		}

		switch assistanceType {
		case "entry":
			w.Entry = workers
		case "break":
			w.Break = workers
		case "finish-break":
			w.FinishBreak = workers
		case "exit":
			w.Exit = workers
		}

	}

	emptyAssistancesWorkers = w

	return
}

// GetCondosChecksByDate ...
func GetCondosChecksByDate(condoID int, date time.Time) (checks []*Checks, err error) {

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("checks.*").From("checks", "workers").Where("workers.condos_id = ?").And("workers.id = checks.workers_id").And("checks.date = ?").OrderBy("checks.id").Desc()

	sql := qb.String()

	v := []*Checks{}

	o := orm.NewOrm()

	_, err = o.Raw(sql, condoID, date.String()).QueryRows(&v)

	if err != nil {
		return
	}

	for _, check := range v {
		searchFK(check.TableName(), check.ID).One(check)

		check.loadRelations()

		if check.Occurrences == nil || len(check.Occurrences) == 0 {
			continue
		}

		for _, occurrence := range check.Occurrences {
			searchFK(occurrence.TableName(), occurrence.ID).One(occurrence)
		}
	}

	checks = v

	return
}

//GetCondosChecksByMonth ...
func GetCondosChecksByMonth(condoID int, year int, month time.Month) (checks []*Checks, err error) {

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("checks.*").From("checks", "workers").Where("workers.condos_id = ?").And("workers.id = checks.workers_id").And("YEAR(checks.date) = YEAR(?)").And("MONTH(checks.date) = MONTH(?)").OrderBy("checks.id").Desc()

	sql := qb.String()

	v := []*Checks{}

	o := orm.NewOrm()

	monthTarget := time.Date(year, month, 1, 1, 1, 1, 1, time.UTC)
	monthTargetString := jodaTime.Format("Y-M-d", monthTarget)

	_, err = o.Raw(sql, condoID, monthTargetString, monthTargetString).QueryRows(&v)

	if err != nil {
		return
	}

	for _, check := range v {
		searchFK(check.TableName(), check.ID).One(check)

		check.loadRelations()

		if check.Occurrences == nil || len(check.Occurrences) == 0 {
			continue
		}

		for _, occurrence := range check.Occurrences {
			searchFK(occurrence.TableName(), occurrence.ID).One(occurrence)
		}
	}

	checks = v

	return
}

func GetCondosWithoutLimit() (condos []*Condos, err error) {

	o := orm.NewOrm()

	var v []*Condos

	_, err = o.QueryTable("condos").Filter("deleted_at__isnull", true).RelatedSel().All(&v)

	if err != nil {
		return
	}

	condos = v

	return

}
