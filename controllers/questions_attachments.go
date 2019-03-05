package controllers

import (
	"condo-control/controllers/services/files"
	"condo-control/models"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/astaxie/beego/validation"
)

// QuestionsAttachmentsController operations for QuestionsAttachments
type QuestionsAttachmentsController struct {
	BaseController
}

//URLMapping ...
func (c *QuestionsAttachmentsController) URLMapping() {
	c.Mapping("Post", c.Post)
	/* c.Mapping("GetOne", c.GetOne) */
	/* c.Mapping("GetAll", c.GetAll) */
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetAllFromTrash", c.GetAllFromTrash)
	c.Mapping("RestoreFromTrash", c.RestoreFromTrash)
}

// Post ...
// @Title Post
// @Description create QuestionsAttachments
// @Accept json
// @Param   Authorization     header   string true       "Resident's Token"
// @Param   name     body   string true       "Resident's Token"
// @Success 200 {object} models.QuestionsAttachments
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Condos Don't Exists or Condos without QuestionsAttachments
// @Failure 409 Condo's QuestionAttachment Limit reached
// @router / [post]
func (c *QuestionsAttachmentsController) Post() {

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

	authToken := c.Ctx.Input.Header("Authorization")
	decodedToken, _, err := VerifyTokenByAllUserTypes(authToken)

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

	var r = c.Ctx.Request

	questionID, err := strconv.Atoi(r.FormValue("question_id"))

	if err != nil {
		c.BadRequest(err)
		return
	}

	question := &models.Questions{
		ID: questionID,
	}

	v := &models.QuestionsAttachments{
		Name:     r.FormValue("name"),
		Question: question,
	}

	_, fileFh, err := c.GetFile("files")

	if err == nil {
		fileUUID, mimeType, err := files.CreateFile(fileFh, "questions-attachments")

		if err != nil {
			c.BadRequest(err)
			return
		}

		v.AttachmentUUID = fileUUID
		v.AttachmentMime = mimeType
	}

	_, err = models.AddQuestionsAttachments(v)
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
// @Description get QuestionsAttachments by id
// @router /:id [get]
/* func (c *QuestionsAttachmentsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetQuestionsAttachmentsByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
} */

// GetAll ...
// @Title Get All
// @Description get QuestionsAttachments
// @router / [get]
/* func (c *QuestionsAttachmentsController) GetAll() {
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

	l, err := models.GetAllQuestionsAttachments(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()

} */

// Put ...
// @Title Put
// @Description update the QuestionsAttachments
// @Accept json
// @Param   Authorization     header   string true       "Resident's Token"
// @Param   id     param   string true       "QuestionAttachment's id"
// @Param   name     body   string false       "QuestionAttachment's new Name"
// @Success 200 {object} models.QuestionsAttachments
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 QuestionsAttachments Don't exists
// @router /:id [put]
func (c *QuestionsAttachmentsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.QuestionsAttachments{ID: id}

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

	err = models.UpdateQuestionsAttachmentsByID(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the QuestionsAttachments
// @router /:id [delete]
func (c *QuestionsAttachmentsController) Delete() {
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

	err = models.DeleteQuestionsAttachments(id, trash)

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
func (c *QuestionsAttachmentsController) GetAllFromTrash() {

	v, err := models.GetQuestionsAttachmentsFromTrash()

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
func (c *QuestionsAttachmentsController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.QuestionsAttachments{ID: id}

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
func (c *QuestionsAttachmentsController) GetAttachmentByUUID() {

	uuid := c.Ctx.Input.Param(":uuid")

	if uuid == "" {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte{})
		return
	}

	imageBytes, mimeType, err := files.GetFile(uuid, "questions-attachments")
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte{})
		return
	}

	c.Ctx.Output.Header("Content-Type", mimeType)
	c.Ctx.Output.SetStatus(200)
	c.Ctx.Output.Body(imageBytes)

}