package routes

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"storality.com/storality/admin/templates"
	"storality.com/storality/internal/app"
	"storality.com/storality/internal/helpers/exceptions"
	"storality.com/storality/internal/helpers/flash"
	"storality.com/storality/internal/helpers/input"
	"storality.com/storality/internal/helpers/transforms"
	"storality.com/storality/internal/models"
)

func (route *Base) Record(app *app.Core, w http.ResponseWriter, r *http.Request, collection *models.Collection, param string) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost && r.Method != http.MethodDelete {
		w.Header().Set("Allowed", fmt.Sprintf("%s, %s, %s", http.MethodGet, http.MethodPost, http.MethodDelete))
		exceptions.ClientError(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(param)
	switch r.Method {
		case http.MethodGet:
			getRecord(route, w, r, app.DB.Records, collection, param)
		case http.MethodPost:
			if err != nil {
				newRecord(w, r, route, app.DB.Records, collection)
			} else {
				updateRecord(w, r, route, app.DB.Records, collection, id)
			}
		case http.MethodDelete:
			deleteRecord(w, r, route, app.DB.Records, collection, id)
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
	
	flashInfo, err := flash.Get(w, r, "flash-info")
	if err != nil {
		exceptions.ServerError(w, err)
	}
	data.Flash = &templates.Flash{
		Type: "info",
		Message: string(flashInfo),
	}

	flashError, err := flash.Get(w, r, "flash-info")
	if err != nil {
		exceptions.ServerError(w, err)
	}
	data.Flash = &templates.Flash{
		Type: "error",
		Message: string(flashError),
	}

	route.Template.Render(w, http.StatusOK, "record.html", data)
}

func newRecord(w http.ResponseWriter, r *http.Request, route *Base, model *models.RecordModel, collection *models.Collection) {
	err := r.ParseForm()
	if err != nil {
		exceptions.ClientError(w, http.StatusBadRequest)
		return
	}
	title := input.Sanitize(r.PostForm.Get("title"))
	content := input.Sanitize(r.PostForm.Get("content"))
	id, err := model.Insert(title, transforms.Slugify(title), content, *collection)
	if err != nil {
		exceptions.ServerError(w, err)
	}
	err = model.UpdateStatus(id, models.StatusPublished)
	if err != nil {
		exceptions.ServerError(w, err)
	}
	info := []byte(fmt.Sprintf("The %s has been saved.", collection.Name))
	flash.Set(w, "flash-info", info)
	http.Redirect(w, r, fmt.Sprintf("%s%s/%d", route.BasePath, transforms.Slugify(collection.Plural), id), http.StatusSeeOther)
}

func updateRecord(w http.ResponseWriter, r *http.Request, route *Base, model *models.RecordModel, collection *models.Collection, id int) {
	err := r.ParseForm()
	if err != nil {
		exceptions.ClientError(w, http.StatusBadRequest)
		return
	}
	title := input.Sanitize(r.PostForm.Get("title"))
	content := input.Sanitize(r.PostForm.Get("content"))
	err = model.Update(id, title, content)
	if err != nil {
		exceptions.ServerError(w, err)
	}
	info := []byte(fmt.Sprintf("The %s has been updated.", collection.Name))
	flash.Set(w, "flash-info", info)
	http.Redirect(w, r, fmt.Sprintf("%s%s/%d", route.BasePath, transforms.Slugify(collection.Plural), id), http.StatusSeeOther)
}

func deleteRecord(w http.ResponseWriter, r *http.Request, route *Base, model *models.RecordModel, collection *models.Collection, id int) {
	fmt.Println("DELETE")
	err := model.Delete(id)
	if err != nil {
		exceptions.ServerError(w, err)
	}
	http.Redirect(w, r, fmt.Sprintf("%s%s?message=deleted", route.BasePath, transforms.Slugify(collection.Plural)), http.StatusSeeOther)
}