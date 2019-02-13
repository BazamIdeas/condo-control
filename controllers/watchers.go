package controllers

import (
	"condo-control/controllers/services/mails"
	"condo-control/models"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"

	"github.com/vjeantet/jodaTime"

	b64 "encoding/base64"

	"github.com/astaxie/beego/validation"
)

// WatchersController operations for Watchers
type WatchersController struct {
	BaseController
}

//URLMapping ...
func (c *WatchersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Login", c.Login)
	c.Mapping("GetSelf", c.GetSelf)
	c.Mapping("GetVerificationsByDate", c.GetVerificationsByDate)
	c.Mapping("GetWatchersVerificationsByMonth", c.GetWatchersVerificationsByMonth)
	c.Mapping("GetByUsername", c.GetByUsername)
	c.Mapping("ChangePublicInfo", c.ChangePublicInfo)
	c.Mapping("ChangePassword", c.ChangePassword)
	c.Mapping("RedirectChangePassword", c.RedirectChangePassword)
	c.Mapping("GenerateChangePasswordToken", c.GenerateChangePasswordToken)

}

// Post ...
// @Title Post
// @Description create Watchers
// @Accept json
// @Param   Authorization     header   string true       "Supervisor's Token"
// @Param   username     body   string true       "watcher's Username"
// @Param   password     body   string true       "watcher's password"
// @Param   phone     body   string false       "watcher's phone"
// @Param   worker     body   object false       "Worker object (first name)"
// @Success 200 {object} models.Watchers
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Zones Don't exists
// @Failure 409 Condo's User limit reached
// @router / [post]
func (c *WatchersController) Post() {

	token := c.Ctx.Input.Header("Authorization")

	decodedToken, err := VerifyToken(token, "Supervisor")

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

	var v models.Watchers

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

	if v.Worker == nil {
		err = errors.New("Worker info is empty")
		c.BadRequest(err)
		return
	}

	if v.Worker.FirstName == "" {

		err = errors.New("Worker's First Name is missing")
		c.BadRequest(err)
		return
	}

	v.Worker.Condo = &models.Condos{ID: condoID}
	v.Worker.Approved = true

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

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Watchers by id
// @router /:id [get]
func (c *WatchersController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetWatchersByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetByUsername ...
// @Title Get By Username
// @Description get By Username by id
// @router /username/:username [get]
func (c *WatchersController) GetByUsername() {

	username := c.Ctx.Input.Param(":username")

	if username == "" {
		err := errors.New("Username is missing")
		c.BadRequest(err)
		return
	}

	v, err := models.GetWatchersByUsername(username)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	err = v.Worker.GetCurrentWorkTimeAssistances()
	if err != nil && err != orm.ErrNoRows {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Watchers
// @router / [get]
func (c *WatchersController) GetAll() {
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

	l, err := models.GetAllWatchers(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Watchers
// @router /:id [put]
func (c *WatchersController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Watchers{ID: id}

	// Validate empty body
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = models.UpdateWatchersByID(&v)

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

// ChangePublicInfo ...
// @Title ChangePublicInfo
// @Description update the Watchers's ChangePublicInfo
// @router /change-public-info [put]
func (c *WatchersController) ChangePublicInfo() {

	token := c.Ctx.Input.Header("Authorization")

	decodedToken, err := VerifyToken(token, "Watcher")

	if err != nil {
		c.BadRequest(err)
		return
	}

	watcherID, err := strconv.Atoi(decodedToken.UserID)
	if err != nil {
		c.BadRequest(err)
		return
	}

	watcher, err := models.GetWatchersByID(watcherID)
	if err != nil {
		c.BadRequestDontExists("Watcher")
		return
	}

	v := models.Watchers{}

	// Validate empty body
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	if v.Password != "" {
		watcher.Password = v.Password
	}

	if v.Phone != "" {
		watcher.Phone = v.Phone
	}

	err = models.UpdateWatchersByID(watcher)

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
// @Description delete the Watchers
// @router /:id [delete]
func (c *WatchersController) Delete() {
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

	err = models.DeleteWatchers(id, trash)

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
func (c *WatchersController) GetAllFromTrash() {

	v, err := models.GetWatchersFromTrash()

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
func (c *WatchersController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Watchers{ID: id}

	err = models.RestoreFromTrash(v.TableName(), v.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}

// Login ...
// @Title Login
// @Description Login
// @router /login [post]
func (c *WatchersController) Login() {

	v := models.Watchers{}

	// Validate empty body

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	// Validate context body

	valid := validation.Validation{}

	valid.Required(v.Username, "username")
	valid.Required(v.Password, "password")

	if valid.HasErrors() {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	id, err := models.LoginWatchers(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	watcherID := strconv.Itoa(id)
	condoID := strconv.Itoa(v.Worker.Condo.ID)

	v.Token, err = c.GenerateToken("Watcher", watcherID, condoID)

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = v

	c.ServeJSON()

}

// GetSelf ...
// @Title Get Self
// @Description Get Self
// @router /self [get]
func (c *WatchersController) GetSelf() {

	token := c.Ctx.Input.Header("Authorization")

	decodedToken, _ := VerifyToken(token, "Supervisor")

	//Disclamer, token already verified
	id, err := strconv.Atoi(decodedToken.CondoID)
	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetWatchersByCondosID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetVerificationsByDate ...
// @Title Get Verifications By Date
// @Description Get Verifications By Date
// @router /:id/verifications/:date [get]
func (c *WatchersController) GetVerificationsByDate() {

	authToken := c.Ctx.Input.Header("Authorization")

	decAuthToken, _ := VerifyToken(authToken, "Supervisor")

	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	watcher, err := models.GetWatchersByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	condoID, _ := strconv.Atoi(decAuthToken.CondoID)

	if watcher.Worker.Condo.ID != condoID {
		err = errors.New("Watcher's Condo and Supervisor's Condo Don't match")
		c.BadRequest(err)
		return
	}

	dateString := c.Ctx.Input.Param(":date")
	date, err := jodaTime.Parse("Y-M-d", dateString)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = watcher.GetVerificationsByDate(date)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = watcher
	c.ServeJSON()

}

// GetWatchersVerificationsByMonth ..
// @Title Get Watchers Verifications By Month
// @Description Get Watchers Verifications By Month
// @Accept json
// @Param   Authorization     header   string true       "Supervisor's Token"
// @Param   year     path   int true       "year's Date"
// @Param   month     path   int true       "month's Date"
// @Success 200 {array} models.Verfications
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Month without Data
// @router /:id/verifications/:year/:month [get]
func (c *WatchersController) GetWatchersVerificationsByMonth() {

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

	_, err = models.GetCondosByID(condoID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	watcherIDstr := c.Ctx.Input.Param(":id")

	watcherID, err := strconv.Atoi(watcherIDstr)
	if err != nil {
		c.BadRequest(err)
		return
	}

	watcher, err := models.GetWatchersByID(watcherID)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	// worker.MonthAssistances =  map[string]map[string]*models.Assistances{}

	verifications, err := models.GetWatchersVerificationsByMonth(watcher.ID, year, month)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	watcher.Verifications = verifications

	c.Data["json"] = watcher
	c.ServeJSON()
}

// GenerateChangePasswordToken ..
// @Title Generate Change Password Token
// @Description Generate Change Password Token
// @Accept json
// @Success 200 {object} models.Watchers
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 email without Data
// @router /:email/change-password/ [post]
func (c *WatchersController) GenerateChangePasswordToken() {

	url := c.GetString("url")
	if url == "" {
		err := errors.New("missing url")
		c.BadRequest(err)
		return
	}

	email := c.Ctx.Input.Param(":email")

	if email == "" {
		err := errors.New("missing email")
		c.BadRequest(err)
		return
	}

	worker, err := models.GetWorkersByEmail(email)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	watcher, err := models.GetWatchersByWorkersID(worker.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	watcherID := strconv.Itoa(watcher.ID)
	condoID := strconv.Itoa(worker.Condo.ID)

	token, err := GenerateGeneralToken(watcherID, condoID, nil, nil, nil)

	if err != nil {
		c.BadRequest(err)
		return
	}

	go func() {

		urlToRoute := c.Ctx.Input.Site() + ":" + strconv.Itoa(c.Ctx.Input.Port()) + c.URLFor("WatchersController.RedirectChangePassword")

		params := &mails.HTMLParams{
			Token:  token,
			URL:    urlToRoute,
			Base64: b64.URLEncoding.EncodeToString([]byte(url)),
		}

		email := &mails.Email{
			To:         []string{email},
			Subject:    "Cambio de Contraseña",
			HTMLParams: params,
		}

		err := mails.SendMail(email, "003")

		if err != nil {
			fmt.Println(err)
		}
	}()

	watcher.Token = token
	c.Data["json"] = watcher
	c.ServeJSON()

}

//RedirectChangePassword ..
// @Title Redirect Change Password
// @Description Redirect Change Password
// @Accept json
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @router /change-password/redirect [get]
func (c *WatchersController) RedirectChangePassword() {

	token := c.GetString("token")
	base64 := c.GetString("base64")

	urlBytes, err := b64.URLEncoding.DecodeString(base64)

	if err != nil {
		return
	}

	c.Ctx.Output.Body([]byte(`<div>
			<a style="color: white; text-decoration: none;" href="` + string(urlBytes) + `/` + token + `">
				<button style="display:block; margin:auto;background-color: #ee7203; border: none; border-radius: 3px; padding: 7px 14px; font-size: 16px; color:  white;">CONTINUAR</button>
			</a>
		</div>`))

	//c.Redirect(string(urlBytes)+"/"+token, 301)

}

//ChangePassword ..
// @Title Change Password
// @Description Change Password
// @Accept json
// @Success 200 {object} models.Watchers
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @router /change-password/:token [put]
func (c *WatchersController) ChangePassword() {

	token := c.Ctx.Input.Param(":token")

	if token == "" {
		err := errors.New("missing token")
		c.BadRequest(err)
		return
	}

	decodedToken, err := VerifyGeneralToken(token)

	if err != nil {
		c.BadRequest(err)
		return
	}

	var v models.Watchers

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	if v.Password == "" {
		err = errors.New("missing Password")
		c.BadRequest(err)
		return
	}

	watcherID, _ := strconv.Atoi(decodedToken.UserID)

	watcher, err := models.GetWatchersByID(watcherID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	watcher.Password = v.Password

	err = models.UpdateWatchersByID(watcher)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = watcher
	c.ServeJSON()

}
