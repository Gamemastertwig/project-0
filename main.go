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
	fmt.Println("What is your class? [T]hief, [M]age, or [W]arrior")
	var c string
	fmt.Scan(&c)
	c = strings.ToLower(c)
	if c == "t" || c == "thief" {
		fmt.Println(n + " is a Thief!")
	} else if c == "m" || c == "mage" {
		fmt.Println(n + " is a Mage!")
	} else if c == "w" || c == "warrior" {
		fmt.Println(n + " is a Warrior!")
	} else {
		fmt.Println("invalid class selected...")
	}
}
