package tasks

import (
	"condo-control/controllers/services/mails"
	"condo-control/models"
	"fmt"
	"time"

	"github.com/vjeantet/jodaTime"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
)

// the toolbox package
func init() {
	EmptyAssistancesTask := toolbox.NewTask("EmptyAssistancesTask", "0 0 0 * * *", func() (err error) {

		now := time.Now().In(orm.DefaultTimeLoc)

		date := jodaTime.Format("Y-M-d", now)

		condos, err := models.GetCondosWithoutLimit()

		if err != nil {
			return
		}

		for _, condo := range condos {

			supervisors, err := models.GetSupervisorsByCondosID(condo.ID)
			if err != nil {
				continue
			}

			emails := []string{}

			for _, supervisor := range supervisors {
				emails = append(emails, supervisor.Worker.Email)
			}

			emptyAssistances, err := models.GetCondosWorkersEmptyAssistancesByDate(condo.ID, now)

			if err != nil {
				continue
			}

			go func() {
				params := &mails.HTMLParams{
					EmptyAssistancesWorkers: emptyAssistances,
				}

				email := &mails.Email{
					To:         emails,
					Subject:    "Informe de Asistencias Faltantes - " + date,
					HTMLParams: params,
				}

				err := mails.SendMail(email, "005")

				if err != nil {
					fmt.Println(err)
				}
			}()

		}

		return nil
	})

	toolbox.AddTask("EmptyAssistancesTask", EmptyAssistancesTask)
	toolbox.StartTask()
	defer toolbox.StopTask()
}
