package routes

import (
	"fmt"
	"net/http"

	"storality.com/storality/internal/models"
)

func CollectionListing(w http.ResponseWriter, collection *models.Collection) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Listing all %s</h1>", collection.Plural)
}