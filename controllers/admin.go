package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

// AdminController operations for Admin
type AdminController struct {
	BaseController
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

	var v struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		c.BadRequest(err)
		return
	}

	email := beego.AppConfig.String("admin::email")
	password := beego.AppConfig.String("admin::password")

	if email == "" || password == "" {
		email = "admin"
		password = "admin"
	}

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
