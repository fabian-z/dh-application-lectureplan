package template

import (
	"html/template"
	"path/filepath"
)

type Templates struct {
	Error          *template.Template
	Events         *template.Template
	DataProtection *template.Template
}

func (t *Templates) Init(path string) error {
	var err error
	t.Error, err = template.ParseFiles(filepath.Join(path, "base.html"), filepath.Join(path, "nav.html"), filepath.Join(path, "error.html"))
	if err != nil {
		return err
	}

	t.Events, err = template.ParseFiles(filepath.Join(path, "base.html"), filepath.Join(path, "nav.html"), filepath.Join(path, "events.html"))
	if err != nil {
		return err
	}
	t.DataProtection, err = template.ParseFiles(filepath.Join(path, "base.html"), filepath.Join(path, "nav.html"), filepath.Join(path, "dataprotection.html"))
	if err != nil {
		return err
	}

	return err
}
