package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["condo-control/controllers:ActivitiesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ActivitiesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ActivitiesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ActivitiesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ActivitiesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ActivitiesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ActivitiesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ActivitiesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ActivitiesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ActivitiesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ActivitiesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ActivitiesController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ActivitiesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ActivitiesController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:BriefsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:BriefsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:BriefsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:BriefsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:BriefsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:BriefsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:BriefsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:BriefsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:BriefsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:BriefsController"],
		beego.ControllerComments{
			Method: "GetOneByCookie",
			Router: `/:service/:cookie`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CartsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CartsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CartsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CartsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CartsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CartsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CartsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CartsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CartsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CartsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CartsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CartsController"],
		beego.ControllerComments{
			Method: "GetOneByCookie",
			Router: `/cookie`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CartsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CartsController"],
		beego.ControllerComments{
			Method: "AddOrDeleteServices",
			Router: `/services`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ClientsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ClientsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ClientsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ClientsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ClientsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ClientsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "ChangePasswordRequest",
			Router: `/change-password/:email`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ClientsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "ChangePassword",
			Router: `/change-password/:token`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ClientsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "GetOneByEmail",
			Router: `/email/:email`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ClientsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CountriesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CountriesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CountriesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CountriesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CountriesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CountriesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CountriesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CountriesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CountriesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CountriesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CountriesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CountriesController"],
		beego.ControllerComments{
			Method: "GetOneByIso",
			Router: `/iso/:iso`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CouponsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CouponsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CouponsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CouponsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CouponsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CouponsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CouponsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CouponsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CouponsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CouponsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CouponsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CouponsController"],
		beego.ControllerComments{
			Method: "GetOneByCode",
			Router: `/value/:code`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CurrenciesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CurrenciesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CurrenciesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CurrenciesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CurrenciesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CurrenciesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CurrenciesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CurrenciesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CurrenciesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CurrenciesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["condo-control/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["condo-control/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["condo-control/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["condo-control/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["condo-control/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["condo-control/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "AddNewsCurrencies",
			Router: `/:id/currencies`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["condo-control/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "DeleteCurrencies",
			Router: `/:id/currencies`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ImagesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ImagesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ImagesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ImagesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ImagesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ImagesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ImagesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ImagesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ImagesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ImagesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ImagesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ImagesController"],
		beego.ControllerComments{
			Method: "ServeImageBySlug",
			Router: `/slug/:slug`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:LocationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:LocationsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:LocationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:LocationsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:LocationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:LocationsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:LocationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:LocationsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:LocationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:LocationsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:OrdersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:OrdersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:OrdersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:OrdersController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:OrdersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:OrdersController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:OrdersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:OrdersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:OrdersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:OrdersController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PaymentsMethodsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PaymentsMethodsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PaymentsMethodsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PaymentsMethodsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PaymentsMethodsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PaymentsMethodsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PaymentsMethodsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PaymentsMethodsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PaymentsMethodsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PaymentsMethodsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PaymentsMethodsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PaymentsMethodsController"],
		beego.ControllerComments{
			Method: "GetOneByIsoAndGateway",
			Router: `/:iso/:gateway`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PortfoliosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PortfoliosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PortfoliosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PortfoliosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PortfoliosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PortfoliosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PortfoliosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PortfoliosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PortfoliosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PortfoliosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PortfoliosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PortfoliosController"],
		beego.ControllerComments{
			Method: "GetByCustomSearch",
			Router: `/custom-search`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PortfoliosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PortfoliosController"],
		beego.ControllerComments{
			Method: "GetPortfoliosByCountry",
			Router: `/iso`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PricesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PricesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PricesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PricesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PricesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PricesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PricesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PricesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PricesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PricesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:SectorsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:SectorsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:SectorsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:SectorsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:SectorsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:SectorsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:SectorsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:SectorsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:SectorsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:SectorsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ServiceFormsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ServiceFormsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ServiceFormsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ServiceFormsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ServiceFormsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ServiceFormsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ServiceFormsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ServiceFormsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ServiceFormsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ServiceFormsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ServiceFormsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ServiceFormsController"],
		beego.ControllerComments{
			Method: "GetOneByService",
			Router: `/service/:slug`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ServicesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ServicesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ServicesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ServicesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ServicesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ServicesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ServicesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ServicesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ServicesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ServicesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ServicesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ServicesController"],
		beego.ControllerComments{
			Method: "GetPricesServiceByCountry",
			Router: `/:id/prices`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:UsersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:UsersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:UsersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:UsersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:UsersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:UsersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:UsersController"],
		beego.ControllerComments{
			Method: "ChangePassword",
			Router: `/change-password`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:UsersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ViewController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ViewController"],
		beego.ControllerComments{
			Method: "Main",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
