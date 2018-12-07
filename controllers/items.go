package controllers

import (
	"condo-control/controllers/services/files"
	"condo-control/models"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/vjeantet/jodaTime"
)

// ItemsController operations for Holidays
type ItemsController struct {
	BaseController
}

//URLMapping ...
func (c *ItemsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("ChangeStatus", c.ChangeStatus)
	c.Mapping("GetAllFromTrash", c.GetAllFromTrash)
	c.Mapping("RestoreFromTrash", c.RestoreFromTrash)
	c.Mapping("MakeComment", c.MakeComment)
	c.Mapping("MakeCommentExecute", c.MakeCommentExecute)
	c.Mapping("GetFilesByUUID", c.GetFilesByUUID)
}

// Post ...
// @Title Post
// @Description create Tasks
// @router / [post]
func (c *ItemsController) Post() {

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

	_, err = models.GetCondosByID(condoID)
	if err != nil {
		c.BadRequestDontExists("Condos")
		return
	}

	var v models.Items

	// Validate empty body

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	if v.Delivery == nil || v.Delivery.ID == 0 {
		err = errors.New("Delivery data is empty")
		c.BadRequest(err)
		return
	}

	_, err = models.GetDeliveriesByID(v.Delivery.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	_, err = models.AddItems(&v)

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
// @Description get Tasks by id
// @router /:id [get]
func (c *ItemsController) GetOne() {

	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetItemsByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Tasks
// @Accept json
// @Param Authorization header string true "Supervisor's Token"
// @Param id param int true "Item's id"
// @Success 200 {object} models.Items
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Task's Dont exists
// @router /:id [put]
func (c *ItemsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Items{ID: id}

	// Validate empty body

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = models.UpdateItemsByID(&v, true)

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
// @Description delete the Items
// @router /:id [delete]
func (c *ItemsController) Delete() {
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

	err = models.DeleteItems(id, trash)

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
func (c *ItemsController) GetAllFromTrash() {

	v, err := models.GetItemsFromTrash()

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
func (c *ItemsController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Items{ID: id}

	err = models.RestoreFromTrash(v.TableName(), v.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}

// ChangeStatus ...
// @Title Change Status
// @Description ChangeStatus
// @router /:id/status/:delivered [put]
func (c *ItemsController) ChangeStatus() {

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

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	deliveredStr := c.Ctx.Input.Param(":delivered")

	var delivered bool

	switch deliveredStr {
	case "true":
		delivered = true
	case "false":
		delivered = false
	default:
		err = errors.New("Invalid Delivered value")
		c.BadRequest(err)
		return
	}

	item, err := models.GetItemsByID(id)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	item.Delivered = delivered

	if item.Delivered {
		item.DateEnd = jodaTime.Format("Y-M-d HH:mm:ss", time.Now())
	} else {
		item.DateEnd = ""
	}

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

	err = models.UpdateItemsByID(item, false)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = item
	c.ServeJSON()

}

// MakeComment ...
// @Title Make Comment
// @Description Make comment
// @router /:id/comment [put]
func (c *ItemsController) MakeComment() {

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

	watcherID, err := strconv.Atoi(decodedToken.UserID)
	if err != nil {
		c.BadRequest(err)
		return
	}

	_, err = models.GetWatchersByID(watcherID)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	item, err := models.GetItemsByID(id)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	var r = c.Ctx.Request

	comment := r.FormValue("comment")

	if comment != "" {
		item.Comment = comment
	}

	_, fileFh, err := c.GetFile("files")

	if err == nil {
		fileUUID, mimeType, err := files.CreateFile(fileFh, "items")

		if err != nil {
			c.BadRequest(err)
			return
		}
		item.ImageMime = mimeType
		item.ImageUUID = fileUUID
	}

	commentToken, err := GenerateGeneralToken(decodedToken.UserID, decodedToken.CondoID, nil, nil, item)

	if err != nil {
		c.BadRequest(err)
		return
	}

	item.Token = commentToken

	c.Data["json"] = item
	c.ServeJSON()

}

// MakeCommentExecute ...
// @Title Make Comment Execute
// @Description Make comment Execute
// @router /:id/comment/:token [put]
func (c *ItemsController) MakeCommentExecute() {

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

	commentToken := c.Ctx.Input.Param(":token")
	decCommentToken, err := VerifyGeneralToken(commentToken)
	if err != nil {
		c.BadRequest(err)
		return
	}

	item := decCommentToken.Item

	err = models.UpdateItemsByID(item, true)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}
}

// GetFilesByUUID ...
// @Title Get By UUID
// @Description Get file By UUID
// @router /image/:uuid [get]
func (c *ItemsController) GetFilesByUUID() {

	uuid := c.Ctx.Input.Param(":uuid")

	if uuid == "" {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte{})
		return
	}

	imageBytes, mimeType, err := files.GetFile(uuid, "items")
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte{})
		return
	}

	c.Ctx.Output.Header("Content-Type", mimeType)
	c.Ctx.Output.SetStatus(200)
	c.Ctx.Output.Body(imageBytes)

}
