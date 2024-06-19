package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	templateCache map[string]*template.Template
	useCache      bool
	inProduction  bool
	Session       *scs.SessionManager
}
