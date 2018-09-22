package controllers

import (
	"GASE/controllers/services/payments"
	"GASE/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
)

// GatewaysController operations for Gateways
type GatewaysController struct {
	BaseController
}

// URLMapping ...
func (c *GatewaysController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("AddNewsCurrencies", c.AddNewsCurrencies)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("DeleteCurrencies", c.DeleteCurrencies)
}

// Post ...
// @Title Post
// @Description create Gateways
// @Param	body		body 	models.Gateways	true		"body for Gateways content"
// @Success 201 {int} models.Gateways
// @Failure 400 body is empty
// @router / [post]
func (c *GatewaysController) Post() {
	var v models.Gateways

	// Validate empty body

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {

		c.BadRequest(err)
		return
	}

	// Validate context body

	valid := validation.Validation{}

	b, err := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	// Validate currencies exists
	var relationsIDs []int

	for _, el := range v.Currencies {

		exists := models.ValidateExists("currencies", el.ID)

		if !exists {
			c.BadRequestDontExists("currenccy")
			return
		}

		relationsIDs = append(relationsIDs, el.ID)
	}

	_, err = models.AddGateways(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	_, err = models.RelationsM2M("INSERT", "gateways", v.ID, "currencies", relationsIDs)

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
// @Description get Gateways by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Gateways
// @Failure 403 :id is empty
// @router /:id [get]
func (c *GatewaysController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetGatewaysByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Gateways
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Gateways
// @Failure 403
// @router / [get]
func (c *GatewaysController) GetAll() {
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

	l, err := models.GetAllGateways(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Gateways
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Gateways	true		"body for Gateways content"
// @Success 200 {object} models.Gateways
// @Failure 403 :id is not int
// @router /:id [put]
func (c *GatewaysController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Gateways{ID: id}

	// Validate context body
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = models.UpdateGatewaysByID(&v)

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
// @Description delete the Gateways
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *GatewaysController) Delete() {
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

	err = models.DeleteGateways(id, trash)

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

// AddNewsCurrencies ...
// @Title AddNewsCurrencies to Gateway
// @Description create Gateways
// @Param	body		body 	models.Gateways	true		"body for Gateways content"
// @Success 201 {int} models.Gateways
// @Failure 400 body is empty
// @router /:id/currencies [post]
func (c *GatewaysController) AddNewsCurrencies() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Gateways{ID: id}

	// Validate context body
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {

		c.ServeErrorJSON(err)
		return
	}

	// Validate currencies exists
	var relationsIDs []int

	for _, el := range v.Currencies {

		exists := models.ValidateExists("currencies", el.ID)

		if !exists {
			c.BadRequestDontExists("currenccy")
			return
		}

		relationsIDs = append(relationsIDs, el.ID)
	}

	count, err := models.RelationsM2M("INSERT", "gateways", v.ID, "currencies", relationsIDs)

	//TODO:

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = MessageResponse{
		Message:       "Added relations: " + strconv.Itoa(count),
		PrettyMessage: "Relaciones Agregadas: " + strconv.Itoa(count),
	}

	c.ServeJSON()
}

// DeleteCurrencies ...
// @Title DeleteCurrencies to Gateway
// @Description Delete Gateways relations M2M
// @Param	body		body 	models.Gateways	true		"body for Gateways content"
// @Success 201 {int} models.Gateways
// @Failure 400 body is empty
// @router /:id/currencies [delete]
func (c *GatewaysController) DeleteCurrencies() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Gateways{ID: id}

	// Validate context body
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {

		c.ServeErrorJSON(err)
		return
	}

	// Validate currencies exists
	var relationsIDs []int

	for _, el := range v.Currencies {

		exists := models.ValidateExists("currencies", el.ID)

		if !exists {
			c.BadRequestDontExists("currenccy")
			return
		}

		relationsIDs = append(relationsIDs, el.ID)
	}

	count, err := models.RelationsM2M("DELETE", "gateways", v.ID, "currencies", relationsIDs)

	if err != nil {

		c.ServeErrorJSON(err)
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = MessageResponse{
		Message:       "Deleted relations: " + strconv.Itoa(count),
		PrettyMessage: "Relaciones Eliminadas: " + strconv.Itoa(count),
	}

	c.ServeJSON()
}

///////////////////////////////////
/////////PAYMENT HANDLER///////////
///////////////////////////////////

func paymentsHandler(orderID int, gateway *models.Gateways, price float32, countries *models.Countries, paymentData map[string]interface{}) (redirect string, paymentID string, err error) {

	switch gateway.Code {

	case "01": //paypal [paypal checkout with button]

		if value, ok := paymentData["id"]; !ok || value.(string) == "" {

			err = errors.New("payment id is missing")

			return
		}

		paypalID := paymentData["id"].(string)

		payment := &payments.WebCheckoutPayment{}

		payment.ID = paypalID

		err = payment.ButtonCheckoutPaypal()

		if err != nil {
			return
		}

		paymentID = payment.ID
		break

	case "02": //stripe [credit card with element.js]

		if tokenValue, ok := paymentData["token"]; !ok || tokenValue.(string) == "" {

			err = errors.New("token is missing")

			return
		}

		tokenStripe := paymentData["token"].(string)

		payment := &payments.CreditCardPayment{}
		payment.Token = tokenStripe

		err = payment.CreditCardStripe()

		if err != nil {
			return
		}

		paymentID = payment.ID
		break

	case "03": //payU

		//TODO: Validar campos

		if tokenValue, ok := paymentData["token"]; !ok || tokenValue.(string) == "" {

			/* 	if  tokenValue, ok := paymentData["token"];

			 */

			return
		}

		tokenPayu := paymentData["token"].(string)

		payment := &payments.CreditCardPayment{}
		payment.Token = tokenPayu

		payment.ExtraData = paymentData

		err = payment.PayuCreditCard()

		if err != nil {
			return
		}

		paymentID = payment.ID

		break
	case "04":

	}

	return
}
