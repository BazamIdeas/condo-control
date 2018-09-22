package controllers

import (
	"GASE/models"
	"encoding/json"
	"errors"

	"github.com/astaxie/beego/validation"
	"github.com/globalsign/mgo/bson"
)

// PaymentsMethodsController definiton.
type PaymentsMethodsController struct {
	BaseController
}

// URLMapping ...
func (c *PaymentsMethodsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Get", c.GetOne)
	c.Mapping("GetOneByIsoAndGateway", c.GetOneByIsoAndGateway)
	c.Mapping("GetAll", c.GetAll)
}

// Post ...
// @Title Post
// @router / [post]
func (c *PaymentsMethodsController) Post() {
	var v models.PaymentsMethods

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	// Validate context body

	valid := validation.Validation{}

	b, err := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, "PaymentsMethods")
		return
	}

	// Validate Exists Country

	exists := models.ValidateExists("Countries", v.Country.ID)

	if !exists {
		c.BadRequestDontExists("Country")
		return
	}

	// Validate Gateway

	exists = models.ValidateExists("Gateways", v.Gateway.ID)

	if !exists {
		c.BadRequestDontExists("Gateway")
		return
	}

	// New Object

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
func (c *PaymentsMethodsController) GetOne() {
	v := models.PaymentsMethods{}

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

	err := v.GetPaymentsMethodsByID(idStr)

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
func (c *PaymentsMethodsController) GetAll() {
	var v models.PaymentsMethods

	PaymentsMethods, err := v.GetAllPaymentsMethods()

	if err != nil {
		c.BadRequest(err)
		return
	}

	if len(PaymentsMethods) == 0 {
		c.ServeErrorJSON(errors.New("No hubo resultados"))
		return
	}

	c.Data["json"] = PaymentsMethods
	c.ServeJSON()
}

// Put ...
// @Title Put
// @router /:id [put]
func (c *PaymentsMethodsController) Put() {
	var v models.PaymentsMethods

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
func (c *PaymentsMethodsController) Delete() {
	var v models.PaymentsMethods

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

// GetOneByIsoAndGateway ...
// @Title GetOneByIsoAndGateway
// @router /:iso/:gateway [get]
func (c *PaymentsMethodsController) GetOneByIsoAndGateway() {
	v := models.PaymentsMethods{}

	iso := c.Ctx.Input.Param(":id")

	if iso == "" {
		c.BadRequest(errors.New("El campo iso no púede ser vacio"))
		return
	}

	gateway := c.Ctx.Input.Param(":gateway")

	if gateway == "" {
		c.BadRequest(errors.New("El campo gateway no púede ser vacio"))
		return
	}

	err := v.GetByIsoAndGateway(iso, gateway)

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}
