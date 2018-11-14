package controllers

import (
	"condo-control/models"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/vjeantet/jodaTime"

	"github.com/astaxie/beego/validation"
)

// WatchersController operations for Watchers
type WatchersController struct {
	BaseController
}

//URLMapping ...
func (c *WatchersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Login", c.Login)
	c.Mapping("GetSelf", c.GetSelf)
	c.Mapping("GetVerificationsByDate", c.GetVerificationsByDate)

}

// Post ...
// @Title Post
// @Description create Watchers
// @Accept json
// @Param   Authorization     header   string true       "Supervisor's Token"
// @Param   username     body   string true       "watcher's Username"
// @Param   password     body   string true       "watcher's password"
// @Param   phone     body   string false       "watcher's phone"
// @Param   worker     body   object false       "Worker object (first name)"
// @Success 200 {object} models.Watchers
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Zones Don't exists
// @Failure 409 Condo's User limit reached
// @router / [post]
func (c *WatchersController) Post() {

	token := c.Ctx.Input.Header("Authorization")

	decodedToken, err := VerifyToken(token, "Supervisor")

	if err != nil {
		c.BadRequest(err)
		return
	}

	condoID, err := strconv.Atoi(decodedToken.CondoID)
	if err != nil {
		c.BadRequest(err)
		return
	}

	condos, err := models.GetCondosByID(condoID)
	if err != nil {
		c.BadRequestDontExists("Condos")
		return
	}

	condosWorkersCount := len(condos.Workers) + 1

	if condosWorkersCount > condos.UserLimit {
		c.Ctx.Output.SetStatus(409)
		c.Data["json"] = MessageResponse{
			Code:    409,
			Message: "Condo's user limit reached",
		}
		c.ServeJSON()
		return
	}

	var v models.Watchers

	// Validate empty body
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

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
		err = errors.New("Worker info is empty")
		c.BadRequest(err)
		return
	}

	if v.Worker.FirstName == "" {

		err = errors.New("Worker's First Name is missing")
		c.BadRequest(err)
		return
	}

	v.Worker.Condo = &models.Condos{ID: condoID}
	v.Worker.Approved = true

	_, err = models.AddWorkers(v.Worker)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	_, err = models.AddWatchers(&v)
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
// @Description get Watchers by id
// @router /:id [get]
func (c *WatchersController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetWatchersByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Watchers
// @router / [get]
func (c *WatchersController) GetAll() {
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

	l, err := models.GetAllWatchers(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Watchers
// @router /:id [put]
func (c *WatchersController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Watchers{ID: id}

	// Validate empty body
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = models.UpdateWatchersByID(&v)

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
// @Description delete the Watchers
// @router /:id [delete]
func (c *WatchersController) Delete() {
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

	err = models.DeleteWatchers(id, trash)

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
func (c *WatchersController) GetAllFromTrash() {

	v, err := models.GetWatchersFromTrash()

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
func (c *WatchersController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Watchers{ID: id}

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
func (c *WatchersController) Login() {

	v := models.Watchers{}

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

	id, err := models.LoginWatchers(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	watcherID := strconv.Itoa(id)
	condoID := strconv.Itoa(v.Worker.Condo.ID)

	v.Token, err = c.GenerateToken("Watcher", watcherID, condoID)

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = v

	c.ServeJSON()

}

// GetSelf ...
// @Title Get Self
// @Description Get Self
// @router /self [get]
func (c *WatchersController) GetSelf() {

	token := c.Ctx.Input.Header("Authorization")

	decodedToken, _ := VerifyToken(token, "Supervisor")

	//Disclamer, token already verified
	id, err := strconv.Atoi(decodedToken.CondoID)
	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetWatchersByCondosID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetVerificationsByDate ...
// @Title Get Verifications By Date
// @Description Get Verifications By Date
// @router /:id/verifications/:date [get]
func (c *WatchersController) GetVerificationsByDate() {

	fmt.Println("hola")

	authToken := c.Ctx.Input.Header("Authorization")

	decAuthToken, _ := VerifyToken(authToken, "Supervisor")

	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	watcher, err := models.GetWatchersByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	condoID, _ := strconv.Atoi(decAuthToken.CondoID)

	if watcher.Worker.Condo.ID != condoID {
		err = errors.New("Watcher's Condo and Supervisor's Condo Don't match")
		c.BadRequest(err)
		return
	}

	dateString := c.Ctx.Input.Param(":date")
	date, err := jodaTime.Parse("Y-M-d", dateString)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = watcher.GetVerificationsByDate(date)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = watcher
	c.ServeJSON()

}
