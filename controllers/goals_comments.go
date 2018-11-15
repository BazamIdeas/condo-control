package controllers

import (
	"condo-control/models"
	"errors"
	"strconv"
)

// GoalsCommentsController operations for Holidays
type GoalsCommentsController struct {
	BaseController
}

//URLMapping ...
func (c *GoalsCommentsController) URLMapping() {
	c.Mapping("Post", c.Post)
	/*c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)*/
}

// Post ...
// @Title Post
// @Description create Goals comments
// @router / [post]
func (c *GoalsCommentsController) Post() {

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

	decodedToken, userType, err := VerifyTokenByAllUserTypes(token)
	if err != nil {
		c.BadRequest(err)
		return
	}

	condoID, _ := strconv.Atoi(decodedToken.CondoID)

	_, err = models.GetCondosByID(condoID)
	if err != nil {
		c.BadRequestDontExists("Condos")
		return
	}

	userID, _ := strconv.Atoi(decodedToken.UserID)

	worker := &models.Workers{}

	switch userType {

	case "Watcher":
		watcher, err := models.GetWatchersByID(userID)
		if err != nil {
			c.ServeErrorJSON(err)
			return
		}

		worker = watcher.Worker
		break

	case "Supervisor":
		supervisor, err := models.GetSupervisorsByID(userID)
		if err != nil {
			c.ServeErrorJSON(err)
			return
		}

		worker = supervisor.Worker
		break

	}

	var r = c.Ctx.Request

	goalID, err := strconv.Atoi(r.FormValue("goal_id"))

	if err != nil {
		c.BadRequest(err)
		return
	}

	goal := &models.Goals{
		ID: goalID,
	}

	v := &models.GoalsComments{
		Description: r.FormValue("description"),
		Goal:        goal,
		Worker:      worker,
	}

	_, err = models.AddGoalsComments(v)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v

	c.ServeJSON()

}
