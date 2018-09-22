package models

import (
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/gofrs/uuid"
)

//Images Model
type Images struct {
	ID        int         `orm:"column(id);pk" json:"id"`
	Priority  int8        `orm:"column(priority)" json:"priority,omitempty"`
	Name      string      `orm:"column(name);size(255)" json:"name,omitempty" valid:"Required"`
	Slug      string      `orm:"column(slug);size(255)" json:"slug,omitempty" valid:"AlphaDash"`
	UUID      string      `orm:"column(uuid);size(255)" json:"uuid,omitempty" valid:"Required"`
	Mimetype  string      `orm:"column(mimetype)" json:"mime_type,omitempty" valid:"Required"`
	URL       string      `orm:"-" json:"url,omitempty"`
	Portfolio *Portfolios `orm:"column(portfolios_id);rel(fk)" json:"portfolio,omitempty"`
	CreatedAt time.Time   `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt time.Time   `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt time.Time   `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName define Name
func (t *Images) TableName() string {
	return "images"
}

//AddImages insert a new Images into database and returns last inserted Id on success.
func AddImages(m *Images, fh *multipart.FileHeader, folderPath string) (id int64, err error) {
	o := orm.NewOrm()

	m.Slug = GenerateSlug(m.TableName(), m.Name)

	UUID, err := uuid.NewV4()

	if err != nil {
		return 0, err
	}

	m.UUID = UUID.String()

	err = o.Begin()

	if err != nil {
		return 0, err
	}

	_, err = o.Insert(m)

	if err != nil {

		errRoll := o.Rollback()

		if errRoll != nil {
			return 0, errRoll
		}

		return 0, err
	}

	f, err := fh.Open()

	if err != nil {

		errRoll := o.Rollback()

		if errRoll != nil {
			return 0, errRoll
		}

		return 0, err
	}

	defer f.Close()

	fileBytes, err := ioutil.ReadAll(f)

	if err != nil {

		errRoll := o.Rollback()

		if errRoll != nil {
			return 0, errRoll
		}

		return 0, err
	}

	newImagePath := folderPath + "/" + m.UUID

	err = ioutil.WriteFile(newImagePath, fileBytes, 644)

	if err != nil {

		errRoll := o.Rollback()

		if errRoll != nil {
			return 0, errRoll
		}

		return 0, err
	}

	err = o.Commit()

	if err != nil {
		return 0, err
	}

	//r := bufio.NewReader(*f)

	return
}

//GetImagesByID retrieves Images by Id. Returns error if Id doesn't exist
func GetImagesByID(id int) (v *Images, err error) {

	v = &Images{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	return
}

//GetAllImages retrieves all Images matches certain condition. Returns empty list if no records exist
func GetAllImages(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Images))
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

	var l []Images
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).Filter("deleted_at__isnull", true).RelatedSel().All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
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
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

//UpdateImagesByID updates Images by Id and returns error if the record to be updated doesn't exist
func UpdateImagesByID(m *Images) (err error) {
	o := orm.NewOrm()
	v := Images{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteImages deletes Images by Id and returns error if
// the record to be deleted doesn't exist
func DeleteImages(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Images{ID: id}
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

//GetImagesBySlug xxx
func GetImagesBySlug(slug string) (v *Images, err error) {

	o := orm.NewOrm()

	v = &Images{}

	err = o.QueryTable(v.TableName()).Filter("slug", slug).One(v)

	if err != nil {
		return nil, err
	}

	return
}
