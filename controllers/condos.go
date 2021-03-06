package controllers

import (
	"condo-control/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/vjeantet/jodaTime"

	"github.com/astaxie/beego/validation"
)

// CondosController operations for Condos
type CondosController struct {
	BaseController
}

//URLMapping ...
func (c *CondosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetSelf", c.GetSelf)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetSupervisorsByCondosID", c.GetSupervisorsByCondosID)
	c.Mapping("GetByRUT", c.GetByRUT)
	c.Mapping("GetSelfVerificationsByMonth", c.GetSelfVerificationsByMonth)
	c.Mapping("GetSelfChecksByDate", c.GetSelfChecksByDate)
	c.Mapping("GetSelfChecksByMonth", c.GetSelfChecksByMonth)
}

// Post ...
// @Title Post
// @Description create Condos
// @Accept json
// @Param Authorization header string true "Supervisor's Token"
// @Param user_limit body int true "max users per condo"
// @Param zone_limit body int true "max zones per condo"
// @Param hour_value body int true "worker's hour value"
// @Param extra_hour_increase body int true "Percentage Increase to Hour Value"
// @Param working_hours body int true "Hours to work"
// @Param assistances_mod body bool true "Assistances Capabilities"
// @Param routes_mod body bool true "Routes Capabilities"
// @Success 201 {object} models.Condos
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @router / [post]
func (c *CondosController) Post() {
	var v models.Condos

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

	_, err = models.AddCondos(&v)

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
// @Description get Condos by id
// @router /:id [get]
func (c *CondosController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetCondosByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetSelfEmptyVerificationsWorker ...
// @Title Get Self Empty Verifications Worker
// @Description Get Self Empty Verifications Worker
// @router /:id/empty-verifications/:date/workers [get]
func (c *CondosController) GetSelfEmptyVerificationsWorker() {

	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.BadRequest(err)
		return
	}

	condo, err := models.GetCondosByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	dateString := c.Ctx.Input.Param(":date")
	date, err := jodaTime.Parse("Y-M-d", dateString)
	if err != nil {
		c.BadRequest(err)
		return
	}

	workersEmptyAssistances, err := models.GetCondosWorkersEmptyAssistancesByDate(condo.ID, date)
	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Data["json"] = workersEmptyAssistances
	c.ServeJSON()
}

// GetByRUT ...
// @Title Get By RUT
// @Description Ge tBy RUT
// @router /rut/:rut [get]
func (c *CondosController) GetByRUT() {

	rut := c.Ctx.Input.Param(":rut")
	if rut == "" {
		err := errors.New("rut is empty")
		c.BadRequest(err)
		return
	}

	v, err := models.GetCondosByRUT(rut)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetSupervisorsByCondosID ...
// @Title Get Supervisors By CondosID
// @Description Get Supervisors By CondosID
// @router /:id/supervisors [get]
func (c *CondosController) GetSupervisorsByCondosID() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	_, err = models.GetCondosByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	supervisors, err := models.GetSupervisorsByCondosID(id)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = supervisors
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Condos
// @router / [get]
func (c *CondosController) GetAll() {
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

	l, err := models.GetAllCondos(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	for index := range l {

		condo := l[index].(models.Condos)

		supervisors, err := models.GetSupervisorsByCondosID(condo.ID)

		if err != nil {
			if err == orm.ErrNoRows {
				continue
			}
			c.ServeErrorJSON(err)
			return
		}

		condo.Supervisors = supervisors

		l[index] = condo

	}

	c.Data["json"] = l
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Condos
// @router /:id [put]
func (c *CondosController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Condos{ID: id}

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

	err = models.UpdateCondosByID(&v)

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
// @Description delete the Condos
// @router /:id [delete]
func (c *CondosController) Delete() {
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

	err = models.DeleteCondos(id, trash)

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
func (c *CondosController) GetAllFromTrash() {

	v, err := models.GetCondosFromTrash()

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
func (c *CondosController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Condos{ID: id}

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
// @Accept json
// @Param   Authorization     header   string true       "Watcher's Token or Supervisor's Token"
// @Success 200 {object} models.Condos
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Condos Don't Exists
// @router /self [get]
func (c *CondosController) GetSelf() {

	token := c.Ctx.Input.Header("Authorization")

	decodedToken, _, err := VerifyTokenByAllUserTypes(token)
	if err != nil {
		c.BadRequest(err)
		return
	}

	condoID, err := strconv.Atoi(decodedToken.CondoID)
	if err != nil {
		c.BadRequest(err)
		return
	}

	condos, err := models.GetCondosByID(condoID)
	if err != nil {
		c.BadRequestDontExists("Condos")
		return
	}

	c.Data["json"] = condos
	c.ServeJSON()
}

// AddWatchersToCondosByRUT ...
// @Title Add Watchers To Condos By RUT
// @Description Add Watchers To Condos By RUT
// @Accept json
// @Success 200 {object} models.Watchers
// @Failure 400 Bad Request
// @Failure 404 Condos Don't Exists
// @router /rut/:rut/watchers [post]
func (c *CondosController) AddWatchersToCondosByRUT() {

	v := models.Watchers{}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	if v.Worker == nil {
		err = errors.New("Worker data is missing")
		c.BadRequest(err)
		return
	}

	valid := validation.Validation{}
	b, _ := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	RUTStr := c.Ctx.Input.Param(":rut")

	if RUTStr == "" {
		err = errors.New("Missing RUT")
		c.BadRequest(err)
		return
	}

	condo, err := models.GetCondosByRUT(RUTStr)
	if err != nil {
		c.BadRequestDontExists("Condos")
		return
	}

	v.Worker.Condo = condo

	_, err = models.AddWorkers(v.Worker)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	_, err = models.AddWatchers(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// AddResidentsToCondosByRUT ...
// @Title Add Residents To Condos By RUT
// @Description Add Residents To Condos By RUT
// @Accept json
// @Success 200 {object} models.Residents
// @Failure 400 Bad Request
// @Failure 404 Condos Don't Exists
// @router /rut/:rut/residents [post]
func (c *CondosController) AddResidentsToCondosByRUT() {

	v := models.Residents{}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v.Approved = false

	valid := validation.Validation{}
	b, _ := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	RUTStr := c.Ctx.Input.Param(":rut")

	if RUTStr == "" {
		err = errors.New("Missing RUT")
		c.BadRequest(err)
		return
	}

	condo, err := models.GetCondosByRUT(RUTStr)
	if err != nil {
		c.BadRequestDontExists("Condos")
		return
	}

	v.Condo = condo

	_, err = models.AddResidents(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

//GetSelfVerificationsByMonth ..
// @Title Get Self Verifications By Month
// @Description Get Self Verifications By Month
// @Accept json
// @Param   Authorization     header   string true       "Supervisor's Token"
// @Param   year     path   int true       "year's Date"
// @Param   month     path   int true       "month's Date"
// @Success 200 {array} models.Verfications
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Month without Data
// @router /verifications/:year/:month [get]
func (c *CondosController) GetSelfVerificationsByMonth() {

	yearString := c.Ctx.Input.Param(":year")
	monthSring := c.Ctx.Input.Param(":month")

	date, err := jodaTime.Parse("Y-M", yearString+"-"+monthSring)

	if err != nil {
		c.BadRequest(err)
		return
	}

	year, month, _ := date.Date()

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

	// worker.MonthAssistances =  map[string]map[string]*models.Assistances{}

	verifications, err := models.GetCondosVerificationsByMonth(condo.ID, year, month)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	var v []*models.Verifications

	for _, verification := range verifications {

		/* 		if c.Ctx.Input.Query("supervisor-comment") == "true" && verification.SupervisorComment == "" {
		   			continue
		   		}

		   		if c.Ctx.Input.Query("watcher-comment") == "true" && verification.WatcherComment == "" {
		   			continue
		   		} */

		v = append(v, verification)
	}

	if len(v) == 0 {
		err = orm.ErrNoRows
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

//GetSelfChecksByMonth ..
// @Title Get Self Checks By Month
// @Description Get Self Checks By Month
// @Accept json
// @Param   Authorization     header   string true       "Supervisor's Token"
// @Param   year     path   int true       "year's Date"
// @Param   month     path   int true       "month's Date"
// @Success 200 {array} models.Checks
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Month without Data
// @router /checks/:year/:month [get]
func (c *CondosController) GetSelfChecksByMonth() {

	yearString := c.Ctx.Input.Param(":year")
	monthSring := c.Ctx.Input.Param(":month")

	date, err := jodaTime.Parse("Y-M", yearString+"-"+monthSring)

	if err != nil {
		c.BadRequest(err)
		return
	}

	year, month, _ := date.Date()

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

	checks, err := models.GetCondosChecksByMonth(condo.ID, year, month)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = checks
	c.ServeJSON()
}

//GetSelfChecksByDate ..
// @Title Get Self Checks By Date
// @Description Get Self Checks By Date
// @Accept json
// @Param   Authorization     header   string true       "Supervisor's Token"
// @Success 200 {array} models.Checks
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Month without Data
// @router /checks/:date [get]
func (c *CondosController) GetSelfChecksByDate() {

	dateString := c.Ctx.Input.Param(":date")
	date, err := jodaTime.Parse("Y-M-d", dateString)

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

	condoID, _ := strconv.Atoi(decAuthToken.CondoID)

	condo, err := models.GetCondosByID(condoID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	checks, err := models.GetCondosChecksByDate(condo.ID, date)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = checks
	c.ServeJSON()
}
