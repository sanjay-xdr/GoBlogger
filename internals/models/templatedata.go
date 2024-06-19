package models

type TemplateData struct {
	Data      map[string]interface{}
	Warning   string
	Success   string
	CSRFToken string
}
