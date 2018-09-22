package controllers

import (
	"GASE/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
)

// ImagesController operations for Images
type ImagesController struct {
	BaseController
}

// URLMapping ...
func (c *ImagesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("ServeImageBySlug", c.ServeImageBySlug)
}

// Post ...
// @Title Post
// @Description create Images
// @Param	body		body 	models.Images	true		"body for Images content"
// @Success 201 {int} models.Images
// @Failure 400 body is empty
// @router / [post]
func (c *ImagesController) Post() {
	//var v models.Images

	//err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	/*
		if err != nil {
			c.BadRequest(err)
			return
		}

		valid := validation.Validation{}

		b, err := valid.Valid(&v)

		if !b {
			c.BadRequestErrors(valid.Errors, v.TableName())
			return
		}
	*/

	err := c.Ctx.Input.ParseFormOrMulitForm(128 << 20)

	if err != nil {
		c.Ctx.Output.SetStatus(413)
		c.ServeJSON()
		return
	}

	var r = c.Ctx.Request

	stringsInts := &map[string]string{
		"portfolio": r.FormValue("portfolio[id]"),
		"priority":  r.FormValue("priority"),
	}

	intValues, err := stringIsValidInt(stringsInts)

	if err != nil {
		c.BadRequest(err)
		return
	}

	foreignsModels := map[string]int{
		"Portfolios": intValues["portfolio"],
	}

	resume := c.doForeignModelsValidation(foreignsModels)

	if !resume {
		return
	}

	v := &models.Images{
		Priority:  int8(intValues["priority"]),
		Portfolio: &models.Portfolios{ID: intValues["portfolio"]},
	}

	portfolio, err := models.GetPortfoliosByID(v.Portfolio.ID)

	if err != nil {
		c.BadRequestDontExists("Portfolio dont exists")
		return
	}

	portfolio.Images = []*models.Images{}

	if !c.Ctx.Input.IsUpload() {
		err := errors.New("Not image file found on request")
		c.BadRequest(err)
		return
	}

	_, fileHeader, err := c.GetFile("images")

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err = addNewImage(fileHeader, v.Portfolio)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	err = generateImageURL(v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Images by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Images
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ImagesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetImagesByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	err = generateImageURL(v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Images
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Images
// @Failure 403
// @router / [get]
func (c *ImagesController) GetAll() {
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

	l, err := models.GetAllImages(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	for key, image := range l {

		i := image.(models.Images)

		err = generateImageURL(&i)

		if err != nil {
			c.BadRequest(err)
			return
		}

		l[key] = &i

	}

	c.Data["json"] = l
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Images
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Images	true		"body for Images content"
// @Success 200 {object} models.Images
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ImagesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Images{ID: id}
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	valid := validation.Validation{}

	b, err := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	foreignsModels := map[string]int{
		"Portfolios": v.Portfolio.ID,
	}

	resume := c.doForeignModelsValidation(foreignsModels)

	if !resume {
		return
	}

	err = models.UpdateImagesByID(&v)

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
// @Description delete the Images
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ImagesController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	trash := false

	if c.Ctx.Input.Query("trash") != "" {
		trash = true
	}

	err = models.DeleteImages(id, trash)

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

// ServeImageBySlug ...
// @Title Serve Image By Slug
// @Description Serve Images by Slug
// @router /slug/:slug [get]
func (c *ImagesController) ServeImageBySlug() {

	slug := c.Ctx.Input.Param(":slug")

	if slug == "" {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte{})
		return
	}

	valid := validation.Validation{}

	result := valid.AlphaDash(slug, "slug")

	if !result.Ok {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte{})
		return
	}
	v, err := models.GetImagesBySlug(slug)

	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte{})
		return
	}

	imageURL := imageFolderDir + "/" + v.UUID

	imageBytes, err := ioutil.ReadFile(imageURL)

	if err != nil {

		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte{})
		return
	}

	c.Ctx.Output.Header("Content-Type", v.Mimetype)
	c.Ctx.Output.Body(imageBytes)

}
