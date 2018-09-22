package controllers

import (
	"GASE/controllers/services/mails"
	"GASE/models"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
)

// ClientsController operations for Clients
type ClientsController struct {
	BaseController
}

// URLMapping ...
func (c *ClientsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Login", c.Login)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetOneByEmail", c.GetOneByEmail)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("ChangePasswordRequest", c.ChangePasswordRequest)
	c.Mapping("ChangePassword", c.ChangePassword)
}

/* //Prepare Controller
func (c *ClientsController) Prepare() {

	URL := c.Ctx.Input.URL()

	fmt.Println(URL)

	if strings.HasPrefix(c.Ctx.Input.URL(), "/email") {
		valid := VerifyToken(c.Ctx.Input.Header("Authorization"))

		if !valid {
			c.DenyAccess()
			return
		}
	}

	return
} */

// Post ...
// @Title Post
// @Description create Clients
// @Param	body		body 	models.Clients	true		"body for Clients content"
// @Success 201 {int} models.Clients
// @Failure 400 body is empty
// @router / [post]
func (c *ClientsController) Post() {
	var v models.Clients

	// Validate empty body

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v.Password = "bazam12345"

	// Validate context body

	valid := validation.Validation{}

	_, err = valid.Valid(&v)

	if valid.HasErrors() {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	id, err := models.AddClients(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	ClientID := strconv.Itoa(int(id))

	v.Token, err = c.GenerateToken("Client", ClientID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	go (func() {
		HTMLParams := &mails.HTMLParams{Client: &v}

		Email := &mails.Email{
			To:         []string{v.Email},
			HTMLParams: HTMLParams,
		}
		mails.SendMail(Email, "001")
	})()

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v

	c.ServeJSON()

}

// GetOne ...
// @Title Get One
// @Description get Clients by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Clients
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ClientsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetClientsByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	v.Password = ""

	if err != nil {
		fmt.Println(err)
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Clients
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Clients
// @Failure 403
// @router / [get]
func (c *ClientsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllClients(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Clients
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Clients	true		"body for Clients content"
// @Success 200 {object} models.Clients
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ClientsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Clients{ID: id}

	// Validate context body
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = models.UpdateClientsByID(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = MessageResponse{
		Message:       "Updated element",
		PrettyMessage: "Elemento Actualizado",
	}

	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Clients
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ClientsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	trash := false

	if c.Ctx.Input.Query("trash") != "" {
		trash = true
	}

	err = models.DeleteClients(id, trash)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = MessageResponse{
		Message:       "Deleted element",
		PrettyMessage: "Elemento Eliminado",
	}

	c.ServeJSON()
}

// Login ...
// @Title Login
// @Description Login for Clients
// @Param	body		body 	models.Clients	true		"body for Clients content"
// @Success 201 {int} models.Clients
// @Failure 400 body is empty
// @router /login [post]
func (c *ClientsController) Login() {
	var v models.Clients

	// Validate empty body

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	// Validate context body

	valid := validation.Validation{}

	valid.Required(v.Email, "email")
	valid.Required(v.Password, "password")

	if valid.HasErrors() {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	id, err := models.LoginClients(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	ClientID := strconv.Itoa(id)

	v.Token, err = c.GenerateToken("Client", ClientID)

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = v

	c.ServeJSON()

}

// GetOneByEmail ...
// @Title GetOneByEmail
// @Description get Carts by Code
// @Param	Code		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Carts
// @Failure 403 :email is empty
// @router /email/:email [get]
func (c *ClientsController) GetOneByEmail() {
	email := c.Ctx.Input.Param(":email")

	if email == "" {
		err := errors.New("No se ha recibido el email")
		c.BadRequest(err)
		return
	}

	v, err := models.GetClientsByEmail(email)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = v
	c.ServeJSON()
}

// ChangePasswordRequest ...
// @Title ChangePasswordRequest
// @Description Change Password Request for Clients
// @router /change-password/:email [post]
func (c *ClientsController) ChangePasswordRequest() {

	email := c.Ctx.Input.Param(":email")

	if email == "" {
		err := errors.New("No se ha recibido el email")
		c.BadRequest(err)
		return
	}

	v, err := models.GetClientsByEmail(email)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	ClientID := strconv.Itoa(v.ID)

	token, err := c.GenerateToken("Client", ClientID, 1)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	v.Password = ""
	v.Token = token

	go (func() {
		HTMLParams := &mails.HTMLParams{Token: v.Token}

		Email := &mails.Email{
			To:         []string{v.Email},
			HTMLParams: HTMLParams,
		}
		mails.SendMail(Email, "002")
	})()

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v
	c.ServeJSON()

}

// ChangePassword ...
// @Title ChangePassword
// @Description Change Password Request for Clients
// @router /change-password/:token [put]
func (c *ClientsController) ChangePassword() {

	token := c.Ctx.Input.Param(":token")

	if token == "" {
		err := errors.New("No se ha recibido el token")
		c.BadRequest(err)
		return
	}

	newPassword := c.Ctx.Input.Query("password")

	if newPassword == "" {
		err := errors.New("No se ha recibido el password")
		c.BadRequest(err)
		return
	}

	decodedToken, err := VerifyToken(token, "Client")

	if err != nil {

		err := errors.New("El token Proporcionado es inválido")
		c.BadRequest(err)
		return
	}

	clientID, _ := strconv.Atoi(decodedToken.ID)

	v, err := models.GetClientsByID(clientID)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	v.Password = newPassword

	err = models.UpdateClientsByID(v)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	go (func() {
		HTMLParams := &mails.HTMLParams{Token: v.Token}

		Email := &mails.Email{
			To:         []string{v.Email},
			HTMLParams: HTMLParams,
		}
		mails.SendMail(Email, "003")
	})()

	c.Ctx.Output.SetStatus(200)
	c.ServeJSON()
}
