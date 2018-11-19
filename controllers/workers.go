package controllers

import (
	"bytes"
	"condo-control/controllers/services/faces"
	"condo-control/models"
	"encoding/csv"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
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
	c.Mapping("AddImage", c.AddImage)
	c.Mapping("GetAssistancesDataByMonth", c.GetAssistancesDataByMonth)
	c.Mapping("GetAssistancesDataByYear", c.GetAssistancesDataByYear)
	c.Mapping("DownloadAssistancesDataByMonth", c.DownloadAssistancesDataByMonth)
	c.Mapping("Approve", c.Approve)
}

// Post ...
// @Title Post
// @Description create Workers
// @Accept json
// @Param   Authorization     header   string true       "Watcher's Token"
// @Param   first_name      body   string true       "New Worker's first name"
// @Success 201 {object} models.Workers
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 409 Condo's user limit reached
// @router / [post]
func (c *WorkersController) Post() {

	token := c.Ctx.Input.Header("Authorization")

	decodedToken, err := VerifyToken(token, "Watcher")

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

	condosWorkersCount := len(condos.Workers) + 1

	if condosWorkersCount > condos.UserLimit {
		c.Ctx.Output.SetStatus(409)
		c.Data["json"] = MessageResponse{
			Code:    409,
			Message: "Condo's user limit reached",
		}
		c.ServeJSON()
		return
	}

	var v models.Workers

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

	v.Condo = &models.Condos{ID: condoID}
	v.Approved = false

	_, err = models.AddWorkers(&v)
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
// @Accept json
// @Param   id     path   string true       "Worker's id"
// @Success 200 {object} models.Workers
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Worker Don't exist
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

//TODO: REMOVE OR ADAPT

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

//TODO: REMOVE OR ADAPT

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

//TODO: REMOVE OR ADAPT

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

//TODO: REMOVE OR ADAPT

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

//TODO: REMOVE OR ADAPT

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
// @Description Get Workers from Watcher's Condo or Supervisor's Condo
// @Accept json
// @Param   Authorization     header   string true       "Watcher's Token or Supervisor's Token"
// @Success 200 {array} models.Workers
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Condos without Workers
// @router /self [get]
func (c *WorkersController) GetSelf() {

	token := c.Ctx.Input.Header("Authorization")

	decodedToken, _, _ := VerifyTokenByAllUserTypes(token)

	//Disclamer, token already verified

	condoID, err := strconv.Atoi(decodedToken.CondoID)

	if err != nil {
		c.BadRequest(err)
		return
	}

	workers, err := models.GetWorkersByCondosID(condoID)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	supervisors, err := models.GetSupervisorsByCondosID(condoID)
	if err != nil && err != orm.ErrNoRows {
		c.ServeErrorJSON(err)
		return
	}

	workersNoSupervisors := []*models.Workers{}

	if supervisors != nil && len(supervisors) > 0 {
		for _, worker := range workers {
			noSupervisor := true
			for _, supervisor := range supervisors {
				if supervisor.Worker.ID == worker.ID {
					noSupervisor = false
					break
				}

			}

			if noSupervisor {
				workersNoSupervisors = append(workersNoSupervisors, worker)
			}
		}
	}

	c.Data["json"] = workersNoSupervisors
	c.ServeJSON()
}

//GetAssistancesDataByMonth ..
// @Title Get Assistances Data By Month
// @Description Get Assistances Data By Month
// @Accept json
// @Param   Authorization     header   string true       "Supervisor's Token"
// @Param   year     path   int true       "year's Date"
// @Param   month     path   int true       "month's Date"
// @Success 200 {object} models.Workers
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Month without Data
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

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	authToken := c.Ctx.Input.Header("Authorization")
	decAuthToken, _ := VerifyToken(authToken, "Supervisor")

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
// @Accept json
// @Param   Authorization     header   string true       "Supervisor's Token"
// @Param   year     path   int true       "year's Date"
// @Success 200 {object} models.Workers
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Year without Data
// @router /:id/data/:year/ [get]
func (c *WorkersController) GetAssistancesDataByYear() {

	yearString := c.Ctx.Input.Param(":year")

	date, err := jodaTime.Parse("Y", yearString)

	if err != nil {
		c.BadRequest(err)
		return
	}

	year, _, _ := date.Date()

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

	err = worker.GetYearAssistancesData(year)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	v := worker

	c.Data["json"] = v
	c.ServeJSON()
}

//VerifyWorkerIdentity ...
func VerifyWorkerIdentity(workerID int, newFaceFh *multipart.FileHeader) (worker *models.Workers, ok bool, err error) {

	worker, err = models.GetWorkersByID(workerID)

	if err != nil {
		return
	}

	oldImageUUID := worker.ImageUUID
	if oldImageUUID == "" {
		err = errors.New("Worker lack registered Face")
		return
	}

	facesDebug, _ := beego.AppConfig.Bool("faces::debug")

	if facesDebug {
		ok = true
		return
	}

	newImageUUID, _, err := faces.CreateFaceFile(newFaceFh)

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

// AddImage ...
// @Title Add Image
// @Description Add Image
// @Accept plain
// @Param   Authorization     header   string true       "Supervisor's Token"
// @Param   id     path   int true       "worker's id"
// @Param   faces     formData   string true       "worker's id"
// @Success 200 {object} models.Workers
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Workers not Found
// @Failure 413 File size too High
// @router /:id/face [post]
func (c *WorkersController) AddImage() {

	err := c.Ctx.Input.ParseFormOrMulitForm(128 << 20)

	if err != nil {
		c.Ctx.Output.SetStatus(413)
		c.ServeJSON()
		return
	}

	if !c.Ctx.Input.IsUpload() {
		err := errors.New("Not image file found on request")
		c.BadRequest(err)
		return
	}

	token := c.Ctx.Input.Header("Authorization")

	decodedToken, _ := VerifyToken(token, "Watcher")

	//Disclamer, token already verified
	condoID, err := strconv.Atoi(decodedToken.CondoID)
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

	worker, err := models.GetWorkersByID(id)

	if err != nil {
		c.BadRequestDontExists("Condos")
		return
	}

	if worker.Condo.ID != condoID {
		err = errors.New("Worker's Condo and Watcher's Condo Don't match")
		c.BadRequest(err)
		return
	}

	_, faceFh, err := c.GetFile("faces")

	if err != nil {
		c.BadRequest(err)
		return
	}

	newImageUUID, mimeType, err := faces.CreateFaceFile(faceFh)
	if err != nil {
		return
	}

	worker.ImageUUID = newImageUUID
	worker.ImageMime = mimeType

	err = models.UpdateWorkersByID(worker)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = worker
	c.ServeJSON()
}

// GetFaceByUUID ...
// @Title Get Face By UUID
// @Description Get Face By UUID
// @Accept plain
// @Param   uuid     path   string true       "worker's face uuid"
// @Success 200 {string} Face Image
// @Failure 400 Bad Request
// @Failure 404 Face not Found
// @router /face/:uuid [get]
func (c *WorkersController) GetFaceByUUID() {

	uuid := c.Ctx.Input.Param(":uuid")

	if uuid == "" {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte{})
		return
	}

	imageBytes, mimeType, err := faces.GetFaceFile(uuid)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte{})
		return
	}

	c.Ctx.Output.Header("Content-Type", mimeType)
	c.Ctx.Output.SetStatus(200)
	c.Ctx.Output.Body(imageBytes)

}

