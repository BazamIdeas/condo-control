package controllers

import (
	"condo-control/controllers/services/files"
	"condo-control/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/vjeantet/jodaTime"
)

// VerificationsController operations for Verifications
type VerificationsController struct {
	BaseController
}

//URLMapping ...
func (c *VerificationsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("NewRouteExecute", c.NewRouteExecute)
	c.Mapping("NewRoute", c.NewRoute)
	c.Mapping("GetImagesByUUID", c.GetImagesByUUID)
	c.Mapping("AddImage", c.AddImage)
	c.Mapping("Comment", c.Comment)
}

// Post ...
// @Title Post
// @Description create Verifications
// @router / [post]
func (c *VerificationsController) Post() {
	var v models.Verifications

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

	//TODO:
	// Validate foreings keys
	/*
		exists := models.ValidateExists("Sectors", v.Sector.ID)

		if !exists {
			c.BadRequestDontExists("Sector")
			return
		} */

	_, err = models.AddVerifications(&v)

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
// @Description get Verifications by id
// @router /:id [get]
func (c *VerificationsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetVerificationsByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Verifications
// @router / [get]
func (c *VerificationsController) GetAll() {
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

	l, err := models.GetAllVerifications(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Verifications
// @router /:id [put]
func (c *VerificationsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Verifications{ID: id}

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

	err = models.UpdateVerificationsByID(&v)

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
// @Description delete the Verifications
// @router /:id [delete]
func (c *VerificationsController) Delete() {
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

	err = models.DeleteVerifications(id, trash)

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
// @router /trashed [patch]
func (c *VerificationsController) GetAllFromTrash() {

	v, err := models.GetVerificationsFromTrash()

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
func (c *VerificationsController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Verifications{ID: id}

	err = models.RestoreFromTrash(v.TableName(), v.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}

// NewRoute ...
// @Title New Route
// @Description New Route
// @router /zones/:id/route [post]
func (c *VerificationsController) NewRoute() {

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

	z, err := models.GetZonesByID(v.ID)

	if err != nil {
		c.BadRequestDontExists("Zones")
		return
	}

	token := c.Ctx.Input.Header("Authorization")

	decodedToken, _ := VerifyToken(token, "Watcher")

	/* if err != nil {
		c.BadRequest(err)
		return
	} */

	condoID, err := strconv.Atoi(decodedToken.CondoID)

	if err != nil {
		c.BadRequest(err)
		return
	}

	if z.Condo.ID != condoID {
		err = errors.New("Zone dont belong to Watcher's Condo")
		c.BadRequest(err)
		return
	}

	now := time.Now()

	if v.Points == nil {

		err = errors.New("Empty Points")
		c.BadRequest(err)
		return
	}

	for _, point := range v.Points {

		p, err := models.GetPointsByID(point.ID)

		if err != nil {
			c.BadRequestDontExists("Points")
			return
		}

		if p.Zone.ID != v.ID {
			err = errors.New("Point's Zone is wrong")
			c.BadRequest(err)
			return
		}

		if point.Verifications == nil {

			err = errors.New("Empty Verifications")
			c.BadRequest(err)
			return
		}

		for _, verification := range point.Verifications {

			date, err := jodaTime.Parse("Y-M-d HH:mm:ss", verification.Date)

			if err != nil {
				c.BadRequest(err)
				return
			}

			if !date.Before(now) {
				err = errors.New("Verification date is a future Date")
				c.BadRequest(err)
				return
			}

			dur := now.Sub(date)

			if dur.Hours() > 6 {
				err = errors.New("Verification date is too old")
				c.BadRequest(err)
				return
			}

			//verification.Point = p

		}

	}

	routeToken, err := GenerateGeneralToken(decodedToken.UserID, decodedToken.CondoID, v.Points, nil, nil)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v.Token = routeToken

	c.Data["json"] = v
	c.ServeJSON()
}

// NewRouteExecute ...
// @Title New Route Execute
// @Description New Route Execute
// @router /route/:token [post]
func (c *VerificationsController) NewRouteExecute() {

	authToken := c.Ctx.Input.Header("Authorization")

	decAuthToken, _ := VerifyToken(authToken, "Watcher")

	routeToken := c.Ctx.Input.Param(":token")

	decRouteToken, err := VerifyGeneralToken(routeToken)

	if err != nil {
		c.BadRequest(err)
		return
	}

	if decRouteToken.CondoID != decAuthToken.CondoID {
		err = errors.New("Zone's Condo and Watcher's Condo Don't match")
		c.BadRequest(err)
		return
	}

	watcherID, _ := strconv.Atoi(decAuthToken.UserID)

	watcher, err := models.GetWatchersByID(watcherID)

	if err != nil {
		c.BadRequestDontExists("Watchers")
		return
	}

	_, faceFh, err := c.GetFile("faces")

	if err != nil {
		c.BadRequest(err)
		return
	}

	_, ok, err := VerifyWorkerIdentity(watcher.Worker.ID, faceFh)

	if err != nil {
		c.BadRequest(err)
		return
	}

	if !ok {
		err = errors.New("Identity Verification Failed")
		c.BadRequest(err)
		return
	}

	verifications := []*models.Verifications{}

	points := decRouteToken.Points

	for _, point := range points {
		for _, verification := range point.Verifications {
			p := &models.Points{ID: point.ID}
			verification.Point = p
			verification.Watcher = watcher
			verification.SupervisorComment = ""
			verifications = append(verifications, verification)

			_, err = models.AddVerifications(verification)
			if err != nil {
				c.ServeErrorJSON(err)
				return
			}
		}
	}

	//err = models.AddManyVerifications(verifications)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = verifications
	c.ServeJSON()

}

// AddImage ...
// @Title add Image
// @Description  add Image
// @router /:id/image [put]
func (c *VerificationsController) AddImage() {

	authToken := c.Ctx.Input.Header("Authorization")
	_, err := VerifyToken(authToken, "Watcher")

	if err != nil {
		c.BadRequest(err)
		return
	}

	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetVerificationsByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}
	if err != nil {
		c.BadRequest(err)
		return
	}

	_, filesFh, err := c.GetFile("files")
	if err != nil {
		c.BadRequest(err)
		return
	}

	fileUUID, fileMime, err := files.CreateFile(filesFh, "verifications")

	v.ImageUUID = fileUUID
	v.ImageMime = fileMime

	err = models.UpdateVerificationsByID(v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}

// GetImagesByUUID ...
// @Title Get  By UUID
// @Description Get image By UUID
// @router /image/:uuid [get]
func (c *VerificationsController) GetImagesByUUID() {

	uuid := c.Ctx.Input.Param(":uuid")

	if uuid == "" {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte{})
		return
	}

	imageBytes, mimeType, err := files.GetFile(uuid, "verifications")
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte{})
		return
	}

	c.Ctx.Output.Header("Content-Type", mimeType)
	c.Ctx.Output.SetStatus(200)
	c.Ctx.Output.Body(imageBytes)

}

// Comment ...
// @Title Comment
// @Description comment the Verification
// @router /:id/comment [put]
func (c *VerificationsController) Comment() {

	authToken := c.Ctx.Input.Header("Authorization")
	_, err := VerifyToken(authToken, "Supervisor")

	if err != nil {
		c.BadRequest(err)
		return
	}

	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Verifications{}

	// Validate empty body
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	if v.SupervisorComment == "" {
		err := errors.New("supervisor_comment is missing")
		c.BadRequest(err)
		return
	}

	// Validate context body
	verification, err := models.GetVerificationsByID(id)
	if err != nil {
		c.BadRequest(err)
		return
	}

	verification.SupervisorComment = v.SupervisorComment

	err = models.UpdateVerificationsByID(verification)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = verification
	c.ServeJSON()
}
