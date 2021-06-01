package template

import (
	"html/template"
	"path/filepath"
)

type Templates struct {
	Base *template.Template
}

func (t *Templates) Init(path string) error {
	var err error
	t.Base, err = template.ParseFiles(filepath.Join(path, "base.html"), filepath.Join(path, "nav.html"), filepath.Join(path, "footer.html"))
	return err
}