//Approve ..
// @Title Approve Worker
// @Description Approve Worker
// @Accept json
// @Param   id     path   string true       "Pending approve Worker's id"
// @Success 200 {object} models.Workers
// @Failure 400 Bad Request
// @Failure 404 Worker not Found
// @Failure 409 Worker already approved
// @router /:id/approve [patch]
func (c *WorkersController) Approve() {

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

	if worker.Approved {
		c.Ctx.Output.SetStatus(409)
		c.Data["json"] = MessageResponse{
			Code:    409,
			Message: "Worker is already approved",
		}
		c.ServeJSON()
		return
	}

	worker.Approved = true
	err = models.UpdateWorkersByID(worker)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	v := worker

	c.Data["json"] = v
	c.ServeJSON()
}

// DownloadAssistancesDataByYear ...
// @Title Download Assistances Data By Year
// @Description Download Assistances Data By Year
// @Accept plain
// @Param   id     path   int true       "Worker's id"
// @Param   year     path   int true       "year's Date"
// @Success 200 {string} Csv data
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Worker or Assistances not Found
// @router /:id/data/:year/download [get]
func (c *WorkersController) DownloadAssistancesDataByYear() {

	yearString := c.Ctx.Input.Param(":year")
	date, err := jodaTime.Parse("Y", yearString)
	if err != nil {
		c.BadRequest(err)
		return
	}

	year, _, _ := date.Date()

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
		c.BadRequestDontExists("Workers")
		return
	}

	condoID, _ := strconv.Atoi(decAuthToken.CondoID)
	if worker.Condo.ID != condoID {
		err = errors.New("Worker's Condo and Supervisor's Condo Don't match")
		c.BadRequest(err)
		return
	}

	err = worker.GetYearAssistancesData(year)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	records := [][]string{{"Mes", "Asistencia en Días", "Horas Adicionales", "Valor Hora", "Adicional Remuneración", "Festivos", "Total Días"}}

	yearData := *worker.YearData

	for monthKey, monthData := range yearData {

		extraWorkedHours := strconv.FormatFloat(float64(monthData.ExtraWorkedHours), 'f', 2, 32)
		extraValue := strconv.FormatFloat(float64(monthData.ExtraValue), 'f', 2, 32)

		extraHourIncreasedFloat := worker.Condo.HourValue + (worker.Condo.HourValue * (worker.Condo.ExtraHourIncrease / 100))
		extraHourIncreased := strconv.FormatFloat(float64(extraHourIncreasedFloat), 'f', 2, 32)

		workedDays := len(*monthData.Days)

		record := []string{strconv.Itoa(monthKey), strconv.Itoa(workedDays - monthData.Holidays), extraWorkedHours, extraHourIncreased, extraValue, strconv.Itoa(monthData.Holidays), strconv.Itoa(workedDays)}
		records = append(records, record)
	}

	buffer := &bytes.Buffer{}
	w := csv.NewWriter(buffer)
	w.WriteAll(records)

	if err := w.Error(); err != nil {
		err = errors.New("Error Writing Csv")
		c.BadRequest(err)
		return
	}

	dateString := jodaTime.Format("Y-M", date)

	http.ServeContent(c.Ctx.ResponseWriter, c.Ctx.Request, dateString+" data.csv", time.Now(), bytes.NewReader(buffer.Bytes()))

}

