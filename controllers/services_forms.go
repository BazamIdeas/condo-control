package controllers

import (
	"GASE/models"
	"encoding/json"
	"errors"

	"github.com/astaxie/beego/validation"
	"github.com/globalsign/mgo/bson"
)

// ServiceFormsController definiton.
type ServiceFormsController struct {
	BaseController
}

// URLMapping ...
func (c *ServiceFormsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Get", c.GetOne)
	c.Mapping("GetOneByService", c.GetOneByService)
	c.Mapping("GetAll", c.GetAll)
}

// Post ...
// @Title Post
// @router / [post]
func (c *ServiceFormsController) Post() {
	var v models.ServiceForms

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	// Validate context body

	valid := validation.Validation{}

	b, err := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, "ServiceForms")
		return
	}

	// Validate Service

	exists := models.ValidateExists("Services", v.Service.ID)

	if !exists {
		c.BadRequestDontExists("Service")
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
// @Title Get One
// @router /:id [get]
func (c *ServiceFormsController) GetOne() {
	v := models.ServiceForms{}

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

	err := v.GetServiceFormsByID(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @router / [get]
func (c *ServiceFormsController) GetAll() {
	var v models.ServiceForms

	ServiceForms, err := v.GetAllServiceForms()

	if err != nil {
		c.BadRequest(err)
		return
	}

	if len(ServiceForms) == 0 {
		c.ServeErrorJSON(errors.New("No hubo resultados"))
		return
	}

	c.Data["json"] = ServiceForms
	c.ServeJSON()
}

// Put ...
// @Title Put
// @router /:id [put]
func (c *ServiceFormsController) Put() {
	var v models.ServiceForms

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

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = v.Update()

	if err != nil {
		c.BadRequest(err)
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
// @router /:id [delete]
func (c *ServiceFormsController) Delete() {
	var v models.ServiceForms

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

// GetOneByService ...
// @Title Get One
// @router /service/:slug [get]
func (c *ServiceFormsController) GetOneByService() {
	v := models.ServiceForms{}

	slug := c.Ctx.Input.Param(":slug")

	if slug == "" {
		c.BadRequest(errors.New("El campo slug no púede ser vacio"))
		return
	}

	err := v.GetServiceFormsByServiceSlug(slug)

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}
