package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"storality.com/storality/admin/routes"
	"storality.com/storality/internal/app"
	"storality.com/storality/internal/helpers/shout"
	"storality.com/storality/internal/helpers/transforms"
)

func router(app *app.Core, basePath string) {
	collections, err := app.DB.Collections.FindAll()
	if err != nil {
		shout.Error.Fatal(err)
	}

	app.Router.HandleFunc(basePath, func(w http.ResponseWriter, r *http.Request){
		routes.Dashboard(w, r, basePath)
	})

	for _, col := range collections {
		collection := col
		app.Router.HandleFunc(fmt.Sprintf("%s%s/", basePath, transforms.Slugify(collection.Plural)), func(w http.ResponseWriter, r *http.Request){
			path := r.URL.Path
			param := strings.TrimPrefix(path, fmt.Sprintf("%s%s/", basePath, transforms.Slugify(collection.Plural)))
			id, err := strconv.Atoi(param)
			if err != nil {
				if param == "new" {
					routes.CollectionNew(w, collection)
				} else {
					routes.CollectionListing(w, collection)
				}
			} else {
				routes.CollectionEdit(w, id, collection, app)
			}
		})
	}
}