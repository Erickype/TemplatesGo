package main

import (
	"os"
	"text/template"

	"github.com/Erickype/TemplatesGo/models"
)

func main() {

	person := &models.Person{Nombre: "Erick", Edad: 22, Hobbies: []string{"Leer", "Caminar", "Escribir"}}

	template, err := template.New("greetings2.txt").ParseFiles("templates/greetings2.txt")

	if err != nil {
		panic(err)
	}

	err = template.Execute(os.Stdout, person)

	if err != nil {
		panic(err)
	}
}
