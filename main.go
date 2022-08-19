package main

import (
	"os"
	"text/template"

	"github.com/Erickype/TemplatesGo/models"
)

func main() {

	person := &models.Person{Nombre: "Erick", Edad: 22}

	template, err := template.New("greetings.txt").ParseFiles("templates/greetings.txt")

	if err != nil {
		panic(err)
	}

	err = template.Execute(os.Stdout, person)

	if err != nil {
		panic(err)
	}
}
