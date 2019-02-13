package controllers

import (
	"condo-control/controllers/services/faces"
	"condo-control/controllers/services/mails"
	"condo-control/models"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"

	b64 "encoding/base64"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// ResidentsController operations for Residents
type ResidentsController struct {
	BaseController
}

//URLMapping ...
func (c *ResidentsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Login", c.Login)
	c.Mapping("GetSelf", c.GetSelf)
	c.Mapping("ChangePublicInfo", c.ChangePublicInfo)
	c.Mapping("ChangePassword", c.ChangePassword)
	c.Mapping("RedirectChangePassword", c.RedirectChangePassword)
	c.Mapping("GenerateChangePasswordToken", c.GenerateChangePasswordToken)
	c.Mapping("AddImage", c.AddImage)

}

// Post ...
// @Title Post
// @Description create Residents
// @Accept json
// @Param   Authorization     header   string true       "Supervisor's Token"
// @Param   username     body   string true       "resident's Username"
// @Param   password     body   string true       "resident's password"
// @Param   phone     body   string false       "resident's phone"
// @Param   Resident     body   object false       "Resident object (first name)"
// @Success 200 {object} models.Residents
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Zones Don't exists
// @Failure 409 Condo's User limit reached
// @router / [post]
func (c *ResidentsController) Post() {

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

	var v models.Residents

	// Validate empty body
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v.Approved = true

	// Validate context body
	valid := validation.Validation{}

	b, _ := valid.Valid(&v)
	if !b {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	v.Condo = condos

	_, err = models.AddResidents(&v)
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
// @Description get Residents by id
// @router /:id [get]
func (c *ResidentsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetResidentsByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Residents
// @router / [get]
func (c *ResidentsController) GetAll() {
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

	l, err := models.GetAllResidents(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Residents
// @router /:id [put]
func (c *ResidentsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Residents{ID: id}

	// Validate empty body
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = models.UpdateResidentsByID(&v)

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
// @Description update the Residents's ChangePublicInfo
// @router /change-public-info [put]
func (c *ResidentsController) ChangePublicInfo() {

	token := c.Ctx.Input.Header("Authorization")

	decodedToken, err := VerifyToken(token, "Resident")

	if err != nil {
		c.BadRequest(err)
		return
	}

	residentID, err := strconv.Atoi(decodedToken.UserID)
	if err != nil {
		c.BadRequest(err)
		return
	}

	resident, err := models.GetResidentsByID(residentID)
	if err != nil {
		c.BadRequestDontExists("Resident")
		return
	}

	v := models.Residents{}

	// Validate empty body
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	if v.Password != "" {
		resident.Password = v.Password
	}

	if v.Phone != "" {
		resident.Phone = v.Phone
	}

	err = models.UpdateResidentsByID(resident)

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
// @Description delete the Residents
// @router /:id [delete]
func (c *ResidentsController) Delete() {
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

	err = models.DeleteResidents(id, trash)

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
func (c *ResidentsController) GetAllFromTrash() {

	v, err := models.GetResidentsFromTrash()

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
func (c *ResidentsController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Residents{ID: id}

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
func (c *ResidentsController) Login() {

	v := models.Residents{}

	// Validate empty body

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	// Validate context body

	valid := validation.Validation{}
	valid.Required(v.Email, "email")
	valid.Required(v.Password, "password")

	if valid.HasErrors() {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	id, err := models.LoginResidents(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	if !v.Approved {
		err = errors.New("Resident no approved")
		c.BadRequest(err)
		return
	}

	residentID := strconv.Itoa(id)
	condoID := strconv.Itoa(v.Condo.ID)

	v.Token, err = c.GenerateToken("Resident", residentID, condoID)

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
func (c *ResidentsController) GetSelf() {

	token := c.Ctx.Input.Header("Authorization")

	decodedToken, _ := VerifyToken(token, "Supervisor")

	//Disclamer, token already verified
	id, err := strconv.Atoi(decodedToken.CondoID)
	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetResidentsByCondosID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GenerateChangePasswordToken ..
// @Title Generate Change Password Token
// @Description Generate Change Password Token
// @Accept json
// @Success 200 {object} models.Residents
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 email without Data
// @router /:email/change-password/ [post]
func (c *ResidentsController) GenerateChangePasswordToken() {

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

	resident, err := models.GetResidentsByEmail(email)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	residentID := strconv.Itoa(resident.ID)
	condoID := strconv.Itoa(resident.Condo.ID)

	token, err := GenerateGeneralToken(residentID, condoID, nil, nil, nil)

	if err != nil {
		c.BadRequest(err)
		return
	}

	go func() {

		urlToRoute := c.Ctx.Input.Site() + ":" + strconv.Itoa(c.Ctx.Input.Port()) + c.URLFor("ResidentsController.RedirectChangePassword")

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

	resident.Token = token
	c.Data["json"] = resident
	c.ServeJSON()

}

//RedirectChangePassword ..
// @Title Redirect Change Password
// @Description Redirect Change Password
// @Accept json
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @router /change-password/redirect [get]
func (c *ResidentsController) RedirectChangePassword() {

	token := c.GetString("token")
	base64 := c.GetString("base64")

	urlBytes, err := b64.URLEncoding.DecodeString(base64)

	if err != nil {
		return
	}

	c.Ctx.Output.Body([]byte(`<div>
			<a style="color: white; text-decoration: none;" href="` + string(urlBytes) + `/` + token + `">
				<button style="backround-color: #ee7203; border: none; border-radius: 3px; padding: 7px 14px; font-size: 16px; color:  white;">CAMBIAR CONTRASEÑA</button>
			</a>
		</div>`))

	//c.Redirect(string(urlBytes)+"/"+token, 301)

}

//ChangePassword ..
// @Title Change Password
// @Description Change Password
// @Accept json
// @Success 200 {object} models.Residents
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @router /change-password/:token [put]
func (c *ResidentsController) ChangePassword() {

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

	var v models.Residents

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

	residentID, _ := strconv.Atoi(decodedToken.UserID)

	resident, err := models.GetResidentsByID(residentID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	resident.Password = v.Password

	err = models.UpdateResidentsByID(resident)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = resident
	c.ServeJSON()

}

// AddImage ...
// @Title Add Image
// @Description Add Image
// @Accept plain
// @Param   Authorization     header   string true       "Supervisor's Token"
// @Param   id     path   int true       "Resident's id"
// @Param   faces     formData   string true       "Resident's id"
// @Success 200 {object} models.Residents
// @Failure 400 Bad Request
// @Failure 403 Invalid Token
// @Failure 404 Residents not Found
// @Failure 413 File size too High
// @router /:id/face [post]
func (c *ResidentsController) AddImage() {

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

	decodedToken, _ := VerifyToken(token, "Resident")

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

	Resident, err := models.GetResidentsByID(id)

	if err != nil {
		c.BadRequestDontExists("Condos")
		return
	}

	if Resident.Condo.ID != condoID {
		err = errors.New("Resident's Condo and Resident's Condo Don't match")
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

	Resident.ImageUUID = newImageUUID
	Resident.ImageMime = mimeType

	err = models.UpdateResidentsByID(Resident)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = Resident
	c.ServeJSON()
}

// GetFaceByUUID ...
// @Title Get Face By UUID
// @Description Get Face By UUID
// @Accept plain
// @Param   uuid     path   string true       "Resident's face uuid"
// @Success 200 {string} Face Image
// @Failure 400 Bad Request
// @Failure 404 Face not Found
// @router /face/:uuid [get]
func (c *ResidentsController) GetFaceByUUID() {

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

// GetByEmail ...
// @Title Get By Email
// @Description get By Email by id
// @router /email/:email [get]
func (c *ResidentsController) GetByEmail() {

	email := c.Ctx.Input.Param(":email")

	if email == "" {
		err := errors.New("email is missing")
		c.BadRequest(err)
		return
	}

	v, err := models.GetResidentsByEmail(email)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

//VerifyResidentIdentity ...
func VerifyResidentIdentity(residentID int, newFaceFh *multipart.FileHeader) (resident *models.Residents, ok bool, err error) {

	resident, err = models.GetResidentsByID(residentID)

	if err != nil {
		return
	}

	oldImageUUID := resident.ImageUUID
	if oldImageUUID == "" {
		err = errors.New("Resident lack registered Face")
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
