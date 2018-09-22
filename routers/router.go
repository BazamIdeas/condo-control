// @APIVersion 1.0.0
// @Title GASE Api
// @Description GASE autogenerate documents for your API
package routers

import (
	"condo-control/controllers"
	"condo-control/middlewares"

	"github.com/astaxie/beego"
)

func init() {

	middlewares.LoadFilters()

	views := beego.NewNamespace("/admin",
		beego.NSInclude(
			&controllers.ViewController{},
		),
	)

	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/condos",
			beego.NSInclude(
				&controllers.CondosController{},
			),
		),
	)

	beego.AddNamespace(ns)
	beego.AddNamespace(views)
}
