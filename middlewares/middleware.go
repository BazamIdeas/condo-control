package middlewares

import (
	"condo-control/controllers"
	"encoding/json"

	"github.com/astaxie/beego/context"
)

var (
	//ControllersNames ...
	ControllersNames = []string{}
)

//DenyAccess =
func DenyAccess(ctx *context.Context, err error, statusCode int) {

	ctx.Output.SetStatus(statusCode)
	ctx.Output.Header("Content-Type", "application/json")

	message := controllers.MessageResponse{
		Message:       "Permission Deny",
		PrettyMessage: "Permiso Denegado",
		Code:          002,
		Error:         err.Error(),
	}

	res, _ := json.Marshal(message)

	ctx.Output.Body(res)
}

//GetURLMapping =
func GetURLMapping(route string) (validation map[string][]string) {

	zones := map[string][]string{
		"/self;GET": {"Watcher", "Supervisor"},
	}

	verifications := map[string][]string{
		"/verifications/zones;POST": {"Watcher"},
	}

	validations := map[string]map[string][]string{
		"zones":         zones,
		"verifications": verifications,
	}

	return validations[route]
}
