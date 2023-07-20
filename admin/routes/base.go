package routes

import (
	"storality.com/storality/admin/templates"
	"storality.com/storality/internal/models"
)

type Base struct {
	BasePath	string
	Template 	*templates.Engine
	Collections []*models.Collection
}