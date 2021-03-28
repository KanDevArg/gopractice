package solid

import (
	"fmt"
)

type (

	human struct {
		name string
	}

	teacher struct {
		human
	}

	writer struct {
		human
	}

	printer struct {
	}

	person interface {
		 getName() string
	}
)

func (h human) getName() string {
	return h.name
}

func (printer) info(p person) {
	fmt.Println("Name: ", p.getName())
}

func substitutionPrinciple(){

	fmt.Println("Substitution Principle")

	print := printer{}

	t := teacher{
		human{ name: "Teacher name"},
	}

	w := writer{
		human{ name: "Writer name"},
	}

	print.info(t)

	print.info(w)

	fmt.Printf("\n")
}




