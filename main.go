package main

import (
	"os"
	"text/template"
)

func main() {
	template, err := template.New("greeting").Parse("Hola, mi nombre es {{.}}")

	if err != nil {
		panic(err)
	}

	err = template.Execute(os.Stdout, "Erick")

	if err != nil {
		panic(err)
	}
}
