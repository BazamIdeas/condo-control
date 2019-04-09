package controllers

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/astaxie/beego"
)

// AdminController operations for Admin
type AdminController struct {
	BaseController
}

type adminStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//URLMapping ...
func (c *AdminController) URLMapping() {
	c.Mapping("Login", c.Login)
}

//Login ...
// @Title Login
// @Description Login Admin
// @router /login [post]
func (c *AdminController) Login() {

	var v adminStruct

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		c.BadRequest(err)
		return
	}

	email, password, _ := getAdminData()

	if v.Email != email || v.Password != password {
		c.BadRequestDontExists("Admin")
		return
	}

	token, err := c.GenerateToken("Admin", "", "")

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Data["json"] = struct {
		Token string `json:"token"`
	}{Token: token}
	c.ServeJSON()
}

//Update ...
// @Title Update
// @Description Update Admin
// @router / [put]
func (c *AdminController) Update() {

	var v adminStruct

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		c.BadRequest(err)
		return
	}

	err = changeAdminData(v.Password)

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.ServeJSON()
}

func getAdminData() (email, password string, jsonFile bool) {

	email = beego.AppConfig.String("admin::email")
	password = beego.AppConfig.String("admin::password")

	if email == "" || password == "" {
		email = "admin"
		password = "admin"
	}

	file, err := os.Open("admin.json")

	if err != nil {
		return
	}

	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		return
	}

	var v adminStruct

	err = json.Unmarshal(fileBytes, &v)

	if err != nil {
		return
	}

	password = v.Password
	jsonFile = true

	return

}

func changeAdminData(password string) (err error) {

	os.Remove("admin.json")

	v := adminStruct{
		Password: password,
	}

	jsonData, err := json.Marshal(&v)

	if err != nil {
		return
	}

	err = ioutil.WriteFile("admin.json", jsonData, 0644)

	if err != nil {
		return
	}

	return

}
