package models

// TemplateData holds data to be displayed on templates
type TemplateData struct {
	IntMap    map[string]int
	StringMap map[string]string
	FloatMap  map[string]float32
	DataMap   map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
