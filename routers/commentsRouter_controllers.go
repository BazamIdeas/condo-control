package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["condo-control/controllers:AssistancesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:AssistancesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:AssistancesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:AssistancesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:AssistancesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:AssistancesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:AssistancesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:AssistancesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:AssistancesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:AssistancesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:AssistancesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:AssistancesController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:AssistancesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:AssistancesController"],
		beego.ControllerComments{
			Method: "NewAssistanceExecute",
			Router: `/:token`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:AssistancesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:AssistancesController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CondosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CondosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CondosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CondosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CondosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CondosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CondosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CondosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CondosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CondosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CondosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CondosController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CondosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CondosController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:HolidaysController"] = append(beego.GlobalControllerRouter["condo-control/controllers:HolidaysController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:HolidaysController"] = append(beego.GlobalControllerRouter["condo-control/controllers:HolidaysController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:HolidaysController"] = append(beego.GlobalControllerRouter["condo-control/controllers:HolidaysController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:HolidaysController"] = append(beego.GlobalControllerRouter["condo-control/controllers:HolidaysController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:HolidaysController"] = append(beego.GlobalControllerRouter["condo-control/controllers:HolidaysController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:HolidaysController"] = append(beego.GlobalControllerRouter["condo-control/controllers:HolidaysController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:HolidaysController"] = append(beego.GlobalControllerRouter["condo-control/controllers:HolidaysController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PointsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PointsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PointsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PointsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PointsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PointsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PointsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PointsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PointsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PointsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PointsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PointsController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:PointsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:PointsController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"],
		beego.ControllerComments{
			Method: "NewRouteExecute",
			Router: `/route/:token`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"],
		beego.ControllerComments{
			Method: "NewRoute",
			Router: `/zones/:id/route`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WatchersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WatchersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WatchersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WatchersController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WatchersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WatchersController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WatchersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WatchersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WatchersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WatchersController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WatchersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WatchersController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WatchersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WatchersController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WatchersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WatchersController"],
		beego.ControllerComments{
			Method: "GetSelf",
			Router: `/self`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WatchersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WatchersController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WorkersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WorkersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WorkersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WorkersController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WorkersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WorkersController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WorkersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WorkersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WorkersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WorkersController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WorkersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WorkersController"],
		beego.ControllerComments{
			Method: "GetAssistancesDataByYear",
			Router: `/:id/data/:year/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WorkersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WorkersController"],
		beego.ControllerComments{
			Method: "GetAssistancesDataByMonth",
			Router: `/:id/data/:year/:month`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WorkersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WorkersController"],
		beego.ControllerComments{
			Method: "GetFaceByUUID",
			Router: `/:id/face`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WorkersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WorkersController"],
		beego.ControllerComments{
			Method: "AddImage",
			Router: `/:id/face`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WorkersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WorkersController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WorkersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WorkersController"],
		beego.ControllerComments{
			Method: "GetSelf",
			Router: `/self`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WorkersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WorkersController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ZonesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ZonesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ZonesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ZonesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ZonesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ZonesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ZonesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ZonesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ZonesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ZonesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ZonesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ZonesController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ZonesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ZonesController"],
		beego.ControllerComments{
			Method: "GetSelf",
			Router: `/self`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ZonesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ZonesController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

}
