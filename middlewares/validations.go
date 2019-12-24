package middlewares

import (
	"condo-control/controllers"
	"encoding/json"

	"github.com/astaxie/beego/context"
)

var (
	// ControllersNames ...
	ControllersNames = []string{
		"admin", "assistances", "condos", "holidays", "points", "supervisors", "verifications", "watchers", "workers", "zones",
	}
	// ExcludeUrls ...
	ExcludeUrls = map[string][]string{
		"login": {"POST"},
	}
)

// DenyAccess ...
func DenyAccess(ctx *context.Context, err error) {

	ctx.Output.SetStatus(401)
	ctx.Output.Header("Content-Type", "application/json")

	message := controllers.MessageResponse{
		Message:       "Permission Deny",
		PrettyMessage: "Permiso Denegado",
		Error:         err.Error(),
	}

	res, _ := json.Marshal(message)

	ctx.Output.Body(res)
	return
}

// MwPattern ...
type MwPattern struct {
	URL       string
	Methods   []string
	UserTypes []string
}

// GetControllerPatterns ...
func GetControllerPatterns(route string) []*MwPattern {

	/** Custom **/
/* 
	condos := []*MwPattern{
		{
			URL:       "/",
			Methods:   []string{"All"},
			UserTypes: []string{"Admin"},
		},
		{
			URL:       "/:id",
			Methods:   []string{"All"},
			UserTypes: []string{"Admin"},
		},
		{
			URL:       "/:id/trash",
			Methods:   []string{"All"},
			UserTypes: []string{"Admin"},
		},
		{
			URL:       "/:id/restore",
			Methods:   []string{"All"},
			UserTypes: []string{"Admin"},
		},
		{
			URL:       "/trashed",
			Methods:   []string{"All"},
			UserTypes: []string{"Admin"},
		},
		{
			URL:       "/self",
			Methods:   []string{"GET"},
			UserTypes: []string{"Supervisor", "Watcher"},
		},
	}

	holidays := []*MwPattern{
		{
			URL:       "/",
			Methods:   []string{"All"},
			UserTypes: []string{"Supervisor"},
		},
		{
			URL:       "/delete",
			Methods:   []string{"DELETE"},
			UserTypes: []string{"Supervisor"},
		},
	} */
	/*
		zones := []*MwPattern{
			{
				URL:       "/",
				Methods:   []string{"POST"},
				UserTypes: []string{"Supervisor"},
			},
			{
				URL:       "/:id",
				Methods:   []string{"PUT"},
				UserTypes: []string{"Supervisor"},
			},
			{
				URL:       "/delete",
				Methods:   []string{"DELETE"},
				UserTypes: []string{"Supervisor"},
			},
			{
				URL:       "/:id/restore",
				Methods:   []string{"All"},
				UserTypes: []string{"Supervisor"},
			},
			{
				URL:       "/trashed",
				Methods:   []string{"All"},
				UserTypes: []string{"Supervisor"},
			},
		}
	*/
	validations := map[string][]*MwPattern{
		/* "condos":   condos,
		"holidays": holidays,
		"zones":    zones, */
	}

	return validations[route]
}
