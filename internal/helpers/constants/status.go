package constants

type Status string

const (
	StatusPublished Status = "published"
	StatusDraft     Status = "draft"
	StatusTrashed   Status = "trashed"
)

var Statuses = []Status{
	StatusPublished,
	StatusDraft,
	StatusTrashed,
}