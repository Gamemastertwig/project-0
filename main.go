// Project-O
// Character Generator
package main

import (
	"fmt"

	"github.com/project-0/npccreater"
)

func init() {
	//a := []string{"warrior", "thief"}

	//Minion.name = userinputs.RequestAnswer("What is your name?")
	//Minion.class = userinputs.MultiChoiceAnswer("What is your class?", a)
	//minion.ability = {0, 0, 0, 0, 0, 0}
	npccreater.SetSex("Male")
	npccreater.SetName("Fred Smith")
	npccreater.SetClass("Fighter")
	npccreater.SetDefaultStats("Melee")
}

// Main Function - This is where it all starts!
func main() {

	fmt.Println(npccreater.Minion)
}
