package routes

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, basePath string) {
	if r.URL.Path != basePath {
		http.NotFound(w, r)
		return
	}

	

}