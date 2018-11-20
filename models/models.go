package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //
)

type mysqlConnData struct {
	user   string
	pass   string
	ip     string
	dbName string
}

func init() {

	RunMode := beego.BConfig.RunMode

	if RunMode == "dev" {
		orm.Debug = true
	}

	//MYSQL
	var mysqlConnData mysqlConnData

	mysqlConnData.user = beego.AppConfig.String(RunMode + "::mysqluser")
	mysqlConnData.pass = beego.AppConfig.String(RunMode + "::mysqlpass")
	mysqlConnData.dbName = beego.AppConfig.String(RunMode + "::mysqldb")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysqlConnData.user+":"+mysqlConnData.pass+"@/"+mysqlConnData.dbName+"?charset=utf8")

	orm.RegisterModel(new(Assistances), new(Condos), new(Holidays), new(Points), new(Supervisors), new(Verifications), new(Watchers), new(Workers), new(Zones), new(Tasks), new(Goals), new(GoalsComments), new(Deliveries), new(Items), new(Notifications))

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
