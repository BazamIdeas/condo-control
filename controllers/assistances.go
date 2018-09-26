package controllers

import (
	"condo-control/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/vjeantet/jodaTime"
)

// AssistancesController operations for Assistances
type AssistancesController struct {
	BaseController
}

//URLMapping ...
func (c *AssistancesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Assistances
// @router / [post]
func (c *AssistancesController) Post() {

	var v models.Assistances

	token := c.Ctx.Input.Header("Authorization")

	decodedToken, _ := VerifyToken(token, "Watcher")

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

	now := time.Now()

	date, err := jodaTime.Parse("Y-M-d HH:mm:ss", v.Date)

	if err != nil {
		c.BadRequest(err)
		return
	}

	if !date.Before(now) {
		err = errors.New("Verification date is a future Date")
		c.BadRequest(err)
		return
	}

	worker, err := models.GetWorkersByID(v.Worker.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	condoID, _ := strconv.Atoi(decodedToken.CondoID)

	if worker.Condo.ID != condoID {
		err = errors.New("Worker dont belong to Watcher's Condo")
		c.BadRequest(err)
		return
	}

	dur := now.Sub(date)

	if dur.Hours() > 6 {
		err = errors.New("Verification date is too old")
		c.BadRequest(err)
		return
	}

	// Validate foreings keys
	/* foreignsModels := map[string]int{
		"Watchers": v.Watcher.ID,
		"Workers":  v.Worker.ID,
	}

	resume := c.doForeignModelsValidation(foreignsModels)

	if !resume {
		return
	}
	*/

	tokenAssistance, err := GenerateGeneralToken(decodedToken.UserID, decodedToken.CondoID, nil, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v.Token = tokenAssistance

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Assistances by id
// @router /:id [get]
func (c *AssistancesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetAssistancesByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Assistances
// @router / [get]
func (c *AssistancesController) GetAll() {
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

	l, err := models.GetAllAssistances(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Assistances
// @router /:id [put]
func (c *AssistancesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Assistances{ID: id}

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

	//TODO:
	// Validate foreings keys

	/* exists := models.ValidateExists("Sectors", v.Sector.ID)

	if !exists {
		c.BadRequestDontExists("Sector")
		return
	} */

	err = models.UpdateAssistancesByID(&v)

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
// @Description delete the Assistances
// @router /:id [delete]
func (c *AssistancesController) Delete() {
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

	err = models.DeleteAssistances(id, trash)

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
func (c *AssistancesController) GetAllFromTrash() {

	v, err := models.GetAssistancesFromTrash()

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
func (c *AssistancesController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Assistances{ID: id}

	err = models.RestoreFromTrash(v.TableName(), v.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}
