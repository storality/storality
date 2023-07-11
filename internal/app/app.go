package app

import (
	"net/http"

	"storality.com/storality/internal/config"
	"storality.com/storality/internal/db"
)

type Core struct {
	DB *db.DB
	Router *http.ServeMux
	Config *config.Config
}

func Bootstrap(cfg config.Config, router *http.ServeMux) *Core {
	app := &Core{}
	app.Config = &cfg
	app.Router = router
	app.DB = db.Connect()
	return app
}