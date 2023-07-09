package web

import (
	"storality.com/storality/internal/app"
	"storality.com/storality/web/routes"
)

type Web struct {}

func Init(app *app.Core) (*Web, error) {
	routes.Handle(app)
	return &Web{}, nil
}