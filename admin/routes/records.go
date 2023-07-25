package routes

import (
	"fmt"
	"net/http"

	"storality.com/storality/admin/templates"
	"storality.com/storality/internal/app"
	"storality.com/storality/internal/helpers/exceptions"
	"storality.com/storality/internal/helpers/transforms"
	"storality.com/storality/internal/models"
)

func (route *Base) Records(app *app.Core, w http.ResponseWriter, r *http.Request, collection *models.Collection) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allowed", http.MethodGet)
		exceptions.ClientError(w, http.StatusMethodNotAllowed)
		return
	}
	if r.Method == http.MethodGet {
		getRecords(app, w, r, collection, route)
	}
}

func getRecords(app *app.Core, w http.ResponseWriter, r *http.Request, collection *models.Collection, route *Base) {
	data := route.Template.CreateTemplateData(r)
	records, err := app.DB.Records.FindMany(&models.Filter{Collection: *collection})
	if err != nil {
		exceptions.ServerError(w, err)
	}
	if len(records) == 0 {
		data.Flash = &templates.Flash{
			Message: fmt.Sprintf("There are no %s", collection.Plural),
			Type: "info",
		}
	} else {
		data.Records = records
	}
	data.Title = transforms.Capitalize(collection.Plural)
	data.Collections = route.Collections
	data.Collection = collection
	route.Template.Render(w, http.StatusOK, "records.html", data)
}