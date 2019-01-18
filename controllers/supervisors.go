package controllers

import (
	"condo-control/models"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
)

// SupervisorsController operations for Supervisors
type SupervisorsController struct {
	BaseController
}

//URLMapping ...
func (c *SupervisorsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("ChangePasswordSelf", c.ChangePasswordSelf)
}

// Post ...
// @Title Post
// @Description create Supervisors
// @router / [post]
func (c *SupervisorsController) Post() {

	v := models.Supervisors{}

	fmt.Println("hola")
	// Validate empty body

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	// Validate context body

	valid := validation.Validation{}

	b, _ := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	if v.Worker == nil {
		err = errors.New("Worker's info is empty")
		c.BadRequest(err)
		return
	}

	if v.Worker.FirstName == "" {

		err = errors.New("Worker's First Name is missing")
		c.BadRequest(err)
		return
	}

	if v.Worker.Condo == nil {
		err = errors.New("Condo's info is empty")
		c.BadRequest(err)
		return
	}

	v.Worker.Approved = true

	_, err = models.AddWorkers(v.Worker)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	_, err = models.AddSupervisors(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Supervisors by id
// @router /:id [get]
func (c *SupervisorsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetSupervisorsByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Supervisors
// @router / [get]
func (c *SupervisorsController) GetAll() {
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

	l, err := models.GetAllSupervisors(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Supervisors
// @router /:id [put]
func (c *SupervisorsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Supervisors{ID: id}

	// Validate empty body

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = models.UpdateSupervisorsByID(&v)

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

// ChangePasswordSelf ...
// @Title Put
// @Description update the ChangePassword
// @router /change-password/self [put]
func (c *SupervisorsController) ChangePasswordSelf() {

	token := c.Ctx.Input.Header("Authorization")

	decodedToken, err := VerifyToken(token, "Supervisor")

	if err != nil {
		c.BadRequest(err)
		return
	}

	supervisorID, _ := strconv.Atoi(decodedToken.UserID)

	supervisor, err := models.GetSupervisorsByID(supervisorID)

	if err != nil {
		c.BadRequestDontExists("Supervisor")
		return
	}

	v := models.Supervisors{}

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	if v.Password == "" {
		err = errors.New("Missing password")
		c.BadRequest(err)
		return
	}

	supervisor.Password = v.Password

	err = models.UpdateSupervisorsByID(supervisor)
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
// @Description delete the Supervisors
// @router /:id [delete]
func (c *SupervisorsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	trash := false

	if c.Ctx.Input.Query("trash") == "true" {
		trash = true
	}

	err = models.DeleteSupervisors(id, trash)

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

// GetAllFromTrash ...
// @Title Get All From Trash
// @Description Get All From Trash
// @router /trashed [patch]
func (c *SupervisorsController) GetAllFromTrash() {

	v, err := models.GetSupervisorsFromTrash()

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}

// RestoreFromTrash ...
// @Title Restore From Trash
// @Description Restore From Trash
// @router /:id/restore [put]
func (c *SupervisorsController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Supervisors{ID: id}

	err = models.RestoreFromTrash(v.TableName(), v.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}

// Login ...
// @Title Login
// @Description Login
// @router /login [post]
func (c *SupervisorsController) Login() {

	v := models.Supervisors{}

	// Validate empty body

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	// Validate context body

	valid := validation.Validation{}

	valid.Required(v.Username, "username")
	valid.Required(v.Password, "password")

	if valid.HasErrors() {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	id, err := models.LoginSupervisors(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	supervisorID := strconv.Itoa(id)
	condoID := strconv.Itoa(v.Worker.Condo.ID)

	v.Token, err = c.GenerateToken("Supervisor", supervisorID, condoID)

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = v

	c.ServeJSON()

}



// GenerateChangePasswordToken ..
// @Title Generate Change Password Token
// @Description Generate Change Password Token
// @Accept json
// @Success 200 {object} models.Watchers
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 email without Data
// @router /:email/change-password/ [post]
func (c *SupervisorsController) GenerateChangePasswordToken() {

	email := c.Ctx.Input.Param(":email")

	if email == "" {
		err := errors.New("missing email")
		c.BadRequest(err)
		return
	}

	worker, err := models.GetWorkersByEmail(email)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	supervisor, err := models.GetSupervisorsByWorkersID(worker.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	supervisorID := strconv.Itoa(supervisor.ID)
	condoID := strconv.Itoa(worker.Condo.ID)

	token, err := GenerateGeneralToken(supervisorID, condoID, nil, nil, nil)

	if err != nil {
		c.BadRequest(err)
		return
	}

	supervisor.Token = token
	c.Data["json"] = supervisor
	c.ServeJSON()

}

//ChangePassword ..
// @Title Change Password
// @Description Change Password
// @Accept json
// @Success 200 {object} models.Watchers
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @router /change-password/:token [put]
func (c *SupervisorsController) ChangePassword() {

	token := c.Ctx.Input.Param(":token")

	if token == "" {
		err := errors.New("missing token")
		c.BadRequest(err)
		return
	}

	decodedToken, err := VerifyGeneralToken(token)

	if err != nil {
		c.BadRequest(err)
		return
	}

	var v models.Supervisors

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	if v.Password == "" {
		err = errors.New("missing Password")
		c.BadRequest(err)
		return
	}

	supervisorID, _ := strconv.Atoi(decodedToken.UserID)

	supervisor, err := models.GetSupervisorsByID(supervisorID)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	supervisor.Password = v.Password

	err = models.UpdateSupervisorsByID(supervisor)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = supervisor
	c.ServeJSON()

}