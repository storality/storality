package routes

import (
	"fmt"
	"net/http"

	"storality.com/storality/internal/app"
)

func Install(w http.ResponseWriter, r *http.Request, basePath string, app *app.Core) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Install Storality</h1>")
}