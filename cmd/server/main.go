package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"storality.com/storality/admin"
	"storality.com/storality/config"
	"storality.com/storality/internal/app"
	"storality.com/storality/web"
)

func main() {
	port := flag.Int("port", 3000, "The port")
	headless := flag.Bool("headless", false, "Run in headless mode")
	driver := flag.String("db-driver", "sqlite3", "The database driver")
	connection := flag.String("db-connection", "stor_data/stor_db.db", "The database connection string")
	flag.Parse()

	router := http.NewServeMux()
	app := app.Bootstrap(*config.Load(*port, *headless, *driver, *connection, "server"), router)

	_, err := admin.Init(app, *headless)
	if err != nil {
		log.Fatal(err)
	}

	if !*headless {
		_, err = web.Init(app)
		if err != nil {
			log.Fatal(err)
		}
	}

	server := &http.Server{
		Addr: ":" + fmt.Sprint(*port),
		Handler: router,
	}

	fmt.Printf("Starting server on :%d", *port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}