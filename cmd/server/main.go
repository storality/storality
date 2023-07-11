package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"storality.com/storality/admin"
	"storality.com/storality/config"
	"storality.com/storality/internal/app"
	"storality.com/storality/internal/helpers/logging"
	"storality.com/storality/web"
)

func main() {
	port := flag.Int("port", 3000, "The port")
	headless := flag.Bool("headless", false, "Run in headless mode")
	dataDir := flag.String("data-dir", "stor_data", "The directory for the data")
	flag.Parse()

	config := *config.Load(*port, *headless, *dataDir, "server")

	router := http.NewServeMux()
	app := app.Bootstrap(config, router)

	_, err := admin.Run(app, config.Headless)
	if err != nil {
		log.Fatal(err)
	}

	if !*headless {
		_, err = web.Run(app)
		if err != nil {
			log.Fatal(err)
		}
	}

	server := &http.Server{
		Addr: ":" + fmt.Sprint(config.Port),
		Handler: router,
	}

	logging.Serve.Printf("Starting server on :%d", config.Port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}