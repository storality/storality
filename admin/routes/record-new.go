package routes

import (
	"fmt"
	"net/http"

	"storality.com/storality/internal/models"
)

func CollectionNew(w http.ResponseWriter, collection *models.Collection) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Creating new %s</h1>", collection.Name)
}