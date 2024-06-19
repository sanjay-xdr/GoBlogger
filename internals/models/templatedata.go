package models

type TemplateData struct {
	Data      map[string]interface{}
	IntMap    map[string]int
	Warning   string
	Success   string
	CSRFToken string
}
