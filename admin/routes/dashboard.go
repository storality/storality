package routes

import (
	"fmt"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request, basePath string) {
	if r.URL.Path != basePath {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintln(w, "Dashboard")
}