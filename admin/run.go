package admin

import (
	"storality.com/storality/internal/app"
)

type Admin struct {
	basePath	string
}

func Run(app *app.Core, headless bool) (*Admin, error) {
	admin := Admin{}

	admin.basePath = "/admin/"
	if headless {
		admin.basePath = "/"
	}

	admin.router(app)
	return &admin, nil
}