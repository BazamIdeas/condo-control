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
				&controllers.GoalsController{},
			),
		),
	)

	beego.AddNamespace(mainNS)
}
