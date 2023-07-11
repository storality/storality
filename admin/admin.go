package admin

import (
	"storality.com/storality/admin/routes"
	"storality.com/storality/internal/app"
)

type Admin struct {}

func Run(app *app.Core, headless bool) (*Admin, error) {
	basePath := "/admin/"
	if headless {
		basePath = "/"
	}
	routes.Handle(app, basePath)
	return &Admin{}, nil
}