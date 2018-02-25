package latte

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/monoculum/formam"
)

var MaxFileSize int64 = 10 * 1024 * 1024

func Bind(r *http.Request, value interface{}) error {
	switch GetRequestContentType(r) {
	case ContentTypeJSON:
		return bindJSON(r, value)
	case ContentTypeForm:
		return bindForm(r, value)
	case ContentTypeMultipartForm:
		return bindMultipartForm(r, value)
	case ContentTypeXML:
		return bindXML(r, value)
	default:
		return ErrUnsupportedType
	}
}

func bindJSON(r *http.Request, value interface{}) error {
	return json.NewDecoder(r.Body).Decode(value)
}

func bindForm(r *http.Request, value interface{}) error {
	switch err := r.ParseForm(); err {
	case nil:
		dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "form"})
		return dec.Decode(r.Form, value)
	default:
		return err
	}
}

func bindMultipartForm(r *http.Request, value interface{}) error {
	return nil
}

func bindXML(r *http.Request, value interface{}) error {
	return xml.NewDecoder(r.Body).Decode(value)
}
