package middlewares

import (
	"condo-control/controllers"
	"errors"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

//LoadMiddlewares ...
func LoadMiddlewares() {
	for _, controllerName := range ControllersNames {
		// Get MwPatterns
		MwPatterns := GetControllerPatterns(controllerName)

		for _, MwPattern := range MwPatterns {
			filterURL := controllerName + MwPattern.URL

			beego.InsertFilter("/*/"+filterURL, beego.BeforeRouter, Middleware(controllerName, MwPattern))
		}
	}
}

//Middleware ...
func Middleware(controllerName string, pattern *MwPattern) func(ctx *context.Context) {

	return func(ctx *context.Context) {

		token := ctx.Input.Header("Authorization")

		verifyToken := true
		// If the url is a excluded url then dont verify the token
		for excludeURL, URLMethods := range ExcludeUrls {
			if strings.Contains(ctx.Input.URL(), excludeURL) {
				verifyToken = false

				for _, m := range URLMethods {
					//If is a optional method like carts
					method := ctx.Input.Method()
					if method != m {
						verifyToken = true
						break
					}
				}
			}
		}

		// Return middleware
		if !verifyToken {
			// beego.Debug("Exclude Url: Dont verify token")
			return
		}

		/* DEBUG */
		beego.Debug("Url: ", controllerName, "| Pattern: ", pattern.URL)
		beego.Debug("Method: ", ctx.Input.Method())
		beego.Debug("Allowed users: ", pattern.UserTypes)

		// Verify global methods
		methods := []string{
			"POST", "DELETE", "PUT", "OPTIONS",
		}

		for _, method := range methods {
			if ctx.Input.Method() == method {
				if method == "OPTIONS" {
					beego.Debug("OPTIONS METHOD: Dont verify token")
					return
				}

				// Deny Access if the token is empty
				if token == "" {
					err := errors.New("Token No enviado")
					DenyAccess(ctx, err)
					return
				}

				_, err := controllers.VerifyToken(token, "Admin")

				if err != nil {
					// beego.Debug("Token error in methods")
					if err.Error() != "Invalid User" {
						err := errors.New("Token Invalido")
						DenyAccess(ctx, err)
						return
					}

					err := errors.New("Usuario Invalido")
					DenyAccess(ctx, err)
					return
				}
			}
		}

		// If Usertype is Guest
		if pattern.UserTypes[0] == "Guest" {
			return
		}

		// Deny Access By Default
		denyAccess := true

		// Deny Access if the token is empty
		if token == "" {
			err := errors.New("Token No enviado")
			DenyAccess(ctx, err)
			return
		}

		for _, userType := range pattern.UserTypes {
			user, err := controllers.VerifyToken(token, userType)

			if err == nil {
				if userType == user.Type {
					denyAccess = false
					break
				}
			}

			if err.Error() != "Invalid User" {
				err := errors.New("Token Invalido")
				DenyAccess(ctx, err)
				return
			}
		}

		if denyAccess {
			// beego.Debug("Invalid User Type in types")
			err := errors.New("Usuario Invalido")
			DenyAccess(ctx, err)
		}
	}
}
