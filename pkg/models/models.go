package models

type TemplateData struct {
	Data []map[string]interface{}

	CSRFToken string // security measure
	Flash     string // flash-message to end user
	Warning   string
	Error     error

	Title      string
	Company    string
	Occupation string
	Years      string
	Price      int
	ImageLink  string
}
