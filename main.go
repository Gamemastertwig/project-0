// Project-O
// Character Generator
package main

import (
	"fmt"
	"strings"
)

// Main Function - This is where it all starts!
func main() {
	var n string
	fmt.Println("What is your Character's Name?")
	fmt.Scanln(&n)
	fmt.Println("Hello, " + n)

	a := false
	for a == false {
		fmt.Println("What is your class? [T]hief, [M]age, or [W]arrior")
		var c string
		fmt.Scan(&c)
		c = strings.ToLower(c)
		if c == "t" || c == "thief" {
			fmt.Println(n + " is a Thief!")
			a = true
		} else if c == "m" || c == "mage" {
			fmt.Println(n + " is a Mage!")
			a = true
		} else if c == "w" || c == "warrior" {
			fmt.Println(n + " is a Warrior!")
			a = true
		} else {
			fmt.Println("invalid class selected...")
		}
	}
}
