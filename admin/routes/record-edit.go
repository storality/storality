package routes

import (
	"fmt"
	"net/http"

	"storality.com/storality/internal/app"
	"storality.com/storality/internal/models"
)

func CollectionEdit(w http.ResponseWriter, id int, collection *models.Collection, app *app.Core) {
	fmt.Fprintf(w, "Editing %s %d", collection.Name, id)
}