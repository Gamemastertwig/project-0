// Project-O
// Character Generator
package main

import (
	"fmt"

	"github.com/project-0/npccreater"
)

type npc struct {
	name    string
	class   string
	ability []int
}

// Minion is a global verable storing a NPC stats
var Minion npc

func inti() {

	//	a := []string{"warrior", "thief", "mage"}

	//Minion.name = npccreater.RequestName("What is your name?")
	//Minion.class = npccreater.RequestClass("What is your class?", a)
	//minion.ability = {0, 0, 0, 0, 0, 0,}
}

// Main Function - This is where it all starts!
func main() {
	a := []string{"warrior", "thief", "mage"}

	Minion.name = npccreater.RequestName("What is your name?")
	Minion.class = npccreater.RequestClass("What is your class?", a)

	fmt.Println(Minion)
}
