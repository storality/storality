package app

import (
	"log"
	"net/http"

	"storality.com/storality/config"
	"storality.com/storality/internal/db"
)

type Core struct {
	DB *db.DB
	Router *http.ServeMux
	Config *config.Config
}

func Bootstrap(cfg config.Config, router *http.ServeMux) *Core {
	log.SetFlags(log.Llongfile)
	app := &Core{}
	app.Config = &cfg
	app.Router = router
	app.DB = db.Connect(cfg.DataDir)
	return app
}