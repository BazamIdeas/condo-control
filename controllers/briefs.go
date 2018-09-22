package controllers

import (
	"GASE/models"
	"encoding/json"
	"errors"

	"github.com/astaxie/beego/validation"
	"github.com/globalsign/mgo/bson"
)

// BriefsController definiton.
type BriefsController struct {
	BaseController
}

// URLMapping ...
func (c *BriefsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Get", c.GetOne)
	c.Mapping("GetOneByCookie", c.GetOneByCookie)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Options", c.Options)
}

// Post ...
// @router / [post]
func (c *BriefsController) Post() {

	var v models.Briefs

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	// Validate context body

	valid := validation.Validation{}

	b, err := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, "Briefs")
		return
	}

	v.ID = bson.NewObjectId()

	err = v.Insert()

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetOne ...
// @router /:id [get]
func (c *BriefsController) GetOne() {

	v := models.Briefs{}

	idStr := c.Ctx.Input.Param(":id")
	if idStr == "" {
		c.BadRequest(errors.New("El campo id no púede ser vacio"))
		return
	}

	validID := bson.IsObjectIdHex(idStr)
	if !validID {
		c.BadRequest(errors.New("El Id no es valido"))
		return
	}

	err := v.GetBriefsByID(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @router /:id [get]
func (c *BriefsController) GetAll() {
	//var v models.Briefs

	Briefs, err := models.GetAllBriefs()

	if err != nil {
		c.BadRequest(err)
		return
	}

	if len(Briefs) == 0 {
		c.ServeErrorJSON(errors.New("No hubo resultados"))
		return
	}

	c.Data["json"] = Briefs
	c.ServeJSON()
}

// Delete ...
// @router /:id [delete]
func (c *BriefsController) Delete() {
	var v models.Briefs

	idStr := c.Ctx.Input.Param(":id")

	if idStr == "" {
		c.BadRequest(errors.New("El campo id no púede ser vacio"))
		return
	}

	validID := bson.IsObjectIdHex(idStr)

	if !validID {
		c.BadRequest(errors.New("El Id no es valido"))
		return
	}

	err := v.Delete(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Data["json"] = MessageResponse{
		Message:       "Deleted element",
		PrettyMessage: "Elemento Eliminado",
	}
	c.ServeJSON()
}

// GetOneByCookie ...
// @router /:service/:cookie [get]
func (c *BriefsController) GetOneByCookie() {
	v := models.Briefs{}

	cookie := c.Ctx.Input.Param(":cookie")

	if cookie == "" {
		c.BadRequest(errors.New("El campo cookie no púede ser vacio"))
		return
	}

	err := v.GetBriefsByCookie(cookie)

	if err != nil {

		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}
