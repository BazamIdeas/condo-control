// @APIVersion 1.0.0
// @Title GASE Api
// @Description GASE autogenerate documents for your API
package routers

import (
	"condo-control/controllers"
	//"condo-control/middlewares"

	"github.com/astaxie/beego"
)

func init() {

	//middlewares.LoadFilters()

	mainNS := beego.NewNamespace("/v1",

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
	)

	beego.AddNamespace(mainNS)
}
