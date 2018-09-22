package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	mgo "github.com/globalsign/mgo"
	_ "github.com/go-sql-driver/mysql" //
	"github.com/gosimple/slug"
)

type mysqlConnData struct {
	user   string
	pass   string
	ip     string
	dbName string
}

//Mongo
var session *mgo.Session

// Conn return mongodb session.
func Conn() *mgo.Session {
	return session.Copy()
}

func init() {

	RunMode := beego.BConfig.RunMode

	if RunMode == "dev" {
		orm.Debug = true
	}

	//MONGO
	url := beego.AppConfig.String("mongodb::url")

	sess, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	session = sess
	session.SetMode(mgo.Monotonic, true)

	//MYSQL
	var mysqlConnData mysqlConnData

	mysqlConnData.user = beego.AppConfig.String(RunMode + "::mysqluser")
	mysqlConnData.pass = beego.AppConfig.String(RunMode + "::mysqlpass")
	mysqlConnData.dbName = beego.AppConfig.String(RunMode + "::mysqldb")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysqlConnData.user+":"+mysqlConnData.pass+"@/"+mysqlConnData.dbName+"?charset=utf8")

	orm.RegisterModel(new(Activities), new(Carts), new(Clients), new(Countries), new(Coupons), new(Currencies), new(Gateways), new(Images), new(Locations), new(Orders), new(Portfolios), new(Prices), new(Sectors), new(Services))

	// Add defaults to database
	count, _ := AddDefaultDataCurrencies()
	if count > 0 {
		fmt.Println("Added Currencies : ", count)
	}

	count, _ = addDefaultDataCountries()
	if count > 0 {
		fmt.Println("Added Countries : ", count)
	}

	count, _ = AddDefaultDataSectors()
	if count > 0 {
		fmt.Println("Added Sectors : ", count)
	}

	count, _ = addDefaultDataActivities()
	if count > 0 {
		fmt.Println("Added Activities : ", count)
	}

	count, _ = AddDefaultDataGateways()
	if count > 0 {
		fmt.Println("Added Gateways : ", count)
	}

	count, _ = addRelationsGatewaysCurrencies()
	if count > 0 {
		fmt.Println("Added relations GatewaysCurrencies : ", count)
	}

	count, _ = AddDefaultDataServices()
	if count > 0 {
		fmt.Println("Added Services : ", count)
	}

	AddDefaultDataPrices()

	AddDefaultDataUsers()

}

//LoadRelations of the model
func searchFK(tableName string, id int) (query orm.QuerySeter) {

	o := orm.NewOrm()

	query = o.QueryTable(tableName).Filter("id", id).Filter("deleted_at__isnull", true).RelatedSel()

	return
}

//ValidateExists FK
func ValidateExists(modelName string, id int) (exists bool) {

	o := orm.NewOrm()
	modelName = strings.ToLower(modelName)
	exists = o.QueryTable(modelName).Filter("id", id).Exist()

	return
}

//GenerateSlug return a slug
func GenerateSlug(modelName string, name string) (generatedSlug string) {

	o := orm.NewOrm()
	generatedSlug = slug.Make(name)
	var slugInt int

	modelName = strings.ToLower(modelName)

	queryO := o.QueryTable(modelName).Filter("slug__startswith", generatedSlug)

	switch modelName {
	case "countries":

		var countries []*Countries
		queryO.All(&countries)

		for _, val := range countries {
			formatSlug(val.Slug, generatedSlug, &slugInt)
		}

	case "activities":

		var activities []*Activities
		queryO.All(&activities)

		for _, val := range activities {
			formatSlug(val.Slug, generatedSlug, &slugInt)
		}

	case "locations":

		var locations []*Locations
		queryO.All(&locations)

		for _, val := range locations {
			formatSlug(val.Slug, generatedSlug, &slugInt)
		}

	case "images":

		var images []*Images
		queryO.All(&images)

		for _, val := range images {
			formatSlug(val.Slug, generatedSlug, &slugInt)
		}

	case "portfolios":

		var portfolio []*Portfolios
		queryO.All(&portfolio)

		for _, val := range portfolio {
			formatSlug(val.Slug, generatedSlug, &slugInt)
		}

	case "sectors":

		var sector []*Sectors
		queryO.All(&sector)

		for _, val := range sector {
			formatSlug(val.Slug, generatedSlug, &slugInt)
		}

	case "services":

		var service []*Services
		queryO.All(&service)

		for _, val := range service {
			formatSlug(val.Slug, generatedSlug, &slugInt)
		}

	}

	if slugInt > 0 {
		generatedSlug = generatedSlug + "-" + strconv.Itoa(slugInt+1)
	}

	return
}

