package templates

import (
	"html/template"
	"os"
)

var funcs = template.FuncMap{
	"increment": func(num int) int {
		return num + 1
	},
}

func LoadTemplate(fileName string, data interface{}) {
	template, err := template.New(fileName).Funcs(funcs).ParseFiles("templates/" + fileName)

	if err != nil {
		panic(err)
	}

	err = template.Execute(os.Stdout, data)

	if err != nil {
		panic(err)
	}
}
