// Project-O
// Character Generator
package main

import (
	"fmt"
	"strings"
)

type npc struct {
	name  string
	class string
	stat  int
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
	for a == false {
		fmt.Println("What is your class? [T]hief, [M]age, or [W]arrior")
		var c string
		fmt.Scan(&c)
		c = strings.ToLower(c)
		if c == "t" || c == "thief" {
			minion.class = "Thief"
			a = true
		} else if c == "m" || c == "mage" {
			minion.class = "Mage"
			a = true
		} else if c == "w" || c == "warrior" {
			minion.class = "Warrior"
			a = true
		} else {
			fmt.Println("invalid class selected...")
		}
	}
}

func setStat(stat int) {
	minion.stat = stat
}

// Main Function - This is where it all starts!
func main() {
	setControlledName()
	setControlledClass()

	fmt.Println(minion)
}
