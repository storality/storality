package admin

import (
	"html/template"

	"storality.com/storality/internal/app"
	"storality.com/storality/internal/helpers/shout"
)

type Admin struct {
	basePath 			string
	templateCache map[string]*template.Template
}

func Run(app *app.Core, headless bool) (*Admin, error) {
	admin := Admin{}

	admin.basePath = "/admin/"
	if headless {
		admin.basePath = "/"
	}

	templateCache, err := CacheTemplates()
	if err != nil {
		shout.Error.Fatal(err)
	}

	admin.templateCache = templateCache
	router(app, &admin)
	return &admin, nil
}