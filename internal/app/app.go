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
	db := db.Connect(cfg.Driver, cfg.Connection)
	app := &Core{
		DB: db,
		Router: router,
		Config: &cfg,
	}
	app.verifyInstall()
	return app
}


func (app *Core) verifyInstall() {
	if _, err := os.Stat("stor_data"); os.IsNotExist(err) {
		err = os.Mkdir("stor_data", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}


}