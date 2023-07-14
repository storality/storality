package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"storality.com/storality/admin/routes"
	"storality.com/storality/internal/app"
	"storality.com/storality/internal/helpers/exceptions"
	"storality.com/storality/internal/helpers/shout"
	"storality.com/storality/internal/helpers/transforms"
	"storality.com/storality/internal/models"
)

func router(app *app.Core, admin *Admin) {
	go func() {
		collections, err := app.DB.Collections.FindAll()
		if err != nil {
			shout.Error.Fatal(err)
		}
		generateRoutes(app, admin.basePath, collections)
	}()

	app.Router.HandleFunc(fmt.Sprintf("%sinstall", admin.basePath), func(w http.ResponseWriter, r *http.Request){
		routes.Install(w, r, admin.basePath, app)
	})

	app.Router.HandleFunc(admin.basePath, func(w http.ResponseWriter, r *http.Request) {
		routes.Index(w, r, admin.basePath)
	})
}

func generateRoutes(app *app.Core, basePath string, collections []*models.Collection) {
	for _, col := range collections {
		collection := col
		app.Router.HandleFunc(fmt.Sprintf("%s%s/", basePath, transforms.Slugify(collection.Plural)), func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			param := strings.TrimPrefix(path, fmt.Sprintf("%s%s/", basePath, transforms.Slugify(collection.Plural)))
			id, err := strconv.Atoi(param)
			if err != nil {
				if param == "new" {
					routes.CollectionNew(w, collection)
				} else {
					if param == "" {
						routes.CollectionListing(w, collection)
					} else {
						exceptions.NotFound(w)
					}
				}
			} else {
				routes.CollectionEdit(w, id, collection, app)
			}
		})
	}
}

