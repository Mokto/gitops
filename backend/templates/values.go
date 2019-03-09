package templates

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
)

// WriteValues replace the helm values.yaml
func WriteValues(templateValues ValuesTemplate) (err error) {
	tmpl, err := template.ParseFiles("/tmp/cloned-repo/app/values.yaml")
	if err != nil {
		return err
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, templateValues)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("/tmp/cloned-repo/app/values.yaml", b.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// ValuesTemplate is the template used to replace values
type ValuesTemplate struct {
	Branch  string
	Tag     string
	Secrets map[string]interface{}
}
