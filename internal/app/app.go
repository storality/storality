package app

import (
	"log"
	"net/http"
	"os"

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
	app.verifyDataDir()
	app.DB = db.Connect(cfg.Driver, cfg.Connection)
	return app
}

func (app *Core) verifyDataDir() {
	if _, err := os.Stat("stor_data"); os.IsNotExist(err) {
		err = os.Mkdir("stor_data", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}