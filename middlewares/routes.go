package middlewares

import (
	"GASE/controllers"
	"errors"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

//InsertFilter ...
func InsertFilter(ctrlName string, url string, pattern string, users []string) {

	beego.InsertFilter(url, beego.BeforeRouter, Middleware(ctrlName, pattern, users))

}

//LoadFilters ...
func LoadFilters() {
	for _, ctrlName := range ControllersNames {
		for pattern, users := range GetURLMapping(ctrlName) {
			if strings.Contains(pattern, ";") {
				methods := strings.Split(pattern, ";")

				InsertFilter(ctrlName, "/*/"+ctrlName+methods[0], pattern, users)
			} else {

				InsertFilter(ctrlName, "/*/"+ctrlName+pattern, pattern, users)
			}
		}
	}
}

//MiddlewareNew =
func MiddlewareNew(userTypes []string) func(ctx *context.Context) {

	return func(ctx *context.Context) {

		token := ctx.Input.Header("Authorization")

		// Deny Access if the token is empty
		if token == "" {
			err := errors.New("Token vacio")
			DenyAccess(ctx, err)
			return
		}

		denyAccess := false

		for _, userType := range userTypes {

			user, err := controllers.VerifyToken(token, userType)

			if err != nil {
				continue
			}

			userTypeDecoded := user.Type

			if userType == userTypeDecoded {
				denyAccess = false
				break
			}

		}

		if denyAccess {
			err := errors.New("Usuario Invalido")
			DenyAccess(ctx, err)
		}
	}
}

//Middleware =
func Middleware(controller string, pattern string, userTypes []string) func(ctx *context.Context) {

	return func(ctx *context.Context) {

		token := ctx.Input.Header("Authorization")

		//Verify global method

		methods := []string{
			"POST", "DELETE", "PUT", "OPTIONS",
		}

		excludeUrls := []string{
			"login", "carts", "change-password", "custom-search",
		}

		verifyToken := true

		for _, excludeURL := range excludeUrls {
			// If the url is a excluded url then dont verify the token
			if strings.Contains(ctx.Input.URL(), excludeURL) {

				verifyToken = false

				//If is a optional method like carts
				switch ctx.Input.Method() {
				case "DELETE":
					if excludeURL == "carts" {
						verifyToken = true
						break
					}
				case "GET":
					if excludeURL == "carts" {
						verifyToken = true
						break
					}
				}
				break
			}
		}

		// Return middleware
		if !verifyToken {
			return
		}

		// Verify if the method require a admin token
		for _, method := range methods {
			if ctx.Input.Method() == method {

				if method == "OPTIONS" {
					return
				}

				_, err := controllers.VerifyToken(token, "Admin")

				if err != nil {
					beego.Debug("No valido")
					err := errors.New("Usuario Invalido")
					DenyAccess(ctx, err)
				}
			}
		}

		if userTypes[0] == "Guest" {
			return
		}

		urlMapping := GetURLMapping(controller)

		beego.Debug("Url: ", controller, "| Match: ", pattern)
		beego.Debug("Allowed users: ", userTypes)
		beego.Debug("Method: ", ctx.Input.Method())

		denyAccess := true

		// if has methods
		if strings.Contains(pattern, ";") {
			splitted := strings.Split(pattern, ";")
			methods := splitted[1]

			if strings.Contains(methods, ctx.Input.Method()) {

				userTypes = urlMapping[pattern]

			}
		}

		beego.Debug("Allowed users: ", userTypes)

		// Deny Access if the token is empty
		if token == "" {
			err := errors.New("Token Invalido")
			DenyAccess(ctx, err)
			return
		}

		for _, userType := range userTypes {

			user, err := controllers.VerifyToken(token, userType)

			if err != nil {
				continue
			}

			userTypeDecoded := user.Type

			if userType == userTypeDecoded {
				denyAccess = false
				break
			}

		}

		if denyAccess {
			err := errors.New("Usuario Invalido")
			DenyAccess(ctx, err)
		}
	}
}