func formatSlug(actualSlug string, generatedSlug string, slugInt *int) (err error) {

	replacedSlug := strings.Replace(actualSlug, generatedSlug, "", 1)

	replacedSlug = strings.Replace(replacedSlug, "-", "", 1)

	if replacedSlug == "" {
		replacedSlug = "1"
	}

	slugNumber, err := strconv.Atoi(replacedSlug)

	if err != nil {
		return
	}

	if slugNumber > *slugInt {
		*slugInt = slugNumber
	}

	return
}

// RelationsM2M insert | delete M2M Relations on database
func RelationsM2M(operation string, entityFieldName string, entityID int, relationFieldName string, relationsIDs []int) (count int, err error) {

	tableName := entityFieldName + "_" + relationFieldName + "s"
	entityFieldName = entityFieldName + "_id"
	relationFieldName = relationFieldName + "_id"

	QueryRaw := "INSERT INTO " + tableName + " (" + entityFieldName + "," + relationFieldName + ") VALUES (?,?)"

	if operation == "DELETE" {

		QueryRaw = "DELETE FROM " + tableName + " WHERE " + entityFieldName + " = ? AND " + relationFieldName + " = ?"

	}

	o := orm.NewOrm()
	err = o.Begin()

	if err != nil {
		return 0, err
	}

	for _, relationID := range relationsIDs {

		Args := []int{entityID, relationID}

		_, err := o.Raw(QueryRaw, Args).Exec()

		if err != nil {
			errRoll := o.Rollback()
			if errRoll != nil {
				return 0, errRoll
			}
			return 0, err
		}

		count++
	}

	err = o.Commit()

	if err != nil {
		return 0, err
	}

	return
}

// AddRelationsM2MQuantity insert M2M Relations on database
func AddRelationsM2MQuantity(entityFieldName string, entityID int, relationFieldName string, relations []map[string]int) (count int, err error) {

	tableName := entityFieldName + "_" + relationFieldName + "s"
	entityFieldName = entityFieldName + "_id"
	relationFieldName = relationFieldName + "_id"

	o := orm.NewOrm()
	err = o.Begin()

	if err != nil {
		return 0, err
	}

	QueryRaw := "INSERT INTO " + tableName + " (" + entityFieldName + "," + relationFieldName + ", quantity) VALUES (?,?,?)"

	for _, relation := range relations {

		Args := []int{entityID, relation["id"], relation["quantity"]}

		_, err := o.Raw(QueryRaw, Args).Exec()

		if err != nil {
			errRoll := o.Rollback()
			if errRoll != nil {
				return 0, errRoll
			}
			return 0, err
		}

		count++
	}

	err = o.Commit()

	if err != nil {
		return 0, err
	}

	return
}

// GetMD5Hash =
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

//RestoreFromTrash ...
func RestoreFromTrash(tableName string, id int) (err error) {

	o := orm.NewOrm()

	num, err := o.QueryTable(tableName).Filter("id", id).Filter("deleted_at__isnull", false).Update(orm.Params{
		"deleted_at": nil,
	})

	if err != nil {
		return
	}

	if num == 0 {
		err = errors.New("not found")
	}

	return
}

/* func filterDeleted(model interface{}) {

	var timeNil time.Time

	switch m := model.(type) {
	case *Portfolios:

		if m.Activity != nil && m.Activity.DeletedAt != timeNil {
			m.Activity = nil
		} else if m.Activity != nil {
			//filterDeleted(m.Activity)
		}

		if m.Location != nil && m.Location.DeletedAt != timeNil {
			m.Location = nil
		} else if m.Location != nil {
			//filterDeleted(m.Location)
		}

		if m.Service != nil && m.Service.DeletedAt != timeNil {
			m.Service = nil
		} else if m.Service != nil {
			//filterDeleted(m.Service)
		}

		if m.Images != nil {
			for i, image := range m.Images {
				if image != nil && image.DeletedAt != timeNil {
					m.Images[i] = nil
				} else if m.Service != nil {
					//filterDeleted(image)
				}

			}
		}

	}
} */
