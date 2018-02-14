package latte

import (
	"bytes"
	"html/template"
)

type View struct {
	Template string
	Layout   string
}

type Renderer struct {
	templates *template.Template
}

func (r *Renderer) JSON() error {
	return nil
}

func (r *Renderer) HTML() error {
	return nil
}

func (r *Renderer) Clone() *Renderer {
	return nil
}

func (r *Renderer) UseLayout(layout string) *Renderer {
	return r
}

func (r *Renderer) executeTemplate(name string, bindings interface{}) (*bytes.Buffer, error) {
	b := new(bytes.Buffer)
	return b, r.templates.ExecuteTemplate(b, name, bindings)
}

func (r *Renderer) addViewFuncs(name string, bindings interface{}) error {
	viewFuncs := template.FuncMap{
		"yield": func() (template.HTML, error) {
			buf, err := r.executeTemplate(name, bindings)
			return template.HTML(buf.String()), err
		},
	}

	if tmpl := r.templates.Lookup(name); tmpl != nil {
		tmpl.Funcs(viewFuncs)
		return nil
	}

	return ErrNotFound
}
