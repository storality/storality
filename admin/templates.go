package admin

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"storality.com/storality/internal/helpers/exceptions"
	"storality.com/storality/internal/models"
)

type Template struct {
	Collection *models.Collection
	Collections []*models.Collection
}

var functions = template.FuncMap{}

func CacheTemplates() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	views, err := filepath.Glob("./ui/html/views/*html")
	if err != nil {
		return nil, err
	}

	for _, view := range views {
		name := filepath.Base(view)
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/document.html")
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseFiles(view)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}

func (admin *Admin) render(w http.ResponseWriter, status int, view string, data *Template) {
	ts, ok := admin.templateCache[view]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", view)
		exceptions.ServerError(w, err)
	}

	buf := new(bytes.Buffer)
	err := ts.ExecuteTemplate(buf, "document", data)
	if err != nil {
		exceptions.ServerError(w, err)
		return
	}
	w.WriteHeader(status)
	buf.WriteTo(w)
}

func (admin *Admin) CreateTemplateData(r *http.Request) *Template {
	return &Template{}
}