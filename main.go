package main

import (
	"github.com/Erickype/TemplatesGo/models"
	"github.com/Erickype/TemplatesGo/templates"
)

func main() {

	person := &models.Person{Nombre: "Erick", Edad: 22}

	templates.LoadTemplate("greetings.txt", person)
}
