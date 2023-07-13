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
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintln(w, "<h1>Dashboard</h1>")
}