package controllers

import (
	"condo-control/controllers/services/files"
	"condo-control/models"
	"encoding/json"
	"errors"
	"strconv"
)

// OccurrencesController operations for Holidays
type OccurrencesController struct {
	BaseController
}

//URLMapping ...
func (c *OccurrencesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetAllFromTrash", c.GetAllFromTrash)
	c.Mapping("RestoreFromTrash", c.RestoreFromTrash)
	c.Mapping("GetAttachmentByUUID", c.GetAttachmentByUUID)
}

// Post ...
// @Title Post
// @Description create Occurrences
// @router / [post]
func (c *OccurrencesController) Post() {

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

	watcherID, _ := strconv.Atoi(decodedToken.UserID)

	watcher, err := models.GetWatchersByID(watcherID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	var r = c.Ctx.Request

	v := models.Occurrences{
		Comment: r.FormValue("comment"),
	}

	if v.Comment == "" {
		err = errors.New("Comment data is missing")
		c.BadRequest(err)
		return
	}

	var object models.Objects
	err = json.Unmarshal([]byte(r.FormValue("objects")), &object)
	if err != nil {
		c.BadRequest(err)
		return
	}

	var check models.Checks
	err = json.Unmarshal([]byte(r.FormValue("checks")), &check)
	if err != nil {
		c.BadRequest(err)
		return
	}

	v.Check = &check
	v.Object = &object

	if v.Object == nil || v.Object.ID == 0 {
		err = errors.New("object data is missing")
		c.BadRequest(err)
		return
	}

	if v.Check == nil || v.Check.ID == 0 {
		err = errors.New("check data is missing")
		c.BadRequest(err)
		return
	}

	_, fileFh, err := c.GetFile("files")

	if err == nil {
		fileUUID, mimeType, err := files.CreateFile(fileFh, "occurrences")

		if err != nil {
			c.BadRequest(err)
			return
		}

		v.ImageUUID = fileUUID
		v.ImageMime = mimeType
	}

	_, err = models.AddOccurrences(&v)

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
// @Description get Occurrences by id
// @router /:id [get]
func (c *OccurrencesController) GetOne() {

	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetOccurrencesByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Occurrences
// @Accept json
// @Param Authorization header string true "Supervisor's Token"
// @Param id param int true "Task's id"
// @Success 200 {object} models.Occurrences
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Task's Dont exists
// @router /:id [put]
func (c *OccurrencesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Occurrences{ID: id}

	// Validate empty body

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = models.UpdateOccurrencesByID(&v)

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
// @Description delete the Occurrences
// @router /:id [delete]
func (c *OccurrencesController) Delete() {
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

	err = models.DeleteOccurrences(id, trash)

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
func (c *OccurrencesController) GetAllFromTrash() {

	v, err := models.GetOccurrencesFromTrash()

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
func (c *OccurrencesController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Occurrences{ID: id}

	err = models.RestoreFromTrash(v.TableName(), v.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}

// GetAttachmentByUUID ...
// @Title Get  By UUID
// @Description Get file By UUID
// @router /attachment/:uuid [get]
func (c *OccurrencesController) GetAttachmentByUUID() {

	uuid := c.Ctx.Input.Param(":uuid")

	if uuid == "" {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte{})
		return
	}

	imageBytes, mimeType, err := files.GetFile(uuid, "occurrences")
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte{})
		return
	}

	c.Ctx.Output.Header("Content-Type", mimeType)
	c.Ctx.Output.SetStatus(200)
	c.Ctx.Output.Body(imageBytes)

}
