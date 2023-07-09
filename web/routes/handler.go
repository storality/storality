package routes

import (
	"fmt"
	"log"
	"net/http"

	"storality.com/storality/internal/app"
)

func Handle(app *app.Core) {
	app.Router.HandleFunc("/", index)
	collections, err := app.DB.Collections.FindAll()
	log.Print(collections)
	if err != nil {
		log.Fatal(err)
	}
	for _, collection := range collections {
		slug := "/" + collection.Plural
		app.Router.HandleFunc(slug, collectionListing)
		app.Router.HandleFunc(slug + "/", func(w http.ResponseWriter, r *http.Request){
			http.Redirect(w, r, slug, http.StatusPermanentRedirect)
		})
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprintln(w, "404 - Not Found")
		return
	}
	fmt.Fprintln(w, "INDEX")
}

func collectionListing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Collection Listing")
}