package controllers

import (
	"GASE/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
)

// CartsController operations for Carts
type CartsController struct {
	BaseController
}

// URLMapping ...
func (c *CartsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("AddOrDeleteServices", c.AddOrDeleteServices)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetOneByCookie", c.GetOneByCookie)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Carts
// @Param	body		body 	models.Carts	true		"body for Carts content"
// @Success 201 {int} models.Carts
// @Failure 400 body is empty
// @router / [post]
func (c *CartsController) Post() {
	var v models.Carts

	// Validate empty body

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	// Validate Format Types
	valid := validation.Validation{}

	b, err := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	// Validate Services exists
	var servicesRelations []map[string]int

	for _, el := range v.Services {

		exists := models.ValidateExists("Services", el.ID)

		if !exists {
			c.BadRequestDontExists("Service")
			return
		}

		relation := map[string]int{"id": el.ID, "quantity": el.Quantity}

		servicesRelations = append(servicesRelations, relation)
	}

	_, err = models.AddCarts(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	// Add Services relations
	_, err = models.AddRelationsM2MQuantity("carts", v.ID, "services", servicesRelations)

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
// @Description get Carts by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Carts
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CartsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetCartsByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Carts
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Carts
// @Failure 403
// @router / [get]
func (c *CartsController) GetAll() {
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

	l, err := models.GetAllCarts(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Carts
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Carts	true		"body for Carts content"
// @Success 200 {object} models.Carts
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CartsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Carts{ID: id}

	// Validate context body
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = models.UpdateCartsByID(&v)

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
// @Description delete the Carts
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CartsController) Delete() {
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

	err = models.DeleteCarts(id, trash)

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

// GetOneByCookie ...
// @Title Get One
// @Description get Carts by Cookie
// @Param	Cookie		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Carts
// @Failure 403 :cookie is empty
// @router /cookie [get]
func (c *CartsController) GetOneByCookie() {

	//Validate Cookie
	cookie := c.Ctx.Input.Header("Cart-Id")

	if cookie == "" {
		err := errors.New("No se ha recibido el id del Carrito")
		c.BadRequest(err)
		return
	}

	header := c.Ctx.Input.Header("Country-Iso")

	if header == "" {
		header = "US"
	}

	_, err := models.GetCountriesByIso(header)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	// Validate Cart

	v, err := models.GetOrCreateCartsByCookie(cookie, header)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// AddOrDeleteServices ...
// @Title AddOrDeleteServices to Carts
// @Description AddOrDeleteServices to Cart
// @Param	body		body 	models.Services	true		"body for Services content"
// @Success 201 {int} models.Carts
// @Failure 400 body is empty
// @router /services [put]
func (c *CartsController) AddOrDeleteServices() {

	//Validate action

	action := c.Ctx.Input.Query("action")

	if action == "" && action != "add" && action != "remove" {
		err := errors.New("No se ha recibido la accion a realizar")
		c.BadRequest(err)
		return
	}

	//Validate Cookie

	cookie := c.Ctx.Input.Header("Cart-Id")

	if cookie == "" {
		err := errors.New("No se ha recibido el id del Carrito")
		c.BadRequest(err)
		return
	}

	// Validate Country

	header := c.Ctx.Input.Header("Country-Iso")

	if header == "" {
		header = "US"
	}

	_, err := models.GetCountriesByIso(header)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	cart, err := models.GetOrCreateCartsByCookie(cookie, header)

	// Validate empty body

	var v models.Carts
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v.ID = cart.ID

	// Validate Services exists
	var servicesRelations []map[string]int
	var servicesRelationsIDs []int

	for _, el := range v.Services {

		exists := models.ValidateExists("Services", el.ID)

		if !exists {
			c.BadRequestDontExists("Service")
			return
		}

		if action == "add" {

			relation := map[string]int{"id": el.ID, "quantity": el.Quantity}
			servicesRelations = append(servicesRelations, relation)

		} else {

			servicesRelationsIDs = append(servicesRelationsIDs, el.ID)

		}

	}

	if action == "add" {

		// Insert Services relations
		_, err = models.AddRelationsM2MQuantity("carts", v.ID, "services", servicesRelations)

		if err != nil {
			c.ServeErrorJSON(err)
			return
		}

		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v

	} else {

		// Delete Services relations
		count, err := models.RelationsM2M("DELETE", "carts", v.ID, "services", servicesRelationsIDs)

		if err != nil {
			c.ServeErrorJSON(err)
			return
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = MessageResponse{
			Message:       "Deleted relations: " + strconv.Itoa(count),
			PrettyMessage: "Relaciones Eliminadas: " + strconv.Itoa(count),
		}
	}

	c.ServeJSON()

}
