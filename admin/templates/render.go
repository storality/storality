package templates

import (
	"bytes"
	"fmt"
	"net/http"

	"storality.com/storality/internal/helpers/exceptions"
)

func (engine *Engine) Render(w http.ResponseWriter, status int, view string, data *Engine) {
	ts, ok := engine.Cache[view]
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