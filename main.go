package main

import (
	"text/template"
	"os"
	"fmt"
)

var details *template.Template //define a global pointer to a template

func init() {
	details = template.Must(template.ParseGlob("templates/*")) //gives you back a pointer to a template and checks for errors
}

type Person struct {
	Name string //these must be capital names, which make them public for other classes to see
	BirthYear int
	BirthMonth int
	BirthDay int
}

type Professor struct {
	P1 Person
	Expertise string
}

func (p Person) displayInfo() {
	fmt.Printf("Person: %s  Date of Birth: %d-%d-%d \n", p.Name, p.BirthMonth, p.BirthDay, p.BirthYear)
}

func main() {

	//details.ExecuteTemplate(os.Stdout, "index.html", nil)

	person := Person{"Ben", 4, 8, 1990}
	person.displayInfo()
	profess := Professor{Person{"Mr. Johnson",5,1990,46}, "Comp Sci"}

	cars := map[string]string {
		"Accord": "Honda",
		"Civic": "Honda",
		"CRV": "Honda",
	}

	details.ExecuteTemplate(os.Stdout,"index.html", profess)

	details.ExecuteTemplate(os.Stdout,"map.html", cars)
}
