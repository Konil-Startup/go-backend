package models

type TemplateData struct {
	Data map[string]interface{}

	CSRFToken string // security measure
	Flash     string // flash-message to end user
	Warning   string
	Error     error
}
