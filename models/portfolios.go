package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//Portfolios Model
type Portfolios struct {
	ID          int         `orm:"column(id);auto" json:"id"`
	Name        string      `orm:"column(name);size(255)" json:"name,omitempty" valid:"Required"`
	Slug        string      `orm:"column(slug);size(255)" json:"slug,omitempty" valid:"AlphaDash"`
	Description string      `orm:"column(description)" json:"description,omitempty" valid:"Required"`
	Client      string      `orm:"column(client);size(255)" json:"client,omitempty" valid:"Required"`
	Priority    int8        `orm:"column(priority)" json:"priority,omitempty"`
	Location    *Locations  `orm:"column(locations_id);rel(fk)" json:"location,omitempty"`
	Service     *Services   `orm:"column(services_id);rel(fk)" json:"service,omitempty"`
	Activity    *Activities `orm:"column(activities_id);rel(fk)" json:"activity,omitempty"`
	Images      []*Images   `orm:"reverse(many)" json:"images,omitempty"`
	CreatedAt   time.Time   `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt   time.Time   `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt   time.Time   `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName define Name
func (t *Portfolios) TableName() string {
	return "portfolios"
}

func (t *Portfolios) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"Images"}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddPortfolios insert a new Portfolios into database and returns last inserted Id on success.
func AddPortfolios(m *Portfolios) (id int64, err error) {
	o := orm.NewOrm()

	m.Slug = GenerateSlug(m.TableName(), m.Name)

	id, err = o.Insert(m)
	return
}

// GetPortfoliosByID retrieves Portfolios by Id. Returns error if Id doesn't exist
func GetPortfoliosByID(id int) (v *Portfolios, err error) {

	v = &Portfolios{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return v, err
}

// GetAllPortfolios retrieves all Portfolios matches certain condition. Returns empty list if no records exist
func GetAllPortfolios(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Portfolios))
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

	var l []Portfolios
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

//UpdatePortfoliosByID updates Portfolios by Id and returns error if the record to be updated doesn't exist
func UpdatePortfoliosByID(m *Portfolios) (err error) {
	o := orm.NewOrm()
	v := Portfolios{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePortfolios deletes Portfolios by Id and returns error if
// the record to be deleted doesn't exist
func DeletePortfolios(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Portfolios{ID: id}

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

// GetPortfoliosByCustomSearch retrieves Portfolios by Id. Returns error if Id doesn't exist
func GetPortfoliosByCustomSearch(filters map[string]int, limit int, offset int) (portfolios []*Portfolios, err error) {

	qb, err := orm.NewQueryBuilder("mysql")

	if err != nil {
		return nil, err
	}

	qb = qb.Select("portfolios.*")

	var (
		//internalFilters map[string]int

		externalKeys []string
		internalKeys []string
	)

	externalFilters := make(map[string]int)

	if countryID, ok := filters["countries"]; ok {

		delete(filters, "countries")

		if _, ok = filters["locations"]; !ok {
			externalFilters["countries"] = countryID
			externalKeys = append(externalKeys, "countries")
		}

	}

	if sectorID, ok := filters["sectors"]; ok {

		delete(filters, "sectors")

		if _, ok = filters["activities"]; !ok {
			externalFilters["sectors"] = sectorID
			externalKeys = append(externalKeys, "sectors")
		}

	}

	for filterKey := range filters {
		internalKeys = append(internalKeys, filterKey)
	}
	/*
		extraTables := ""

		if len(internalKeys) > 0 {
			extraTables = ", " + strings.Join(internalKeys, ",")
		} */

	qb = qb.From("portfolios" /*  + extraTables */)

	////

	for _, externalKey := range externalKeys {

		var interTable string

		if externalKey == "countries" {

			interTable = "locations"

		} else if externalKey == "sectors" {

			interTable = "activities"

		}

		qb = qb.InnerJoin(interTable).On("portfolios." + interTable + "_id = " + interTable + ".id")

		qb = qb.InnerJoin(externalKey).On(interTable + "." + externalKey + "_id = " + externalKey + ".id")
	}

	for _, internalKey := range internalKeys {

		qb = qb.InnerJoin(internalKey).On("portfolios." + internalKey + "_id = " + internalKey + ".id")

	}

	////

	var whereString string

	for _, internalKey := range internalKeys {

		id := strconv.Itoa(filters[internalKey])

		whereString += internalKey + ".id = " + id

		if internalKey != internalKeys[len(internalKeys)-1] {
			whereString += " AND "
		}
	}

	for _, externalKey := range externalKeys {

		if whereString != "" {
			whereString += " AND "
		}

		id := strconv.Itoa(externalFilters[externalKey])

		whereString += externalKey + ".id = " + id

	}

	if whereString != "" {
		qb.Where(whereString)
	}

	qb = qb.OrderBy("portfolios.priority").Asc()

	if limit != 0 {
		qb = qb.Limit(limit)
	}

	if offset != 0 {
		qb = qb.Offset(offset)
	}

	sql := qb.String()

	fmt.Println(sql)

	o := orm.NewOrm()

	_, err = o.Raw(sql).QueryRows(&portfolios)

	if err != nil {
		return nil, err
	}

	if portfolios == nil {
		return nil, orm.ErrNoRows
	}

	for key := range portfolios {

		searchFK(portfolios[key].TableName(), portfolios[key].ID).One(portfolios[key])
		portfolios[key].loadRelations()
	}

	return
}
