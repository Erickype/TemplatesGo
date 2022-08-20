package templates

import (
	"html/template"
	"os"
)

func LoadTemplate(fileName string, data interface{}) {
	template, err := template.New(fileName).ParseFiles("templates/" + fileName)

	if err != nil {
		panic(err)
	}

	err = template.Execute(os.Stdout, data)

	if err != nil {
		panic(err)
	}
}
