package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"storality.com/storality/internal/app"
	"storality.com/storality/internal/helpers/exceptions"
	"storality.com/storality/internal/helpers/shout"
	"storality.com/storality/internal/helpers/transforms"
	"storality.com/storality/internal/models"
)

func router(app *app.Core, admin *Admin) {
	collections, err := app.DB.Collections.FindAll()
	if err != nil {
		shout.Error.Fatal(err)
	}

	filerServer := http.FileServer(http.Dir("./admin/ui/static/"))
	app.Router.Handle("/static/", http.StripPrefix("/static", filerServer))

	for _, col := range collections {
		collection := col
		app.Router.HandleFunc(fmt.Sprintf("%s%s/", admin.basePath, transforms.Slugify(collection.Plural)), func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			param := strings.TrimPrefix(path, fmt.Sprintf("%s%s/", admin.basePath, transforms.Slugify(collection.Plural)))
			id, err := strconv.Atoi(param)
			if err != nil {
				if param == "new" {
					handleRecord(w, r, admin, app, collections, collection, id)
				} else {
					if param == "" {
						handleRecords(w, r, admin, collections, collection)
					} else {
						exceptions.NotFound(w)
					}
				}
			} else {
				handleRecord(w, r, admin, app, collections, collection, id)
			}
		})
	}

	app.Router.HandleFunc(admin.basePath, func(w http.ResponseWriter, r *http.Request) {
		handleIndex(w, r, admin, collections)
	})
}

func handleIndex(w http.ResponseWriter, r *http.Request, admin *Admin, collections []*models.Collection) {
	r.Header.Set("Allow", "GET")
	if r.URL.Path != admin.basePath {
		exceptions.NotFound(w)
		return
	}
	data := admin.CreateTemplateData(r)
	data.Title = "Dashboard"
	data.Collections = collections
	admin.render(w, http.StatusOK, "index.html", data)
}

func handleRecords(w http.ResponseWriter, r *http.Request, admin *Admin, collections []*models.Collection, collection *models.Collection) {
	r.Header.Set("Allow", "GET")
	data := admin.CreateTemplateData(r)
	data.Title = transforms.Capitalize(collection.Plural)
	data.Collections = collections
	data.Collection = collection
	admin.render(w, http.StatusOK, "records.html", data)
}

func handleRecord(w http.ResponseWriter, r *http.Request, admin *Admin, app *app.Core, collections []*models.Collection, collection *models.Collection, id int) {
	r.Header.Set("Allow", "GET, POST")
	if r.Method == "POST" {
		
	} else {
		data := admin.CreateTemplateData(r)
		data.Title = transforms.Capitalize(collection.Name)
		data.Collections = collections
		data.Collection = collection
		admin.render(w, http.StatusOK, "record.html", data)
	}
}

