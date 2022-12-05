package config

import "html/template"

//AppConfig Holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
