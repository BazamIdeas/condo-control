package controllers

import (
	"condo-control/models"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/astaxie/beego/validation"
)

// GoalsController operations for Holidays
type GoalsController struct {
	BaseController
}

//URLMapping ...
func (c *GoalsController) URLMapping() {
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
func (c *GoalsController) Post() {

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

	var v models.Goals

	// Validate empty body

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	if v.Task == nil || v.Task.ID == 0 {
		err = errors.New("Task data is empty")
		c.BadRequest(err)
		return
	}

	_, err = models.GetTasksByID(v.Task.ID)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	v.Status = "pending"

	valid := validation.Validation{}

	b, _ := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	_, err = models.AddGoals(&v)

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
// @Description get Goals by id
// @router /:id [get]
func (c *GoalsController) GetOne() {

	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetGoalsByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

///////////////////

// Put ...
// @Title Put
// @Description update the Goals
// @Accept json
// @Param Authorization header string true "Supervisor's Token"
// @Param id param int true "Goal's id"
// @Success 200 {object} models.Goals
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Goal's Dont exists
// @router /:id [put]
func (c *GoalsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Goals{ID: id}

	// Validate empty body

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = models.UpdateGoalsByID(&v)

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
// @Description delete the Goals
// @router /:id [delete]
func (c *GoalsController) Delete() {
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

	err = models.DeleteGoals(id, trash)

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
func (c *GoalsController) GetAllFromTrash() {

	v, err := models.GetGoalsFromTrash()

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
func (c *GoalsController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Tasks{ID: id}

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
// @Description Change Status
// @router /:id/status/:status [put]
func (c *GoalsController) ChangeStatus() {

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

	goal, err := models.GetGoalsByID(id)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	goal.Status = status

	err = models.UpdateGoalsByID(goal)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = goal
	c.ServeJSON()

}
