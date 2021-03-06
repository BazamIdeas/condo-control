package controllers

import (
	"condo-control/models"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego/validation"
)

// ZonesController operations for Zones
type ZonesController struct {
	BaseController
}

//URLMapping ...
func (c *ZonesController) URLMapping() {
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
// @Description create Zones
// @Accept json
// @Param   Authorization     header   string true       "Supervisor's Token"
// @Param   name     body   string true       "Supervisor's Token"
// @Success 200 {object} models.Zones
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Condos Don't Exists or Condos without Zones
// @Failure 409 Condo's Zone Limit reached
// @router / [post]
func (c *ZonesController) Post() {
	var v models.Zones

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
	decAuthToken, err := VerifyToken(authToken, "Supervisor")

	if err != nil {
		c.BadRequest(err)
		return
	}

	condoID, _ := strconv.Atoi(decAuthToken.CondoID)

	condo, err := models.GetCondosByID(condoID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	condosZonesCount := len(condo.Zones) + 1

	if condosZonesCount > condo.ZoneLimit {
		c.Ctx.Output.SetStatus(409)
		c.Data["json"] = MessageResponse{
			Code:    409,
			Message: "Condo's zones limit reached",
		}
		c.ServeJSON()
		return
	}

	v.Condo = &models.Condos{ID: condo.ID}

	_, err = models.AddZones(&v)

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
// @Description get Zones by id
// @router /:id [get]
/* func (c *ZonesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetZonesByID(id)
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
// @Param   Authorization     header   string true       "Watcher's Token or Supervisor's Token"
// @Success 200 {array} models.Zones
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Condos Don't Exists or Condos without Zones
// @router /self [get]
func (c *ZonesController) GetSelf() {

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

	c.Data["json"] = v.Zones
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Zones
// @router / [get]
/* func (c *ZonesController) GetAll() {
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

	l, err := models.GetAllZones(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()

} */

// Put ...
// @Title Put
// @Description update the Zones
// @Accept json
// @Param   Authorization     header   string true       "Supervisor's Token"
// @Param   id     param   string true       "Zone's id"
// @Param   name     body   string false       "Zone's new Name"
// @Success 200 {object} models.Zones
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Zones Don't exists
// @router /:id [put]
func (c *ZonesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Zones{ID: id}

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

	err = models.UpdateZonesByID(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Zones
// @router /:id [delete]
func (c *ZonesController) Delete() {
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

	err = models.DeleteZones(id, trash)

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
func (c *ZonesController) GetAllFromTrash() {

	v, err := models.GetZonesFromTrash()

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
func (c *ZonesController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Zones{ID: id}

	err = models.RestoreFromTrash(v.TableName(), v.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}
