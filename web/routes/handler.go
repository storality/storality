package routes

import (
	"fmt"
	"net/http"

	"storality.com/storality/internal/app"
	"storality.com/storality/internal/helpers/shout"
)

func Handle(app *app.Core) {
	app.Router.HandleFunc("/", index)
	collections, err := app.DB.Collections.FindAll()
	if err != nil {
		shout.Error.Fatal(err)
	}
	for _, collection := range collections {
		slug := "/" + collection.Plural
		app.Router.HandleFunc(slug, func(w http.ResponseWriter, r *http.Request){
			collectionListing(w, slug, app)
		})
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

func collectionListing(w http.ResponseWriter, slug string, app *app.Core) {
	collection, err := app.DB.Collections.FindBySlug(slug)
	if err != nil {
		fmt.Fprintln(w, "404 - Not Found")
		return
	}
	fmt.Fprintln(w, "Listing for the " + collection.Name + " collection")
}