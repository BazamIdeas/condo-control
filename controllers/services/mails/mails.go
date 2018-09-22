package mails

import (
	"GASE/models"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/astaxie/beego"
	gomail "gopkg.in/gomail.v2"
)

type HTMLParams struct {
	ClientName string
	AdminName  string
	Client     *models.Clients
	Service    *models.Services
	Services   []*models.Services
	Order      *models.Orders
	Country    *models.Countries
	Location   *models.Locations
	Sectors    *models.Sectors
	Activities *models.Activities
	Gateway    *models.Gateways
	Portfolio  *models.Portfolios
	Brief      *models.Briefs
	Token      string
}

//Email ...
type Email struct {
	To          []string
	Bcc         []string
	Cc          []string
	Subject     string
	html        string
	Headers     [][]string
	Attachments []string
	HTMLParams  *HTMLParams
}

var (
	from    = beego.AppConfig.String("email::from")
	name    = beego.AppConfig.String("email::name")
	pass    = beego.AppConfig.String("email::pass")
	host    = beego.AppConfig.String("email::host")
	port, _ = beego.AppConfig.Int("email::port")
)

var (
	rootDir, _     = filepath.Abs(beego.AppConfig.String("assets::jumps"))
	mailFolderPath = beego.AppConfig.String("assets::mailTemplatePath")
	mailFolderDir  = rootDir + "/" + mailFolderPath
)

func init() {
	checkOrCreateImagesFolder(mailFolderDir)
}

/* func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
} */

func getTemplate(code string, HTMLParams *HTMLParams) (tplHTML *string, err error) {

	filesInfo, err := ioutil.ReadDir(mailFolderDir)

	if err != nil {
		return
	}

	var tplName string

	for _, filesInfo := range filesInfo {

		fileInfoName := filesInfo.Name()

		fileInfoNameSplited := strings.Split(fileInfoName, ".")

		if len(fileInfoNameSplited) == 3 && fileInfoNameSplited[1] == code {

			tplName = fileInfoName
			break
		}

	}

	if tplName == "" {
		err = errors.New("Template is missing")
		return
	}

	tpl, err := template.ParseFiles(mailFolderDir + "/" + tplName)

	if err != nil {
		return
	}

	var tplBuffer bytes.Buffer

	err = tpl.Execute(&tplBuffer, HTMLParams)

	if err != nil {
		fmt.Println(err)
		return
	}

	tplString := tplBuffer.String()

	return &tplString, nil
}

func (e *Email) loadTemplate(code string) (err error) {

	tplHTML, err := getTemplate(code, e.HTMLParams)

	if err != nil {
		return
	}

	e.html = *tplHTML

	return
}

//SendMails ...
func SendMails(Emails []*Email, code string) {

	ch := make(chan *gomail.Message)

	for _, Email := range Emails {
		go SendMail(Email, code)
	}

	close(ch)
}

//SendMail ...
func SendMail(Email *Email, code string) (err error) {
	m := gomail.NewMessage()

	err = Email.loadTemplate(code)

	if err != nil {
		return
	}

	for _, header := range Email.Headers {
		m.SetHeader(header[0], header[1])
	}

	for _, Attachment := range Email.Attachments {
		m.Attach(Attachment)
	}

	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(from, name)},
		"To":      Email.To,
		"Cc":      Email.Cc,
		"Bcc":     Email.Bcc,
		"Subject": {Email.Subject},
		"X-Date":  {m.FormatDate(time.Now())},
	})

	m.SetBody("text/html", Email.html)

	d := gomail.NewDialer(host, port, from, pass)

	err = d.DialAndSend(m)

	if err != nil {
		return
	}

	return

}

func checkOrCreateImagesFolder(imageFolderDir string) (err error) {

	if _, err := os.Stat(imageFolderDir); os.IsNotExist(err) {
		os.MkdirAll(imageFolderDir, 644)
	}
	return
}
