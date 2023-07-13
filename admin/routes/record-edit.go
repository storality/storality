package routes

import (
	"fmt"
	"net/http"

	"storality.com/storality/internal/app"
	"storality.com/storality/internal/models"
)

func CollectionEdit(w http.ResponseWriter, id int, collection *models.Collection, app *app.Core) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Editing %s %d</h1>", collection.Name, id)
}