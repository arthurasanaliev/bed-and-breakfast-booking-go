package config

import (
	"html/template"
)

// AppConfig holds app configurations
type AppConfig struct {
	UseCache      bool
	InProduction  bool
	TemplateCache map[string]*template.Template
}
