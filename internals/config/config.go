package config

import (
	"text/template"
)

type AppConfig struct {
	templateCache map[string]*template.Template
	useCache      bool
	inProduction  bool
}
