package controllers

import (
	"condo-control/controllers/services/files"
	"condo-control/models"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/astaxie/beego/validation"
)

// NotificationsController operations for Notifications
type NotificationsController struct {
	BaseController
}

//URLMapping ...
func (c *NotificationsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetImageByUUID", c.GetImageByUUID)
	c.Mapping("Approve", c.Approve)
	c.Mapping("View", c.View)
	c.Mapping("GetSelf", c.GetSelf)
}

// Post ...
// @Title Post
// @Description create Notifications
// @router / [post]
func (c *NotificationsController) Post() {

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

	var r = c.Ctx.Request

	v := &models.Notifications{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	}

	decodedToken, err := VerifyToken(token, "Watcher")

	if err != nil {
		c.BadRequest(err)
		return
	}

	watcherID, err := strconv.Atoi(decodedToken.UserID)
	if err != nil {
		c.BadRequest(err)
		return
	}

	watcher, err := models.GetWatchersByID(watcherID)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	v.Worker = watcher.Worker

	_, fileFh, err := c.GetFile("files")

	if err == nil {
		fileUUID, mimeType, err := files.CreateFile(fileFh, "notifications")

		if err != nil {
			c.BadRequest(err)
			return
		}

		v.ImageUUID = fileUUID
		v.ImageMime = mimeType
	}

	_, err = models.AddNotifications(v)

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
// @Description get Notifications by id
// @router /:id [get]
func (c *NotificationsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetNotificationsByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// Put ...
// @Title Put
// @router /:id [put]
func (c *NotificationsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Notifications{ID: id}

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

	err = models.UpdateNotificationsByID(&v)

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
// @Description delete the Notifications
// @router /:id [delete]
func (c *NotificationsController) Delete() {
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

	err = models.DeleteNotifications(id, trash)

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
func (c *NotificationsController) GetAllFromTrash() {

	v, err := models.GetNotificationsFromTrash()

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
func (c *NotificationsController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Notifications{ID: id}

	err = models.RestoreFromTrash(v.TableName(), v.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}

// GetImageByUUID ...
// @Title Get  By UUID
// @Description Get file By UUID
// @router /image/:uuid [get]
func (c *NotificationsController) GetImageByUUID() {

	uuid := c.Ctx.Input.Param(":uuid")

	if uuid == "" {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte{})
		return
	}

	imageBytes, mimeType, err := files.GetFile(uuid, "notifications")
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte{})
		return
	}

	c.Ctx.Output.Header("Content-Type", mimeType)
	c.Ctx.Output.SetStatus(200)
	c.Ctx.Output.Body(imageBytes)

}

// Approve ...
// @Title Approve
// @router /:id/status/:approved [put]
func (c *NotificationsController) Approve() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	notification, err := models.GetNotificationsByID(id)
	if err != nil {

		c.BadRequest(err)
		return
	}

	approvedStr := c.Ctx.Input.Param(":approved")

	var approved bool

	switch approvedStr {
	case "true":
		approved = true
	case "false":
		approved = false
	default:
		err = errors.New("Invalid Approved value")
		c.BadRequest(err)
		return
	}

	notification.Approved = approved

	err = models.UpdateNotificationsByID(notification)
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

// View ...
// @Title View
// @router /:id/view/:viewed [put]
func (c *NotificationsController) View() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	notification, err := models.GetNotificationsByID(id)
	if err != nil {
		c.BadRequest(err)
		return
	}

	viewedStr := c.Ctx.Input.Param(":viewed")

	var viewed bool

	switch viewedStr {
	case "true":
		viewed = true
	case "false":
		viewed = false
	default:
		err = errors.New("Invalid Approved value")
		c.BadRequest(err)
		return
	}

	notification.View = viewed

	err = models.UpdateNotificationsByID(notification)
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

// GetByCondosSelf ...
// @Title Get By Condos Self
// @router /condos/self [get]
func (c *NotificationsController) GetByCondosSelf() {

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

	condo, err := models.GetCondosByID(condoID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	notifications, err := models.GetNotificationsByCondosID(condo.ID)

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Data["json"] = notifications
	c.ServeJSON()
}

// GetSelf ...
// @Title Get Self
// @router /self [get]
func (c *NotificationsController) GetSelf() {

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
		c.ServeErrorJSON(err)
		return
	}

	watcherID, err := strconv.Atoi(decodedToken.UserID)
	if err != nil {
		c.BadRequest(err)
		return
	}

	watcher, err := models.GetWatchersByID(watcherID)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	worker, err := models.GetWorkersByID(watcher.Worker.ID)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = worker.Notifications
	c.ServeJSON()
}
