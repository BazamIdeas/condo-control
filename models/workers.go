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

type dayAssistances struct {
	Entry            *Assistances `json:"entry,omitempty"`
	Break            *Assistances `json:"break,omitempty"`
	FinishBreak      *Assistances `json:"finish_break,omitempty"`
	Exit             *Assistances `json:"exit,omitempty"`
	Day              string       `json:"day,omitempty"`
	DayEnd           string       `json:"day_end,omitempty"`
	TotalWorkedHours float32      `json:"total_worked_hours"`
	ExtraWorkedHours float32      `json:"extra_worked_hours"`
	BaseValue        float32      `json:"base_value"`
	ExtraValue       float32      `json:"extra_value"`
	TotalValue       float32      `json:"total_value"`
	IsHoliday        bool         `json:"is_holiday"`
}

type monthDaysAssistances map[string]*dayAssistances

type monthDetail struct {
	Days             *monthDaysAssistances `json:"days"`
	TotalWorkedHours float32               `json:"total_worked_hours"`
	ExtraWorkedHours float32               `json:"extra_worked_hours"`
	ExtraValue       float32               `json:"extra_value"`
	BaseValue        float32               `json:"base_value"`
	TotalValue       float32               `json:"total_value"`
	Holidays         int                   `json:"holidays"`
}

type yearDetail map[int]*monthDetail

