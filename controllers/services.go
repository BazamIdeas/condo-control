package controllers

import (
	"GASE/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
)

// ServicesController operations for Services
type ServicesController struct {
	BaseController
}

// URLMapping ...
func (c *ServicesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetPricesServiceByCountry", c.GetPricesServiceByCountry)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Services
// @Param	body		body 	models.Services	true		"body for Services content"
// @Success 201 {int} models.Services
// @Failure 400 body is empty
// @router / [post]
func (c *ServicesController) Post() {
	var v models.Services

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	valid := validation.Validation{}

	b, err := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, "Services")
	}

	_, err = models.AddServices(&v)

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
// @Description get Services by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Services
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ServicesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetServicesByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Services
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Services
// @Failure 403
// @router / [get]
func (c *ServicesController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	//TODO: QUERY FILTER POR PAIS Y MONEDAS

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

	l, err := models.GetAllServices(query, fields, sortby, order, offset, limit)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	CountryIso := c.Ctx.Input.Header("Country-Iso")

	if CountryIso != "" {

		country, err := models.GetCountriesByIso(CountryIso)

		if err != nil {
			c.ServeErrorJSON(err)
			return
		}

		for i, service := range l {

			s := service.(models.Services)
			prices := s.Prices

			for _, price := range prices {
				if country.Currency.ID == price.Currency.ID {

					s.Price = price
					l[i] = s
					break
				}
			}

		}
	}

	c.Data["json"] = l
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Services
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Services	true		"body for Services content"
// @Success 200 {object} models.Services
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ServicesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Services{ID: id}
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = models.UpdateServicesByID(&v)

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
// @Description delete the Services
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ServicesController) Delete() {
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

	err = models.DeleteServices(id, trash)

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

// GetPricesServiceByCountry ...
// @Title GetPricesServiceByCountry
// @Description get Prices Services by Country Iso
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Services
// @Failure 403 :id is empty
// @router /:id/prices [get]
func (c *ServicesController) GetPricesServiceByCountry() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	CountryIso := c.Ctx.Input.Header("Country-Iso")

	if CountryIso == "" {
		CountryIso = "US"
	}

	_, err = models.GetCountriesByIso(CountryIso)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	v := &models.Services{ID: id}

	err = v.GetPricesServicesByID(CountryIso)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}
