package routes

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"storality.com/storality/internal/app"
	"storality.com/storality/internal/helpers/exceptions"
	"storality.com/storality/internal/helpers/input"
	"storality.com/storality/internal/helpers/transforms"
	"storality.com/storality/internal/models"
)

func (route *Base) Record(app *app.Core, w http.ResponseWriter, r *http.Request, collection *models.Collection, param string) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		w.Header().Set("Allowed", fmt.Sprintf("%s, %s", http.MethodGet, http.MethodPost))
		exceptions.ClientError(w, http.StatusMethodNotAllowed)
		return
	}
	_, err := strconv.Atoi(param)
	if r.Method == http.MethodPost {
		if err != nil {
			newRecord(w, r, route, app.DB.Records, collection)
		} else {
			updateRecord(w, r, app.DB.Records, collection, param)
		}
	} else {
		getRecord(route, w, r, app.DB.Records, collection, param)
	}
}

func getRecord(route *Base,w http.ResponseWriter, r *http.Request, model *models.RecordModel, collection *models.Collection, param string) {
	data := route.Template.CreateTemplateData(r)
	data.Title = transforms.Capitalize(collection.Name)
	data.Collections = route.Collections
	data.Collection = collection
	id, err := strconv.Atoi(param)
	if err == nil {
		record, err := model.FindById(id)
		if err != nil {
			if errors.Is(err, exceptions.ErrNoRecord) {
				exceptions.NotFound(w)
			} else {
				exceptions.ServerError(w, err)
			}
			return
		}
		data.Record = record
	} else {
		data.Record = &models.Record{
			Title: "",
			Content: "",
		}
	}
	route.Template.Render(w, http.StatusOK, "record.html", data)
}

func newRecord(w http.ResponseWriter, r *http.Request, route *Base, model *models.RecordModel, collection *models.Collection) {
	title := input.Sanitize(r.FormValue("title"))
	content := input.Sanitize(r.FormValue("content"))
	
	id, err := model.Insert(title, transforms.Slugify(title), content, *collection)
	if err != nil {
		exceptions.ServerError(w, err)
	}
	http.Redirect(w, r, fmt.Sprintf("%s%s/%d", route.BasePath, transforms.Slugify(collection.Plural), id), http.StatusSeeOther)
}

func updateRecord(w http.ResponseWriter, r *http.Request, model *models.RecordModel, collection *models.Collection, param string) {

}