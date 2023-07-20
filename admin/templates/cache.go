package templates

import (
	"html/template"
	"path/filepath"
)

func (engine *Engine) CreateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	views, err := filepath.Glob("./admin/ui/html/views/*.html")
	if err != nil {
		return nil, err
	}

	for _, view := range views {
		name := filepath.Base(view)
		ts, err := template.New(name).Funcs(functions).ParseFiles("./admin/ui/html/document.html")
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob("./admin/ui/html/partials/*.html")
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