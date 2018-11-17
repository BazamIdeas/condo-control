package controllers

import (
	"condo-control/models"
	"encoding/json"
	"errors"
	"strconv"
)

// DeliviriesController operations for Holidays
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

	_, err = models.GetDeliveriesByID(v.Deliveries.ID)

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
func (c *DeliveriesController) GetOne() {

	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetDeliveriesByID(id)
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
// @Param id param int true "Delivery's id"
// @Success 200 {object} models.Deliveries
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Task's Dont exists
// @router /:id [put]
func (c *DeliveriesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Deliveries{ID: id}

	// Validate empty body

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = models.UpdateDeliveriesByID(&v)

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
// @Description delete the Deliveries
// @router /:id [delete]
func (c *DeliveriesController) Delete() {
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

	err = models.DeleteDeliveries(id, trash)

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
func (c *DeliveriesController) GetAllFromTrash() {

	v, err := models.GetDeliveriesFromTrash()

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
func (c *DeliveriesController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Deliveries{ID: id}

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
// @router /:id/status/:status [put]
func (c *DeliveriesController) ChangeStatus() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	status := c.Ctx.Input.Param(":status")

	if status == "" || (status != "pending" && status != "completed" && status != "finished") {

		err = errors.New("Wrong status")
		c.BadRequest(err)
		return
	}

	delivery, err := models.GetDeliveriesByID(id)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	delivery.Status = status

	err = models.UpdateDeliveriesByID(delivery)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = delivery
	c.ServeJSON()

}
