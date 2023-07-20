package routes

import (
	"net/http"

	"storality.com/storality/internal/helpers/exceptions"
)

func (route *Base) Index(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Allow", "GET")
	if r.URL.Path != route.BasePath {
		exceptions.NotFound(w)
		return
	}
	data := route.Template.CreateTemplateData(r)
	data.Title = "Dashboard"
	data.Collections = route.Collections
	route.Template.Render(w, http.StatusOK, "index.html", data)
}