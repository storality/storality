package routes

import (
	"net/http"

	"storality.com/storality/internal/helpers/transforms"
	"storality.com/storality/internal/models"
)

func (route *Base) Records(w http.ResponseWriter, r *http.Request, collection *models.Collection) {
	r.Header.Set("Allow", "GET")
	data := route.Template.CreateTemplateData(r)
	data.Title = transforms.Capitalize(collection.Plural)
	data.Collections = route.Collections
	data.Collection = collection
	route.Template.Render(w, http.StatusOK, "records.html", data)
}