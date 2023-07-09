package routes

import (
	"fmt"
	"net/http"

	"storality.com/storality/internal/app"
)

func Handle(app *app.Core, basePath string) {
	app.Router.HandleFunc(basePath, index)
	app.Router.HandleFunc(basePath + "login", login)
	app.Router.HandleFunc(basePath + "register", register)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "DASHBOARD")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "LOGIN")
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "REGISTER")
}