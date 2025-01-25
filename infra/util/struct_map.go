package util

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const APPLICATION_FORM string = "application/x-www-form-urlencoded"
const APPLICATION_JSON string = "application/json"

type StructMap struct {
	content  []byte
	request  *http.Request
	response http.ResponseWriter
}

func NewStructMap(r *http.Request) *StructMap {
	StructMap := &StructMap{
		content:  []byte{},
		request:  &http.Request{},
		response: nil,
	}

	if r != nil {
		StructMap.LoadData(r)
	}

	return StructMap
}

func (o *StructMap) GetRequestHeaderValue(key string, r *http.Request) string {

	var value string
	for _, val := range r.Header[key] {
		value = val
	}
	return value
}

func (o *StructMap) Vars(r *http.Request) map[string]string {

	vars := make(map[string]string, len(r.URL.Query()))
	for k, v := range r.URL.Query() {
		vars[k] = v[0]
	}
	return vars
}

func (o *StructMap) BindData(reference interface{}) error {

	return json.Unmarshal(o.content, reference)
}

func (o *StructMap) GenericStruct(content []byte) map[string]interface{} {

	reference := make(map[string]interface{})
	json.Unmarshal(content, &reference)
	return reference
}

func (o *StructMap) GetLoadedData() []byte {

	return o.content
}

func (o *StructMap) LoadData(r *http.Request) {

	defer r.Body.Close()

	switch r.Method {
	case "GET", "DELETE":
		field := o.Vars(r)
		reference := make(map[string]interface{})
		for key, value := range field {
			reference[key] = value
		}
		query := r.URL.Query()
		params, _ := url.ParseQuery(query.Encode())
		slice := []string{}
		for key, values := range params {
			_, exists := reference[key]
			if !exists {
				if len(values) > 1 {
					slice = append(slice, values...)
					reference[key] = strings.Join(slice, ",")
				} else {
					for key, value := range params {
						reference[key] = strings.Join(value, "")
					}
				}
			}
		}
		o.content, _ = json.Marshal(reference)
	case "PUT", "POST":
		switch o.GetRequestHeaderValue("Content-Type", r) {
		case APPLICATION_FORM:
			r.ParseForm()
			contentMap := make(map[string]string)
			for key, values := range r.Form {
				for _, value := range values {
					contentMap[key] = value
				}
			}
			o.content, _ = json.Marshal(contentMap)
		case APPLICATION_JSON:
			o.content, _ = io.ReadAll(r.Body)
			r.Body = io.NopCloser(bytes.NewBuffer(o.content))
			if r.Method == "PUT" {
				field := o.Vars(r)
				reference := make(map[string]interface{})
				json.Unmarshal(o.content, &reference)
				if len(field["id"]) > 0 {
					reference["id"] = field["id"]
				}
				o.content, _ = json.Marshal(reference)
			}
		}
	}
}
