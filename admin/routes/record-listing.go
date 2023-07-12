package routes

import (
	"fmt"
	"net/http"

	"storality.com/storality/internal/models"
)

func CollectionListing(w http.ResponseWriter, collection *models.Collection) {
	fmt.Fprintf(w, "Listing all %s", collection.Plural)
}