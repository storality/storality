package routes

import (
	"net/http"

	"storality.com/storality/internal/helpers/transforms"
	"storality.com/storality/internal/models"
)

func (route *Base) Record(w http.ResponseWriter, r *http.Request, collection *models.Collection, id int) {
	r.Header.Set("Allow", "GET, POST")
	if r.Method == "POST" {
		
	} else {
		data := route.Template.CreateTemplateData(r)
		data.Title = transforms.Capitalize(collection.Name)
		data.Collections = route.Collections
		data.Collection = collection
		route.Template.Render(w, http.StatusOK, "record.html", data)
	}
}