//Workers Model
type Workers struct {
	ID               int                     `orm:"column(id);pk" json:"id"`
	FirstName        string                  `orm:"column(first_name);size(255)" json:"first_name,omitempty" valid:"Required"`
	Condo            *Condos                 `orm:"rel(fk);column(condos_id)" json:"condos,omitempty"`
	Assistances      []*Assistances          `orm:"reverse(many)" json:"assistances,omitempty"`
	Notifications    []*Notifications        `orm:"reverse(many)" json:"notifications,omitempty"`
	ImageUUID        string                  `orm:"column(image_uuid)" json:"image_uuid,omitempty"`
	ImageMime        string                  `orm:"column(image_mime)" json:"-"`
	FaceID           string                  `orm:"column(face_id)" json:"-"`
	Approved         bool                    `orm:"column(approved)" json:"approved"`
	TodayAssistances map[string]*Assistances `orm:"-" json:"today_assistances,omitempty"`
	MonthData        *monthDetail            `orm:"-" json:"month_data,omitempty"`
	YearData         *yearDetail             `orm:"-" json:"year_data,omitempty"`
	CreatedAt        time.Time               `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt        time.Time               `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt        time.Time               `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName =
func (t *Workers) TableName() string {
	return "workers"
}

func (t *Workers) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"Assistances", "Notifications"}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddWorkers insert a new Workers into database and returns
// last inserted Id on success.
func AddWorkers(m *Workers) (id int64, err error) {
	o := orm.NewOrm()
	//m.Slug = GenerateSlug(m.TableName(), m.Name)
	id, err = o.Insert(m)

	if err != nil {
		return
	}

	m.ID = int(id)

	return
}

// GetWorkersByID retrieves Workers by Id. Returns error if
// Id doesn't exist
func GetWorkersByID(id int) (v *Workers, err error) {
	v = &Workers{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// GetAllWorkers retrieves all Workers matches certain condition. Returns empty list if
// no records exist
func GetAllWorkers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Workers))
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

	var l []Workers
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

// UpdateWorkersByID updates Workers by Id and returns error if
// the record to be updated doesn't exist
func UpdateWorkersByID(m *Workers) (err error) {
	o := orm.NewOrm()
	v := Workers{ID: m.ID}
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

// DeleteWorkers deletes Workers by Id and returns error if
// the record to be deleted doesn't exist
func DeleteWorkers(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Workers{ID: id}
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

//GetWorkersFromTrash return Workers soft Deleted
func GetWorkersFromTrash() (workers []*Workers, err error) {

	o := orm.NewOrm()

	var v []*Workers

	_, err = o.QueryTable("workers").Filter("deleted_at__isnull", false).All(&v)

	if err != nil {
		return
	}

	workers = v

	return

}

//GetWorkersByCondosID ...
func GetWorkersByCondosID(condosID int) (workers []*Workers, err error) {

	//condos, err := GetCondosByID(condosID)

	o := orm.NewOrm()

	v := []*Workers{}

	_, err = o.QueryTable("workers").Filter("deleted_at__isnull", true).Filter("condos_id", condosID).RelatedSel().All(&v)

	if err != nil {
		return
	}

	for _, worker := range v {
		//worker.GetTodayAssistances()
		worker.GetCurrentWorkTimeAssistances()
	}

	workers = v

	return

}

//GetTodayAssistances ...
func (t *Workers) GetTodayAssistances() (err error) {

	o := orm.NewOrm()
	todayDate := time.Now().Local().Format("2006-01-02")
	todayAssistances := []*Assistances{}

	qs := o.QueryTable("assistances").Filter("workers_id", t.ID)
	qs = qs.Filter("date__gte", todayDate+" 00:00:00").Filter("date__lte", todayDate+" 23:59:59")

	_, err = qs.All(&todayAssistances)

	if err != nil {

		return
	}

	mapAssistances := map[string]*Assistances{}

	if err != nil {
		t.TodayAssistances = mapAssistances
	}

	for _, todayAssistance := range todayAssistances {
		mapAssistances[todayAssistance.Type] = todayAssistance
	}

	t.TodayAssistances = mapAssistances

	return

}

/*
//GetMonthAssistancesData ...
func (t *Workers) GetMonthAssistancesData(year int, month time.Month) (err error) {

	qb, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		return
	}

	holidays, err := GetHolidaysByCondosID(year, month, t.Condo.ID)

	if err != nil {
		return
	}

	monthTarget := time.Date(year, month, 1, 1, 1, 1, 1, time.UTC)
	monthTargetString := jodaTime.Format("Y-M-d", monthTarget)

	//Construct query object
	qb.Select("assistances.id", "assistances.type", "assistances.date", "DAY(assistances.date) AS day").From("assistances").Where("assistances.workers_id = ?").And("YEAR(assistances.date) = YEAR(?)").And("MONTH(assistances.date) = MONTH(?)").OrderBy("assistances.date").Asc()

	sql := qb.String()

	o := orm.NewOrm()

	var paramMaps []orm.Params

	_, err = o.Raw(sql).SetArgs(t.ID, monthTargetString, monthTargetString).Values(&paramMaps)

	if err != nil {
		return
	}

	monthData := monthDaysAssistances{}

	for _, paramMap := range paramMaps {

		if dayData, ok := monthData[paramMap["day"].(string)]; ok {
			assistanceID, _ := strconv.Atoi(paramMap["id"].(string))
			assistance := &Assistances{ID: assistanceID, Date: paramMap["date"].(string), Type: paramMap["type"].(string)}
			dayData.assignTypes(assistance)
			continue
		}

		dayData := &dayAssistances{Day: paramMap["day"].(string)}
		assistanceID, _ := strconv.Atoi(paramMap["id"].(string))
		assistance := &Assistances{ID: assistanceID, Date: paramMap["date"].(string), Type: paramMap["type"].(string)}

		dayData.assignTypes(assistance)

		monthData[paramMap["day"].(string)] = dayData
	}

	monthDetail := &monthDetail{}

	for day, dayData := range monthData {

		if dayData.Entry == nil || dayData.Exit == nil {
			delete(monthData, day)
			continue
		}

		if (dayData.Break != nil && dayData.FinishBreak == nil) || (dayData.Break == nil && dayData.FinishBreak != nil) {
			delete(monthData, day)
			continue
		}

		if _, ok := holidays[day]; ok {
			dayData.IsHoliday = true
			monthDetail.Holidays++
		}

		entryDate, _ := jodaTime.Parse("Y-M-d HH:mm:ss", dayData.Entry.Date)
		exitDate, _ := jodaTime.Parse("Y-M-d HH:mm:ss", dayData.Exit.Date)

		workedDuration := exitDate.Sub(entryDate)

		workedHours := workedDuration.Hours()

		if dayData.Break != nil && dayData.FinishBreak != nil {
			breakDate, _ := jodaTime.Parse("Y-M-d HH:mm:ss", dayData.Break.Date)
			finishBreakDate, _ := jodaTime.Parse("Y-M-d HH:mm:ss", dayData.FinishBreak.Date)
			breakDuration := finishBreakDate.Sub(breakDate)
			breakHours := breakDuration.Hours()

			workedHours = workedHours - breakHours
		}

		dayData.TotalWorkedHours = float32(workedHours)

		if t.Condo == nil {
			continue
		}

		extraHours := dayData.TotalWorkedHours - float32(t.Condo.WorkingHours)

		var extraHoursValue float32

		if extraHours > 0 {

			dayData.ExtraWorkedHours = extraHours
			extraHourIncreased := t.Condo.HourValue + (t.Condo.HourValue * (t.Condo.ExtraHourIncrease / 100))

			extraHoursValue = extraHours * extraHourIncreased

			dayData.ExtraValue = extraHoursValue
		} else {
			dayData.ExtraValue = 0
		}

		monthDetail.ExtraWorkedHours += dayData.ExtraWorkedHours
		monthDetail.ExtraValue += dayData.ExtraValue

		if dayData.ExtraValue == 0 {
			dayData.BaseValue = dayData.TotalWorkedHours * t.Condo.HourValue
		} else {
			dayData.BaseValue = (dayData.TotalWorkedHours - dayData.ExtraWorkedHours) * t.Condo.HourValue
		}

		dayData.TotalValue = dayData.BaseValue + dayData.ExtraValue

		monthDetail.BaseValue += dayData.BaseValue
		monthDetail.TotalValue += dayData.TotalValue

	}

	monthDetail.Days = &monthData

	t.MonthData = monthDetail

	return
}
*/

//GetMonthAssistancesData ...
func (t *Workers) GetMonthAssistancesData(year int, month time.Month) (err error) {

	qb, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		return
	}

	holidays, err := GetHolidaysByCondosID(year, month, t.Condo.ID)

	if err != nil {
		return
	}

	monthTarget := time.Date(year, month, 1, 1, 1, 1, 1, time.UTC)
	monthTargetString := jodaTime.Format("Y-M-d", monthTarget)

	//Construct query object
	qb.Select("assistances.id", "assistances.type", "assistances.date", "DAY(assistances.date) AS day").From("assistances").Where("assistances.workers_id = ?").And("YEAR(assistances.date) = YEAR(?)").And("MONTH(assistances.date) = MONTH(?)").And("assistances.type = ?").OrderBy("assistances.date").Asc()

	sql := qb.String()

	o := orm.NewOrm()

	var paramMaps []orm.Params

	_, err = o.Raw(sql).SetArgs(t.ID, monthTargetString, monthTargetString, "entry").Values(&paramMaps)

	if err != nil {
		return
	}

	monthData := monthDaysAssistances{}

	for _, paramMap := range paramMaps {

		dayData := &dayAssistances{Day: paramMap["day"].(string)}
		assistanceID, _ := strconv.Atoi(paramMap["id"].(string))
		assistance := &Assistances{ID: assistanceID, Date: paramMap["date"].(string), Type: paramMap["type"].(string)}

		dayData.assignTypes(assistance)

		ox := orm.NewOrm()

		var nextAssistances []*Assistances

		_, err = ox.QueryTable("assistances").Filter("work_time_id", assistance.ID).All(&nextAssistances)

		if err != nil && err != orm.ErrNoRows {
			return
		}

		if err == orm.ErrNoRows {
			continue
		}

		for _, nextAssistance := range nextAssistances {
			dayData.assignTypes(nextAssistance)

			if nextAssistance.Type == "exit" {
				dayData.DayEnd = nextAssistance.Date
			}
		}

		monthData[paramMap["day"].(string)] = dayData

	}

	monthDetail := &monthDetail{}

	for day, dayData := range monthData {

		if dayData.Entry == nil || dayData.Exit == nil {
			delete(monthData, day)
			continue
		}

		if (dayData.Break != nil && dayData.FinishBreak == nil) || (dayData.Break == nil && dayData.FinishBreak != nil) {
			delete(monthData, day)
			continue
		}

		if _, ok := holidays[day]; ok {
			dayData.IsHoliday = true
			monthDetail.Holidays++
		}

		entryDate, _ := jodaTime.Parse("Y-M-d HH:mm:ss", dayData.Entry.Date)
		exitDate, _ := jodaTime.Parse("Y-M-d HH:mm:ss", dayData.Exit.Date)

		workedDuration := exitDate.Sub(entryDate)

		workedHours := workedDuration.Hours()

		if dayData.Break != nil && dayData.FinishBreak != nil {
			breakDate, _ := jodaTime.Parse("Y-M-d HH:mm:ss", dayData.Break.Date)
			finishBreakDate, _ := jodaTime.Parse("Y-M-d HH:mm:ss", dayData.FinishBreak.Date)
			breakDuration := finishBreakDate.Sub(breakDate)
			breakHours := breakDuration.Hours()

			workedHours = workedHours - breakHours
		}

		dayData.TotalWorkedHours = float32(workedHours)

		if t.Condo == nil {
			continue
		}

		extraHours := dayData.TotalWorkedHours - float32(t.Condo.WorkingHours)

		var extraHoursValue float32

		if extraHours > 0 {

			dayData.ExtraWorkedHours = extraHours
			extraHourIncreased := t.Condo.HourValue + (t.Condo.HourValue * (t.Condo.ExtraHourIncrease / 100))

			extraHoursValue = extraHours * extraHourIncreased

			dayData.ExtraValue = extraHoursValue
		} else {
			dayData.ExtraValue = 0
		}

		monthDetail.ExtraWorkedHours += dayData.ExtraWorkedHours
		monthDetail.ExtraValue += dayData.ExtraValue

		if dayData.ExtraValue == 0 {
			dayData.BaseValue = dayData.TotalWorkedHours * t.Condo.HourValue
		} else {
			dayData.BaseValue = (dayData.TotalWorkedHours - dayData.ExtraWorkedHours) * t.Condo.HourValue
		}

		dayData.TotalValue = dayData.BaseValue + dayData.ExtraValue

		monthDetail.BaseValue += dayData.BaseValue
		monthDetail.TotalValue += dayData.TotalValue

	}

	monthDetail.Days = &monthData

	t.MonthData = monthDetail
	return
}

func (d *dayAssistances) assignTypes(assistance *Assistances) {
	switch assistance.Type {

	case "entry":
		d.Entry = assistance
		break
	case "break":
		d.Break = assistance
		break
	case "finish-break":
		d.FinishBreak = assistance
		break
	case "exit":
		d.Exit = assistance
		break
	}
}

var fullYear = []time.Month{
	time.January,
	time.February,
	time.March,
	time.April,
	time.May,
	time.June,
	time.July,
	time.August,
	time.September,
	time.October,
	time.November,
	time.December,
}

//GetYearAssistancesData ...
func (t *Workers) GetYearAssistancesData(year int) (err error) {

	yearData := yearDetail{}

	for _, singleMonth := range fullYear {
		errMonth := t.GetMonthAssistancesData(year, singleMonth)
		if errMonth != nil && errMonth != orm.ErrNoRows {
			err = errMonth
			return
		}

		monthData := *t.MonthData
		yearData[int(singleMonth)] = &monthData
		t.MonthData = nil
	}

	t.YearData = &yearData

	return

}

//GetCurrentWorkTimeAssistances ...
func (t *Workers) GetCurrentWorkTimeAssistances() (err error) {

	o := orm.NewOrm()
	//todayDate := time.Now().Local().Format("2006-01-02")
	lastAssistanceEntry := &Assistances{}

	qs := o.QueryTable("assistances").Filter("workers_id", t.ID).Filter("type", "entry")
	qs = qs.Limit(1).OrderBy("-date")

	//qs = qs.Filter("date__gte", todayDate+" 00:00:00").Filter("date__lte", todayDate+" 23:59:59")

	err = qs.One(lastAssistanceEntry)

	mapAssistances := map[string]*Assistances{}

	if err != nil {
		t.TodayAssistances = mapAssistances
		return
	}

	mapAssistances[lastAssistanceEntry.Type] = lastAssistanceEntry

	nextAssistances := []*Assistances{}

	// TODO: VALIDAR QUE LA JORNADA ACTUAL NO PUEDA SUPERAR LAS 10 HORAS PASADAS

	qs = o.QueryTable("assistances").Filter("workers_id", t.ID).Filter("work_time_id", lastAssistanceEntry.ID)
	qs = qs.Limit(3)

	_, err = qs.All(&nextAssistances)

	if err != nil {
		t.TodayAssistances = mapAssistances
		return
	}

	for _, nextAssistance := range nextAssistances {
		mapAssistances[nextAssistance.Type] = nextAssistance
	}

	t.TodayAssistances = mapAssistances

	return

}
