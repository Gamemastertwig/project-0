// Project-O
// Character Generator
package main

import (
	"fmt"
	"strings"
)

type npc struct {
	name   string
	class  string
	abiity []int
}

var minion npc

func inti() {

}

func setName(name string) {
	minion.name = name
}

func setControlledName() {
	var n string
	fmt.Println("What is your Character's Name?")
	fmt.Scanln(&n)
	minion.name = n
}

func setClass(class string) {
	minion.class = class
}

func setControlledClass() {
	a := false
	for !a {
		fmt.Println("What is your class? \n[F]ighter \n[A]rcher \n[H]ealer")
		fmt.Println("[S]orcerer \n[B]bard \n[L]ibrarian")
		var c string
		fmt.Scan(&c)
		c = strings.ToLower(c)

		// placeholder will need to get this from a config file later
		switch c {
		case "f", "fighter":
			minion.class = "Fighter"
			a = true
		case "a", "archer":
			minion.class = "Archer"
			a = true
		case "h", "healer":
			minion.class = "Healer"
			a = true
		case "s", "sorcerer":
			minion.class = "Sorcerer"
			a = true
		case "b", "bard":
			minion.class = "Bard"
			a = true
		case "l", "librarian":
			minion.class = "Librarian"
			a = true
		default:
			fmt.Println("Invalid class selected! Please try again...")
		}
	}
}

func setAbilities() {
	//placeholed will want to get this from a config file later (o its not hard coded)
	allAbilities := [][]int{
		{13, 11, 12, 9, 10, 8}, //Melee
		{11, 13, 12, 10, 9, 8}, //Ranged
		{10, 8, 12, 9, 13, 11}, //Divine
		{8, 12, 10, 13, 9, 11}, //Arcane(INT)
		{8, 12, 10, 11, 9, 13}, //Arcane(CHA)
		{10, 12, 11, 13, 8, 9}, //Skill
	}

	currentClass := minion.class
	var classInt int

	// placeholder will need to get this from a config file later
	switch currentClass {
	case "Fighter":
		classInt = 0
	case "Archer":
		classInt = 1
	case "Healer":
		classInt = 2
	case "Sorcerer":
		classInt = 3
	case "Bard":
		classInt = 4
	case "Librarian":
		classInt = 5
	default:
		fmt.Println("No valid class found!")
	}

	if classInt >= 0 && classInt < len(allAbilities[classInt]) {
		minion.abiity = allAbilities[classInt]
	}
}

// Main Function - This is where it all starts!
func main() {
	setControlledName()
	setControlledClass()
	setAbilities()

	fmt.Println(minion)
}