// DownloadAssistancesDataByMonth ...
// @Titile Download Assistances Data By Month
// @Description Download Assistances Data By Month
// @Accept plain
// @Param   id     path   int true       "Worker's id"
// @Param   year     path   int true       "year's Date"
// @Param   month     path   int true       "month's Date"
// @Success 200 {object} models.Workers
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Worker or Assistances not Found
// @router /:id/data/:year/:month/download [get]
func (c *WorkersController) DownloadAssistancesDataByMonth() {

	yearString := c.Ctx.Input.Param(":year")
	monthSring := c.Ctx.Input.Param(":month")

	date, err := jodaTime.Parse("Y-M", yearString+"-"+monthSring)

	if err != nil {
		c.BadRequest(err)
		return
	}

	year, month, _ := date.Date()

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

	err = worker.GetMonthAssistancesData(year, month)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	records := [][]string{{"Día", "Entrada", "Inicio Colación", "Termino Colación", "Salida", "Horas Trabajadas", "Diferencial", "Festivo"}}

	Days := *worker.MonthData.Days

	for dayKey, day := range Days {

		isHoliday := ""

		if day.IsHoliday {
			isHoliday = "1"
		} else {
			isHoliday = "0"
		}

		totalWorkedHours := strconv.FormatFloat(float64(day.TotalWorkedHours), 'f', 2, 32)
		extraWorkedHours := strconv.FormatFloat(float64(day.ExtraWorkedHours), 'f', 2, 32)

		record := []string{dayKey, day.Entry.Date, "", "", day.Exit.Date, totalWorkedHours, extraWorkedHours, isHoliday}

		if day.Break != nil && day.FinishBreak != nil {
			record[2] = day.Break.Date
			record[3] = day.FinishBreak.Date
		}

		records = append(records, record)
	}

	buffer := &bytes.Buffer{}
	w := csv.NewWriter(buffer)
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		err = errors.New("Error Writing Csv")
		c.BadRequest(err)
		return
	}

	dateString := jodaTime.Format("Y-M", date)

	http.ServeContent(c.Ctx.ResponseWriter, c.Ctx.Request, dateString+"data.csv", time.Now(), bytes.NewReader(buffer.Bytes()))

}
