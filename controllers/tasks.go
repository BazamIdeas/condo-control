package controllers

import (
	"condo-control/models"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/vjeantet/jodaTime"

	"github.com/astaxie/beego/validation"
)

// TasksController operations for Holidays
type TasksController struct {
	BaseController
}

//URLMapping ...
func (c *TasksController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetSelf", c.GetSelf)
	c.Mapping("GetByCondosID", c.GetByCondosID)
	c.Mapping("GetByWorkersID", c.GetByWorkersID)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("ChangeStatus", c.ChangeStatus)
	c.Mapping("GetAllFromTrash", c.GetAllFromTrash)
	c.Mapping("RestoreFromTrash", c.RestoreFromTrash)
}

// GetByCondosID ...
// @Title Get By CondosID
// @Description Get By CondosID
// @router /condos/self [get]
func (c *TasksController) GetByCondosID() {

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
		c.ServeErrorJSON(err)
		return
	}

	v := []*models.Tasks{}

	for _, worker := range condos.Workers {

		workerTasks, err := models.GetTasksByWorkersID(worker.ID)

		if err != nil && err != orm.ErrNoRows {
			c.BadRequest(err)
			return
		}

		if err == orm.ErrNoRows {
			continue
		}

		v = append(v, workerTasks...)

	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = v

	c.ServeJSON()

}

// GetByWorkersID ...
// @Title Get By WorkersID
// @Description Get By WorkersID
// @router /workers/:id [get]
func (c *TasksController) GetByWorkersID() {

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
		c.ServeErrorJSON(err)
		return
	}

	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	worker, err := models.GetWorkersByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	workerTasks, err := models.GetTasksByWorkersID(worker.ID)
	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = workerTasks

	c.ServeJSON()

}

// GetSelf ...
// @Title Get Self
// @Description Get Self
// @router /self [get]
func (c *TasksController) GetSelf() {

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

	id, err := strconv.Atoi(decodedToken.UserID)

	if err != nil {
		c.BadRequest(err)
		return
	}

	watcher, err := models.GetWatchersByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	worker, err := models.GetWorkersByID(watcher.Worker.ID)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	workerTasks, err := models.GetTasksByWorkersID(worker.ID)
	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = workerTasks

	c.ServeJSON()

}

// Post ...
// @Title Post
// @Description create Tasks
// @router / [post]
func (c *TasksController) Post() {

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

	var v models.Tasks

	// Validate empty body

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	if v.Worker == nil || v.Worker.ID == 0 {
		err = errors.New("Worker data is empty")
		c.BadRequest(err)
		return
	}

	_, err = models.GetWorkersByID(v.Worker.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	valid := validation.Validation{}

	b, _ := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	_, err = models.AddTasks(&v)

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
func (c *TasksController) GetOne() {

	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetTasksByID(id)
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
// @Param id param int true "Task's id"
// @Success 200 {object} models.Tasks
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Task's Dont exists
// @router /:id [put]
func (c *TasksController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Tasks{ID: id}

	// Validate empty body

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = models.UpdateTasksByID(&v, true)

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
// @Description delete the Tasks
// @router /:id [delete]
func (c *TasksController) Delete() {
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

	err = models.DeleteTasks(id, trash)

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
func (c *TasksController) GetAllFromTrash() {

	v, err := models.GetTasksFromTrash()

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
func (c *TasksController) RestoreFromTrash() {

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
// @Description ChangeStatus
// @router /:id/status/:approved [put]
func (c *TasksController) ChangeStatus() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	approvedStr := c.Ctx.Input.Param(":approved")

	var approved bool

	switch approvedStr {
	case "true":
		approved = true
		break

	case "false":
		approved = false
		break
	default:
		err = errors.New("Invalid Approved value")
		c.BadRequest(err)
		return
	}

	task, err := models.GetTasksByID(id)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	task.Approved = approved

	if task.Approved {
		task.DateEnd = jodaTime.Format("Y-M-d HH:mm:ss", time.Now())
	} else {
		task.DateEnd = ""
	}

	err = models.UpdateTasksByID(task, false)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = task
	c.ServeJSON()

}
