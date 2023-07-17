package web

import (
	"storality.com/storality/internal/app"
)

type Web struct {}

func Run(app *app.Core) (*Web, error) {
	return &Web{}, nil
}