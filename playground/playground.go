package playground

import (
	"fmt"
)

var words  = [][]string{
	{"break", "lake", "go"},
	{"are", "how", "are"},
	{"nada", "then", "are"},
}

func PlayInitDeclaration(){
	var name, desc string = "name1", "descripcion1"

	var colours  = []string{ "Blue", "Red", "Yellow"}

	fmt.Println(name, desc)
	fmt.Println(colours)

	search("are")

}


func search(w string) {
	// DoSearch:
		for  i:=0; i < len(words); i++ {
			for  k:= 0; k < len(words[i]); k++ {
				if words[i][k] == w {
					fmt.Println("Found", w)
					// continue DoSearch
				}
			}
		}
}