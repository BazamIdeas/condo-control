package middlewares

import (
	"GASE/controllers"
	"encoding/json"

	"github.com/astaxie/beego/context"
)

var (
	//ControllersNames ...
	ControllersNames = []string{
		"portfolios", "activities", "carts", "clients", "countries", "coupons", "currencies", "gateways", "images", "locations", "orders", "prices", "sectors", "services", "briefs", "users",
	}
)

//DenyAccess =
func DenyAccess(ctx *context.Context, err error) {

	ctx.Output.SetStatus(401)
	ctx.Output.Header("Content-Type", "application/json")

	message := controllers.MessageResponse{
		Message:       "Permission Deny",
		PrettyMessage: "Permiso Denegado",
		Code:          002,
		Error:         err.Error(),
	}

	res, _ := json.Marshal(message)

	ctx.Output.Body([]byte(string(res)))
	return
}

//GetURLMapping =
func GetURLMapping(route string) (validation map[string][]string) {

	carts := map[string][]string{
		";GET":            {"Admin"},
		"/:id;GET,DELETE": {"Admin"},
	}

	clients := map[string][]string{
		";GET":              {"Admin"},
		"/email/:email;GET": {"Admin", "Client"},
		//"/:id;GET,PUT,DELETE": {"Client"},
	}

	countries := map[string][]string{
		"/:id;GET,PUT,DELETE": {"Admin"},
	}

	coupons := map[string][]string{
		";GET":                {"Admin"},
		"/:id;GET,PUT,DELETE": {"Admin"},
	}

	currencies := map[string][]string{
		"/:id;PUT,DELETE": {"Admin"},
	}

	gateways := map[string][]string{
		"/:id;PUT,DELETE": {"Admin"},
	}

	images := map[string][]string{
		";GET":                {"Admin"},
		"/:id;GET,PUT,DELETE": {"Admin"},
	}

	locations := map[string][]string{
		"/:id;PUT,DELETE": {"Admin"},
	}

	orders := map[string][]string{
		";GET":            {"Admin"},
		"/:id;PUT,DELETE": {"Client", "Admin"},
	}

	prices := map[string][]string{
		";GET":                {"Admin"},
		"/:id;GET,PUT,DELETE": {"Admin"},
	}

	sectors := map[string][]string{
		"/:id;PUT,DELETE": {"Admin"},
	}

	services := map[string][]string{
		"/:id;PUT,DELETE": {"Admin"},
	}

	briefs := map[string][]string{
		"/:id;PUT,DELETE": {"Admin"},
	}

	activities := map[string][]string{
		"/:id;PUT,DELETE": {"Admin"},
	}

	users := map[string][]string{
		";GET":                {"Admin"},
		"/:id;GET,PUT,DELETE": {"Admin"},
	}

	portfolios := map[string][]string{
		"/:id;PUT,DELETE": {"Admin"},
	}

	payments := map[string][]string{
		";GET":                {"Admin"},
		"/:id;GET,PUT,DELETE": {"Admin"},
	}

	forms := map[string][]string{
		";GET":                {"Admin"},
		"/:id;GET,PUT,DELETE": {"Admin"},
	}

	validations := map[string]map[string][]string{
		"carts":            carts,
		"clients":          clients,
		"countries":        countries,
		"coupons":          coupons,
		"currencies":       currencies,
		"gateways":         gateways,
		"images":           images,
		"locations":        locations,
		"orders":           orders,
		"prices":           prices,
		"sectors":          sectors,
		"services":         services,
		"briefs":           briefs,
		"users":            users,
		"portfolios":       portfolios,
		"payments-methods": payments,
		"forms":            forms,
		"activities":       activities,
	}

	return validations[route]
}
