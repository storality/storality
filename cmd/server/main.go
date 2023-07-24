package main

import (
	"fmt"
	"net/http"

	"storality.com/storality/admin"
	"storality.com/storality/internal/app"
	"storality.com/storality/internal/config"
	"storality.com/storality/internal/helpers/shout"
	"storality.com/storality/internal/middleware"
	"storality.com/storality/web"
)

func main() {
	config := config.New("server")
	router := http.NewServeMux()
	app := app.Bootstrap(*config, router)

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
		Handler: middleware.RecoverPanic(middleware.LogRequest(middleware.SecureHeaders(router))),
	}
	
	fmt.Printf("\033[38;5;209mStarting server on :%d\033[0m\n", config.Port)
	err = server.ListenAndServe()
	if err != nil {
		shout.Error.Fatal(err)
	}
}