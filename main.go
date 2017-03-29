package main

import (
	"text/template"
	"os"
	"fmt"
	"log"
)

var details *template.Template //define a global pointer to a template
var tpl *template.Template

func init() {
	details = template.Must(template.ParseGlob("templates/*")) //gives you back a pointer to a template and checks for errors
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
type course struct {
	Number string
	Name   string
	Units  string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	AcaYear string
	Fall    semester
	Spring  semester
	Summer  semester
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

	years := []year{
		year{
			AcaYear: "2020-2021",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					course{"CSCI-40", "Introduction to Programming in Go", "4"},
					course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					course{"CSCI-50", "Advanced Go", "5"},
					course{"CSCI-190", "Advanced Web Programming with Go", "5"},
					course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
		year{
			AcaYear: "2021-2022",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					course{"CSCI-40", "Introduction to Programming in Go", "4"},
					course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					course{"CSCI-50", "Advanced Go", "5"},
					course{"CSCI-190", "Advanced Web Programming with Go", "5"},
					course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, years)
	if err != nil {
		log.Fatalln(err)
	}

	details.ExecuteTemplate(os.Stdout,"index.html", profess)

	details.ExecuteTemplate(os.Stdout,"map.html", cars)
}
