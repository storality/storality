package routes

import (
	"fmt"
	"net/http"

	"storality.com/storality/internal/models"
)

func CollectionNew(w http.ResponseWriter, collection *models.Collection) {
	fmt.Fprintf(w, "Creating new %s", collection.Name)
}