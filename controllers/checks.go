package controllers

import (
	"condo-control/models"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/vjeantet/jodaTime"
)

// ChecksController operations for Holidays
type ChecksController struct {
	BaseController
}

//URLMapping ...
func (c *ChecksController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetAllFromTrash", c.GetAllFromTrash)
	c.Mapping("RestoreFromTrash", c.RestoreFromTrash)
}

// Post ...
// @Title Post
// @Description create Checks
// @router / [post]
func (c *ChecksController) Post() {

	err := c.Ctx.Input.ParseFormOrMulitForm(128 << 20)

	if err != nil {
		c.Ctx.Output.SetStatus(413)
		c.ServeJSON()
		return
	}

	if !c.Ctx.Input.IsUpload() {
		err := errors.New("Not image file found on request")
		c.BadRequest(err)
		return
	}

	token := c.Ctx.Input.Header("Authorization")

	decodedToken, err := VerifyToken(token, "Watcher")

	if err != nil {
		c.BadRequest(err)
		return
	}

	condoID, err := strconv.Atoi(decodedToken.CondoID)
	if err != nil {
		c.BadRequest(err)
		return
	}

	_, err = models.GetCondosByID(condoID)
	if err != nil {
		c.BadRequestDontExists("Condos")
		return
	}

	/* 	var v models.Checks

	   	// Validate empty body

	   	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	   	if err != nil {
	   		c.BadRequest(err)
	   		return
	   	} */

	var r = c.Ctx.Request

	v := models.Checks{
		Comment: r.FormValue("comment"),
		Date:    jodaTime.Format("Y-M-d HH:mm:ss", time.Now().In(orm.DefaultTimeLoc)),
	}

	valid := validation.Validation{}

	b, _ := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	watcherID, _ := strconv.Atoi(decodedToken.UserID)

	watcher, err := models.GetWatchersByID(watcherID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	v.Worker = watcher.Worker

	_, faceFh, err := c.GetFile("faces")

	if err != nil {
		c.BadRequest(err)
		return
	}

	_, ok, err := VerifyWorkerIdentity(watcher.Worker.ID, faceFh)

	if err != nil {
		c.BadRequest(err)
		return
	}

	if !ok {
		err = errors.New("Identity Verification Failed")
		c.BadRequest(err)
		return
	}

	_, err = models.AddChecks(&v)

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
// @Description get Checks by id
// @router /:id [get]
func (c *ChecksController) GetOne() {

	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetChecksByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Checks
// @Accept json
// @Param Authorization header string true "Supervisor's Token"
// @Param id param int true "Task's id"
// @Success 200 {object} models.Checks
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Task's Dont exists
// @router /:id [put]
func (c *ChecksController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Checks{ID: id}

	// Validate empty body

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = models.UpdateChecksByID(&v)

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
// @Description delete the Checks
// @router /:id [delete]
func (c *ChecksController) Delete() {
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

	err = models.DeleteChecks(id, trash)

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
func (c *ChecksController) GetAllFromTrash() {

	v, err := models.GetChecksFromTrash()

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
func (c *ChecksController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Checks{ID: id}

	err = models.RestoreFromTrash(v.TableName(), v.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}
