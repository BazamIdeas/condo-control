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

// QuestionsController operations for Questions
type QuestionsController struct {
	BaseController
}

//URLMapping ...
func (c *QuestionsController) URLMapping() {
	c.Mapping("Post", c.Post)
	/* c.Mapping("GetOne", c.GetOne) */
	/* c.Mapping("GetAll", c.GetAll) */
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetSelf", c.GetSelf)
	c.Mapping("GetAllFromTrash", c.GetAllFromTrash)
	c.Mapping("RestoreFromTrash", c.RestoreFromTrash)
}

// Post ...
// @Title Post
// @Description create Questions
// @Accept json
// @Param   Authorization     header   string true       "Resident's Token"
// @Param   name     body   string true       "Resident's Token"
// @Success 200 {object} models.Questions
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Condos Don't Exists or Condos without Questions
// @Failure 409 Condo's Question Limit reached
// @router / [post]
func (c *QuestionsController) Post() {
	var v models.Questions

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

	authToken := c.Ctx.Input.Header("Authorization")
	decodedToken, userType, err := VerifyTokenByAllUserTypes(authToken)

	if err != nil {
		c.BadRequest(err)
		return
	}

	condoID, _ := strconv.Atoi(decodedToken.CondoID)

	condo, err := models.GetCondosByID(condoID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	v.Condo = &models.Condos{ID: condo.ID}

	userID, _ := strconv.Atoi(decodedToken.UserID)

	if userType == "Resident" {

		resident, err := models.GetResidentsByID(userID)

		if err != nil {
			c.ServeErrorJSON(err)
			return
		}

		if resident.Committee {
			v.Approved = true
		} else {
			v.Approved = false
			v.CommitteeOnly = false
		}
	}

	now := time.Now().In(orm.DefaultTimeLoc)

	v.Date = jodaTime.Format("Y-M-d HH:mm:ss", now)

	_, err = jodaTime.Parse("Y-M-d", v.DateEnd)

	if err != nil {
		err = errors.New("invalid date format")
		c.BadRequest(err)
		return
	}

	_, err = models.AddQuestions(&v)

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
// @Description get Questions by id
// @router /:id [get]
/* func (c *QuestionsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetQuestionsByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
} */

// GetSelf ...
// @Title Get Self
// @Description Get Self
// @Accept json
// @Param   Authorization     header   string true       "Watcher's Token or Resident's Token"
// @Success 200 {array} models.Questions
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Condos Don't Exists or Condos without Questions
// @router /self [get]
func (c *QuestionsController) GetSelf() {

	token := c.Ctx.Input.Header("Authorization")
	decodedToken, _, _ := VerifyTokenByAllUserTypes(token)

	id, err := strconv.Atoi(decodedToken.CondoID)
	if err != nil {
		c.BadRequest(err)
		return
	}
	//TODO: use another function, filter by deleted needed
	v, err := models.GetCondosByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	if v.Questions == nil || len(v.Questions) == 0 {
		err = orm.ErrNoRows
		c.BadRequest(err)
		return
	}

	c.Data["json"] = v.Questions
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Questions
// @router / [get]
/* func (c *QuestionsController) GetAll() {
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

	l, err := models.GetAllQuestions(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()

} */

// Put ...
// @Title Put
// @Description update the Questions
// @Accept json
// @Param   Authorization     header   string true       "Resident's Token"
// @Param   id     param   string true       "Question's id"
// @Param   name     body   string false       "Question's new Name"
// @Success 200 {object} models.Questions
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Questions Don't exists
// @router /:id [put]
func (c *QuestionsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Questions{ID: id}

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

	err = models.UpdateQuestionsByID(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Questions
// @router /:id [delete]
func (c *QuestionsController) Delete() {
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

	err = models.DeleteQuestions(id, trash)

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
// @router /trashed [get]
func (c *QuestionsController) GetAllFromTrash() {

	v, err := models.GetQuestionsFromTrash()

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
func (c *QuestionsController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Questions{ID: id}

	err = models.RestoreFromTrash(v.TableName(), v.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}
