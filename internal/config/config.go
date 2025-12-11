package config

import (
	"html/template"
	"log"

	"github.com.br/Leodf/bookings/internal/model"
	"github.com/alexedwards/scs/v2"
)

// AppConffig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	MailChan      chan model.MailData
}
