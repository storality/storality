package templates

import (
	"html/template"
	"net/http"

	"storality.com/storality/internal/models"
)

type Engine struct {
	BasePath 		string
	Route				string
	Title 			string
	Collection 	*models.Collection
	Collections []*models.Collection
	Cache				map[string]*template.Template
}

func (engine *Engine) CreateTemplateData(r *http.Request) *Engine {
	return &Engine{
		BasePath: engine.BasePath,
		Route: r.URL.Path,
	}
}