package main

import (
	"flag"
	"fmt"
	"net/http"

	"storality.com/storality/admin"
	"storality.com/storality/internal/app"
	"storality.com/storality/internal/config"
	"storality.com/storality/internal/helpers/shout"
	"storality.com/storality/web"
)

func main() {
	port := flag.Int("port", 3000, "The port")
	headless := flag.Bool("headless", false, "Run in headless mode")
	flag.Parse()

	config := *config.Write(*port, *headless, "server")

	router := http.NewServeMux()
	app := app.Bootstrap(config, router)

	_, err := admin.Run(app, config.Headless)
	if err != nil {
		shout.Error.Fatal(err)
	}

	if !config.Headless {
		_, err = web.Run(app)
		if err != nil {
			shout.Error.Fatal(err)
		}
	}

	server := &http.Server{
		Addr: ":" + fmt.Sprint(config.Port),
		Handler: router,
	}
	fmt.Printf("\033[38;5;209mStarting server on :%d\033[0m\n", config.Port)
	err = server.ListenAndServe()
	if err != nil {
		shout.Error.Fatal(err)
	}
}