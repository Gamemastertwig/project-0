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
	minion.name = "Fred"
	minion.class = "Unknown"
	minion.ability = {0, 0, 0, 0, 0, 0}
}

// Main Function - This is where it all starts!
func main() {


	fmt.Println(minion)
}
