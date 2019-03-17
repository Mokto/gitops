package templates

import (
	"bytes"
	"gitops/backend/utils"
	"html/template"
	"io/ioutil"
	"log"
	"strings"
)

// WriteValues replace the helm values.yaml
func WriteValues(basePath string, templateValues ValuesTemplate) (err error) {
	path := utils.ComposeStrings(basePath, "/chart/values.yaml")

	templateValues.BranchUrlSafe = strings.ReplaceAll(strings.ReplaceAll(templateValues.Branch, "/", "."), "/", "-")

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, templateValues)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(path, b.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// ValuesTemplate is the template used to replace values
type ValuesTemplate struct {
	Branch  string
	BranchUrlSafe string
	Secrets map[string]interface{}
}
