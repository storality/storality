package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"storality.com/storality/admin/routes"
	"storality.com/storality/admin/templates"
	"storality.com/storality/internal/app"
	"storality.com/storality/internal/helpers/exceptions"
	"storality.com/storality/internal/helpers/shout"
	"storality.com/storality/internal/helpers/transforms"
)

func (admin *Admin) router(app *app.Core) {
	collections, err := app.DB.Collections.FindAll()
	if err != nil {
		shout.Error.Fatal(err)
	}

	templateEngine := templates.Engine{
		BasePath: admin.basePath,
	}

	templateEngine.Cache, err = templateEngine.CreateCache()
	if err != nil {
		shout.Error.Fatal(err)
	}

	routes := routes.Base{
		BasePath: admin.basePath,
		Collections: collections,
		Template: &templateEngine,
	}

	filerServer := http.FileServer(http.Dir("./admin/ui/static/"))
	app.Router.Handle("/static/", http.StripPrefix("/static", filerServer))

	app.Router.HandleFunc(admin.basePath, func(w http.ResponseWriter, r *http.Request) {
		routes.Index(w, r)
	})

	for _, col := range routes.Collections {
		collection := col
		app.Router.HandleFunc(fmt.Sprintf("%s%s/", admin.basePath, transforms.Slugify(collection.Plural)), func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			param := strings.TrimPrefix(path, fmt.Sprintf("%s%s/", admin.basePath, transforms.Slugify(collection.Plural)))
			id, err := strconv.Atoi(param)
			if err != nil {
				if param == "new" {
					routes.Record(w, r, collection, id)
				} else {
					if param == "" {
						routes.Records(w, r, collection)
					} else {
						exceptions.NotFound(w)
					}
				}
			} else {
				routes.Record(w, r, collection, id)
			}
		})
	}
}

