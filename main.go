// Project-O
// Character Generator
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Gamemastertwig/project-0/fileinputs"
	_ "github.com/Gamemastertwig/project-0/fileinputs"
	"github.com/Gamemastertwig/project-0/npc"
	"github.com/Gamemastertwig/project-0/randimizer"
)

var minion npc.Npc

var r bool         // random flag
var m bool         // make flag
var h bool         // help flag
var mArgs []string // non-flag arguments from make flag

func init() {
	flag.BoolVar(&r, "random", false, "Generates a fully random NPC using set defaults") // random
	flag.BoolVar(&m, "make", false, "Generates a NPC using user defined variables")      // make
	flag.BoolVar(&h, "help", false, "Display help menu and exits application")           // help1
	flag.BoolVar(&h, "h", false, "Display help menu and exits application")              // help2
	flag.Parse()
	if m {
		mArgs = flag.Args()
	}

	// testing variables
	minion.SetSex(randimizer.StringRandimizer([]string{"Male", "Female"}))
	if minion.Sex == "Female" {
		minion.SetName(randimizer.StringRandimizer(fileinputs.SliceFromTXTFile("gName.txt")))
	} else {
		minion.SetName(randimizer.StringRandimizer(fileinputs.SliceFromTXTFile("bName.txt")))
	}

	minion.SetClass("Fighter")
	minion.SetDefaultStats("melee") // need error control here...
	//minion.SetStats([]int{12, 12, 12, 12, 12, 12})
}

// Main Function - This is where it all starts!
func main() {
	if h {
		help()
		os.Exit(0)
	}

	minion.DisplayNpcBlock()
	//fmt.Println(minion)
	//content := fileinputs.SliceFromTXTFile("gName.txt")
	//fmt.Println(content)
	// fmt.Printf("make = %t\n", m)
	// fmt.Printf("random = %t\n", r)
	// fmt.Printf("help, h = %t\n", h)
	// if len(mArgs) != 0 {
	// 	fmt.Println(mArgs)
	// }
}

func help() {
	// displays "hardcoded" help menu
	fmt.Printf("This application is designed to help the user genrate basic NPCs.\n")
	fmt.Printf("The core is designed around the D20 Pathfinder ruleset but some\n")
	fmt.Printf("feature are available to allow for customization.\n\n")
	fmt.Printf("-make     Generates a NPC using user defined variables.\n")
	fmt.Printf("          variables are in the following order: sex name class\n")
	fmt.Printf("          abilityValue1 abilityValue2 ...\n\n")
	fmt.Printf("          example: -make Male Fred Fighter 13 11 12 9 10 8\n")
	fmt.Printf("-random   Generates a fully random NPC using set defaults.\n")
	fmt.Printf("-h, -help Displays this menu and exits the application\n")
}
