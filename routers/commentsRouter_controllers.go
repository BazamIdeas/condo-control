package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["condo-control/controllers:AdminController"] = append(beego.GlobalControllerRouter["condo-control/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

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
			Method: "GetSupervisorsByCondosID",
			Router: `/:id/supervisors`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CondosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CondosController"],
		beego.ControllerComments{
			Method: "GetByRUT",
			Router: `/rut/:rut`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CondosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CondosController"],
		beego.ControllerComments{
			Method: "AddWatcherToCondosByRUT",
			Router: `/rut/:rut/watchers`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CondosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CondosController"],
		beego.ControllerComments{
			Method: "GetSelf",
			Router: `/self`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CondosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CondosController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:CondosController"] = append(beego.GlobalControllerRouter["condo-control/controllers:CondosController"],
		beego.ControllerComments{
			Method: "GetSelfVerificationsByMonth",
			Router: `/verifications/:year/:month`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"],
		beego.ControllerComments{
			Method: "ChangeStatus",
			Router: `/:id/status/:approved`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"],
		beego.ControllerComments{
			Method: "GetByCondosID",
			Router: `/condos/self`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"],
		beego.ControllerComments{
			Method: "GetSelf",
			Router: `/self`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"] = append(beego.GlobalControllerRouter["condo-control/controllers:DeliveriesController"],
		beego.ControllerComments{
			Method: "GetByWorkersID",
			Router: `/workers/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:GoalsCommentsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:GoalsCommentsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:GoalsCommentsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:GoalsCommentsController"],
		beego.ControllerComments{
			Method: "GetAttachmentByUUID",
			Router: `/attachment/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:GoalsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:GoalsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:GoalsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:GoalsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:GoalsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:GoalsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:GoalsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:GoalsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:GoalsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:GoalsController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:GoalsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:GoalsController"],
		beego.ControllerComments{
			Method: "ChangeStatus",
			Router: `/:id/status/:completed`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:GoalsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:GoalsController"],
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

	beego.GlobalControllerRouter["condo-control/controllers:ItemsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ItemsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ItemsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ItemsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ItemsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ItemsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ItemsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ItemsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ItemsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ItemsController"],
		beego.ControllerComments{
			Method: "MakeComment",
			Router: `/:id/comment`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ItemsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ItemsController"],
		beego.ControllerComments{
			Method: "MakeCommentExecute",
			Router: `/:id/comment/:token`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ItemsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ItemsController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ItemsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ItemsController"],
		beego.ControllerComments{
			Method: "ChangeStatus",
			Router: `/:id/status/:delivered`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ItemsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ItemsController"],
		beego.ControllerComments{
			Method: "GetFilesByUUID",
			Router: `/image/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:ItemsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:ItemsController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"],
		beego.ControllerComments{
			Method: "Approve",
			Router: `/:id/status/:approved`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"],
		beego.ControllerComments{
			Method: "View",
			Router: `/:id/view/:viewed`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"],
		beego.ControllerComments{
			Method: "GetByCondosSelf",
			Router: `/condos/self`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"],
		beego.ControllerComments{
			Method: "GetImageByUUID",
			Router: `/image/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"],
		beego.ControllerComments{
			Method: "GetSelf",
			Router: `/self`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:NotificationsController"],
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
			Method: "GenerateChangePasswordToken",
			Router: `/:email/change-password/`,
			AllowHTTPMethods: []string{"post"},
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
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
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
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"],
		beego.ControllerComments{
			Method: "ChangePassword",
			Router: `/change-password/:token`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:SupervisorsController"],
		beego.ControllerComments{
			Method: "ChangePasswordSelf",
			Router: `/change-password/self`,
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

	beego.GlobalControllerRouter["condo-control/controllers:TasksController"] = append(beego.GlobalControllerRouter["condo-control/controllers:TasksController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:TasksController"] = append(beego.GlobalControllerRouter["condo-control/controllers:TasksController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:TasksController"] = append(beego.GlobalControllerRouter["condo-control/controllers:TasksController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:TasksController"] = append(beego.GlobalControllerRouter["condo-control/controllers:TasksController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:TasksController"] = append(beego.GlobalControllerRouter["condo-control/controllers:TasksController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:TasksController"] = append(beego.GlobalControllerRouter["condo-control/controllers:TasksController"],
		beego.ControllerComments{
			Method: "ChangeStatus",
			Router: `/:id/status/:approved`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:TasksController"] = append(beego.GlobalControllerRouter["condo-control/controllers:TasksController"],
		beego.ControllerComments{
			Method: "GetByCondosID",
			Router: `/condos/self`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:TasksController"] = append(beego.GlobalControllerRouter["condo-control/controllers:TasksController"],
		beego.ControllerComments{
			Method: "RequestTasks",
			Router: `/request/:id/supervisor`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:TasksController"] = append(beego.GlobalControllerRouter["condo-control/controllers:TasksController"],
		beego.ControllerComments{
			Method: "GetSelf",
			Router: `/self`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:TasksController"] = append(beego.GlobalControllerRouter["condo-control/controllers:TasksController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:TasksController"] = append(beego.GlobalControllerRouter["condo-control/controllers:TasksController"],
		beego.ControllerComments{
			Method: "GetByWorkersID",
			Router: `/workers/:id`,
			AllowHTTPMethods: []string{"get"},
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
			Method: "Comment",
			Router: `/:id/comment`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"] = append(beego.GlobalControllerRouter["condo-control/controllers:VerificationsController"],
		beego.ControllerComments{
			Method: "AddImage",
			Router: `/:id/image`,
			AllowHTTPMethods: []string{"put"},
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
			Method: "GetImagesByUUID",
			Router: `/image/:uuid`,
			AllowHTTPMethods: []string{"get"},
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
			Method: "GenerateChangePasswordToken",
			Router: `/:email/change-password/`,
			AllowHTTPMethods: []string{"post"},
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
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
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
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WatchersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WatchersController"],
		beego.ControllerComments{
			Method: "GetVerificationsByDate",
			Router: `/:id/verifications/:date`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WatchersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WatchersController"],
		beego.ControllerComments{
			Method: "GetWatchersVerificationsByMonth",
			Router: `/:id/verifications/:year/:month`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WatchersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WatchersController"],
		beego.ControllerComments{
			Method: "ChangePassword",
			Router: `/change-password/:token`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WatchersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WatchersController"],
		beego.ControllerComments{
			Method: "ChangePublicInfo",
			Router: `/change-public-info`,
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

	beego.GlobalControllerRouter["condo-control/controllers:WatchersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WatchersController"],
		beego.ControllerComments{
			Method: "GetByUsername",
			Router: `/username/:username`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WebSocketController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WebSocketController"],
		beego.ControllerComments{
			Method: "Join",
			Router: `/join`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WebSocketController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WebSocketController"],
		beego.ControllerComments{
			Method: "Test",
			Router: `/test`,
			AllowHTTPMethods: []string{"get"},
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
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WorkersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WorkersController"],
		beego.ControllerComments{
			Method: "Approve",
			Router: `/:id/approve`,
			AllowHTTPMethods: []string{"patch"},
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
			Method: "DownloadAssistancesDataByMonth",
			Router: `/:id/data/:year/:month/download`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["condo-control/controllers:WorkersController"] = append(beego.GlobalControllerRouter["condo-control/controllers:WorkersController"],
		beego.ControllerComments{
			Method: "DownloadAssistancesDataByYear",
			Router: `/:id/data/:year/download`,
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
			Method: "GetFaceByUUID",
			Router: `/face/:uuid`,
			AllowHTTPMethods: []string{"get"},
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
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
