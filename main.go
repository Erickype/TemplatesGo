package main

import (
	"os"
	"text/template"

	"github.com/Erickype/TemplatesGo/models"
)

func main() {

	person := &models.Person{Nombre: "Erick", Edad: 22}

	template, err := template.New("greeting").Parse("Hola, mi nombre es {{.Nombre}} y tengo {{.Edad}} años.")

	if err != nil {
		panic(err)
	}

	err = template.Execute(os.Stdout, person)

	if err != nil {
		panic(err)
	}
}
