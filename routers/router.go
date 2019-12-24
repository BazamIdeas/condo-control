package routers

// @APIVersion 1.0.0
// @Title Condo Control Api
// @Description Condo Control autogenerate documents for your API

import (
	"condo-control/controllers"
	"condo-control/middlewares"

	"github.com/astaxie/beego"
)

func init() {

	middlewares.LoadMiddlewares()

	mainNS := beego.NewNamespace("/v1",
		beego.NSNamespace("/admin",
			beego.NSInclude(
				&controllers.AdminController{},
			),
		),
		beego.NSNamespace("/assistances",
			beego.NSInclude(
				&controllers.AssistancesController{},
			),
		),
		beego.NSNamespace("/condos",
			beego.NSInclude(
				&controllers.CondosController{},
			),
		),
		beego.NSNamespace("/holidays",
			beego.NSInclude(
				&controllers.HolidaysController{},
			),
		),
		beego.NSNamespace("/points",
			beego.NSInclude(
				&controllers.PointsController{},
			),
		),
		beego.NSNamespace("/supervisors",
			beego.NSInclude(
				&controllers.SupervisorsController{},
			),
		),
		beego.NSNamespace("/verifications",
			beego.NSInclude(
				&controllers.VerificationsController{},
			),
		),
		beego.NSNamespace("/watchers",
			beego.NSInclude(
				&controllers.WatchersController{},
			),
		),
		beego.NSNamespace("/workers",
			beego.NSInclude(
				&controllers.WorkersController{},
			),
		),
		beego.NSNamespace("/zones",
			beego.NSInclude(
				&controllers.ZonesController{},
			),
		),
		/// NEW MODELS TASKS MOD
		beego.NSNamespace("/tasks",
			beego.NSInclude(
				&controllers.TasksController{},
			),
		),
		beego.NSNamespace("/goals",
			beego.NSInclude(
				&controllers.GoalsController{},
			),
		),
		beego.NSNamespace("/goals-comments",
			beego.NSInclude(
				&controllers.GoalsCommentsController{},
			),
		),
		beego.NSNamespace("/deliveries",
			beego.NSInclude(
				&controllers.DeliveriesController{},
			),
		),
		beego.NSNamespace("/items",
			beego.NSInclude(
				&controllers.ItemsController{},
			),
		),
		beego.NSNamespace("/notifications",
			beego.NSInclude(
				&controllers.NotificationsController{},
			),
		),
		beego.NSNamespace("/checks",
			beego.NSInclude(
				&controllers.ChecksController{},
			),
		),
		beego.NSNamespace("/occurrences",
			beego.NSInclude(
				&controllers.OccurrencesController{},
			),
		),
		beego.NSNamespace("/objects",
			beego.NSInclude(
				&controllers.ObjectsController{},
			),
		),
		beego.NSNamespace("/residents",
			beego.NSInclude(
				&controllers.ResidentsController{},
			),
		),
		beego.NSNamespace("/questions",
			beego.NSInclude(
				&controllers.QuestionsController{},
			),
		),
		beego.NSNamespace("/votes",
			beego.NSInclude(
				&controllers.VotesController{},
			),
		),
		beego.NSNamespace("/questions-attachments",
			beego.NSInclude(
				&controllers.QuestionsAttachmentsController{},
			),
		),
		beego.NSNamespace("/ws",
			beego.NSInclude(
				&controllers.WebSocketController{},
			),
		),
	)

	beego.AddNamespace(mainNS)
}
