package app

import (
	"net/http"

	"storality.com/storality/internal/config"
	"storality.com/storality/internal/db"
	"storality.com/storality/internal/helpers/session"
)

type Core struct {
	DB *db.DB
	Router *http.ServeMux
	Config *config.Config
	SessionManager *session.SessionManager
}

func Bootstrap(cfg config.Config, router *http.ServeMux, sessionManager *session.SessionManager) *Core {
	app := &Core{}
	app.Config = &cfg
	app.Router = router
	app.DB = db.Connect()
	app.SessionManager = sessionManager
	return app
}