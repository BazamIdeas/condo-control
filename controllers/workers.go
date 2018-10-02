package controllers

import (
	"condo-control/controllers/services/faces"
	"condo-control/models"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/vjeantet/jodaTime"
)

//WorkersController ...
type WorkersController struct {
	BaseController
}

//URLMapping ...
func (c *WorkersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetSelf", c.GetSelf)
}

// Post ...
// @Title Post
// @Description create Workers
// @router / [post]
func (c *WorkersController) Post() {

	err := c.Ctx.Input.ParseFormOrMulitForm(128 << 20)

	if err != nil {
		c.Ctx.Output.SetStatus(413)
		c.ServeJSON()
		return
	}

	var r = c.Ctx.Request

	var (
		firstName = r.FormValue("first_name")
		lastName  = r.FormValue("last_name")
	)

	v := &models.Workers{FirstName: firstName, LastName: lastName}

	if !c.Ctx.Input.IsUpload() {
		err := errors.New("Not image file found on request")
		c.BadRequest(err)
		return
	}

	/* file, fileHeader, err := c.GetFile("faces")

	if err != nil {
		c.BadRequest(err)
		return
	}


	//TODO: VALIDATE IMAGE

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.BadRequest(err)
		return
	} */

	valid := validation.Validation{}

	b, _ := valid.Valid(v)

	if !b {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	_, err = models.AddWorkers(v)

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
// @Description get Workers by id
// @router /:id [get]
func (c *WorkersController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetWorkersByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Workers
// @router / [get]
func (c *WorkersController) GetAll() {
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

	l, err := models.GetAllWorkers(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Workers
// @router /:id [put]
func (c *WorkersController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Workers{ID: id}

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

	//TODO:
	// Validate foreings keys

	/* exists := models.ValidateExists("Sectors", v.Sector.ID)

	if !exists {
		c.BadRequestDontExists("Sector")
		return
	} */

	err = models.UpdateWorkersByID(&v)

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
// @Description delete the Workers
// @router /:id [delete]
func (c *WorkersController) Delete() {
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

	err = models.DeleteWorkers(id, trash)

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
func (c *WorkersController) GetAllFromTrash() {

	v, err := models.GetWorkersFromTrash()

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
func (c *WorkersController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Workers{ID: id}

	err = models.RestoreFromTrash(v.TableName(), v.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}

// GetSelf ...
// @Title Get Self
// @Description Get Self
// @router /self [get]
func (c *WorkersController) GetSelf() {

	token := c.Ctx.Input.Header("Authorization")

	decodedToken, _ := VerifyTokenByAllUserTypes(token)

	//Disclamer, token already verified

	id, err := strconv.Atoi(decodedToken.CondoID)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetWorkersByCondosID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

//GetAssistancesDataByMonth ..
// @Title Get Assistances Data By Month
// @Description Get Assistances Data By Month
// @router /:id/data/:year/:month [get]
func (c *WorkersController) GetAssistancesDataByMonth() {

	yearString := c.Ctx.Input.Param(":year")
	monthSring := c.Ctx.Input.Param(":month")

	date, err := jodaTime.Parse("Y-M", yearString+"-"+monthSring)

	if err != nil {
		c.BadRequest(err)
		return
	}

	year, month, _ := date.Date()

	fmt.Println(year)
	fmt.Println(month)
	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	authToken := c.Ctx.Input.Header("Authorization")
	decAuthToken, err := VerifyToken(authToken, "Supervisor")

	if err != nil {
		c.BadRequest(err)
		return
	}

	worker, err := models.GetWorkersByID(id)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	condoID, _ := strconv.Atoi(decAuthToken.CondoID)

	if worker.Condo.ID != condoID {
		err = errors.New("Worker's Condo and Supervisor's Condo Don't match")
		c.BadRequest(err)
		return
	}

	// worker.MonthAssistances =  map[string]map[string]*models.Assistances{}

	err = worker.GetMonthAssistancesData(year, month)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	v := worker

	c.Data["json"] = v
	c.ServeJSON()
}

//GetAssistancesDataByYear ..
// @Title Get Assistances Data By Year
// @Description Get Assistances Data By Year
// @router /:id/data/:year/ [get]
func (c *WorkersController) GetAssistancesDataByYear() {

	v := ""

	c.Data["json"] = v
	c.ServeJSON()
}

//VerifyWorkerIdentity ...
func VerifyWorkerIdentity(workerID int, newFaceFh *multipart.FileHeader) (worker *models.Workers, ok bool, err error) {

	worker, err = models.GetWorkersByID(workerID)

	if err != nil {
		return
	}

	/////TODO:
	//////TEST MODE /////
	ok = true

	return

	///////END TEST MODE

	oldImageUUID := worker.ImageUUID

	newImageUUID, err := faces.CreateFaceFile(newFaceFh)

	if err != nil {
		return
	}

	defer faces.DeleteFaceFile(newImageUUID)

	newFaceID, err := faces.CreateFaceID(newImageUUID)
	if err != nil {
		return
	}

	oldFaceID, err := faces.CreateFaceID(oldImageUUID)
	if err != nil {
		return
	}

	ok, err = faces.CompareFacesIDs(oldFaceID, newFaceID)

	return
}
