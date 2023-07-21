package templates

import (
	"html/template"
	"net/http"

	"storality.com/storality/internal/models"
)

type Notice struct {
	Message string
	Type		string
}

type Engine struct {
	BasePath 		string
	Route				string
	Title 			string
	Collection 	*models.Collection
	Collections []*models.Collection
	Record 			*models.Record
	Records			[]*models.Record
	Cache				map[string]*template.Template
	Notice 			Notice
}

func (engine *Engine) CreateTemplateData(r *http.Request) *Engine {
	return &Engine{
		BasePath: engine.BasePath,
		Route: r.URL.Path,
	}
}