package main

import (
	"github.com/Erickype/TemplatesGo/models"
	"github.com/Erickype/TemplatesGo/templates"
)

func main() {

	person := &models.Person{Nombre: "Erick", Edad: 22, Hobbies: []string{"Leer", "Caminar", "Escribir"}}

	templates.LoadTemplate("greetings2.txt", person)
}
