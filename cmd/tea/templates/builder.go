package templates

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"text/template"
)

func GenerateTemplate(absPath string, data interface{}) string {
	template := WriteTemplate(ReadTemplate(absPath), data)

	return template
}

func ReadTemplate(absPath string) string {
	file, err := filepath.Abs(absPath)
	if err != nil {
		panic(err)
	}

	result, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(result)
}

func WriteTemplate(content string, data interface{}) string {
	tmpl, err := template.New("template").Parse(content)
	if err != nil {
		panic(err)
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, data)
	if err != nil {
		panic(err)
	}

	return result.String()
}
