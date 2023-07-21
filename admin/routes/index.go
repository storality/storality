package routes

import (
	"net/http"

	"storality.com/storality/internal/app"
	"storality.com/storality/internal/helpers/exceptions"
)

func (route *Base) Index(app *app.Core, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allowed", http.MethodGet)
		exceptions.ClientError(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != route.BasePath {
		exceptions.NotFound(w)
		return
	}
	data := route.Template.CreateTemplateData(r)
	data.Title = "Dashboard"
	data.Collections = route.Collections
	route.Template.Render(w, http.StatusOK, "index.html", data)
}