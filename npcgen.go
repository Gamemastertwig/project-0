// Project-O
// Character Generator
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/Gamemastertwig/project-0/fileinputs"
	"github.com/Gamemastertwig/project-0/npc"
	"github.com/Gamemastertwig/project-0/randimizer"
	"github.com/Gamemastertwig/project-0/userinputs"
)

var minion npc.Npc

var r bool         // random flag
var m bool         // make flag
var h bool         // help flag
var mArgs []string // non-flag arguments from make flag
var wg sync.WaitGroup

func init() {
	// sets flag options
	flag.BoolVar(&r, "random", false, "Generates a fully random NPC using set defaults") // random
	flag.BoolVar(&m, "make", false, "Generates a NPC using user defined variables")      // make
	flag.BoolVar(&h, "help", false, "Display help menu and exits application")           // help1
	flag.BoolVar(&h, "h", false, "Display help menu and exits application")              // help2
	flag.Parse()
	if m || r {
		mArgs = flag.Args()
	}
}

// Main Function - This is where it all starts!
func main() {

	// -h, -help
	if h {
		help()
		os.Exit(0)
	} else if r { // -random
		if len(mArgs) >= 1 {
			count, err := strconv.Atoi(mArgs[0])
			if err != nil {
				fmt.Printf("value after -random is not an integer: %s\n", err)
			} else {
				wg.Add(count)
				for i := 0; i < count; i++ {
					go random() // run all on seperate go routines
				}
			}
		} else {
			wg.Add(1)
			go random()
		}
	} else if m { // -make
		custom()
	} else { // default
		userRequest()
	}
	wg.Wait()
}

func userRequest() {
	minion.SetSex(userinputs.RequestAnswer("What is your Gender?"))
	minion.SetName(userinputs.RequestAnswer("What is your Name?"))
	minion.SetClass(userinputs.RequestAnswer("What is Class?"))
	str, err := strconv.Atoi(userinputs.RequestAnswer("What is your Strength?"))
	if err != nil {
		fmt.Printf("Invalid: non-numeric value provided for Strength\n")
	}
	dex, err := strconv.Atoi(userinputs.RequestAnswer("What is your Dexterity?"))
	if err != nil {
		fmt.Printf("Invalid: non-numeric value provided for Dexterity\n")
	}
	con, err := strconv.Atoi(userinputs.RequestAnswer("What is your Consitution?"))
	if err != nil {
		fmt.Printf("Invalid: non-numeric value provided for Consitution\n")
	}
	itl, err := strconv.Atoi(userinputs.RequestAnswer("What is your Inteligence?"))
	if err != nil {
		fmt.Printf("Invalid: non-numeric value provided for Inteligence\n")
	}
	wis, err := strconv.Atoi(userinputs.RequestAnswer("What is your Wisdom?"))
	if err != nil {
		fmt.Printf("Invalid: non-numeric value provided for Wisdom\n")
	}
	cha, err := strconv.Atoi(userinputs.RequestAnswer("What is your Charisma?"))
	if err != nil {
		fmt.Printf("Invalid: non-numeric value provided for Charisma\n")
	}
	minion.SetStats([]int{str, dex, con, itl, wis, cha})

	minion.DisplayNpcBlock()
}

func help() {
	// displays "hardcoded" help menu
	fmt.Printf("This application is designed to help the user genrate basic NPCs.\n")
	fmt.Printf("The core is designed around the D20 Pathfinder ruleset but some\n")
	fmt.Printf("feature are available to allow for customization.\n\n")
	fmt.Printf("Using the application with no argument will allow the user to specify the\n")
	fmt.Printf("values based on answering prompts.\n\n")
	fmt.Printf("-make     Generates a NPC using user defined variables.\n")
	fmt.Printf("          variables are in the following order: sex name class\n")
	fmt.Printf("          str dex con int wis cha\n")
	fmt.Printf("          If values are left off they will be randomized.\n\n")
	fmt.Printf("          example: -make Male Fred Fighter Melee 13 11 12 9 10 8\n\n")
	fmt.Printf("-random   Generates a fully random NPC using set defaults. If followed by a\n")
	fmt.Printf("          value it will generate that many random NPCs\n\n")
	fmt.Printf("          example1: -random 5 || example2: -random\n\n")
	fmt.Printf("-h, -help Displays this menu and exits the application\n")
	fmt.Printf("To provide more random results please add the following text files to the\n")
	fmt.Printf("same location as this application.\n\n")
	fmt.Printf("bName.txt -> Text file holding boy names for randomization, one name per line.\n")
	fmt.Printf("gName.txt -> Text file holding girl names for randomization, one name per line.\n")
	fmt.Printf("class.txt -> Text file holding class names for randomization, one class per line.\n\n")
}

func random() {
	defer wg.Done()
	// set random values based on files if files not present uses hard coded defaults
	var sex, name, class, classType string
	sex = randimizer.StringRandimizer([]string{"Male", "Female"})
	minion.SetSex(sex)

	// set random name based on sex
	if minion.Sex == "Female" {
		names, err := fileinputs.SliceFromTXTFile("gName.txt") // use gName file
		if err != nil {                                        // if file not present use defaults
			name = randimizer.StringRandimizer([]string{"Jane", "Mary", "Susan"})
		} else {
			name = randimizer.StringRandimizer(names)
		}
	} else {
		names, err := fileinputs.SliceFromTXTFile("bName.txt") // use bName file
		if err != nil {                                        // if file not present use defaults
			name = randimizer.StringRandimizer([]string{"John", "Gary", "Richard"})
		} else {
			name = randimizer.StringRandimizer(names)
		}
	}
	minion.SetName(name)

	// set random class
	classes, err := fileinputs.SliceFromTXTFile("class.txt")
	if err != nil {
		class = randimizer.StringRandimizer([]string{"Fighter", "Ranger", "Cleric"})
	} else {
		class = randimizer.StringRandimizer(classes)
	}
	minion.SetClass(class)

	// set random class type NOTE: need to work out a way to link this to classes
	classType = randimizer.StringRandimizer([]string{"Melee", "Ranged",
		"Divine", "IArcane", "CArcane", "Skill"})
	minion.SetDefaultStats(classType) // need error control here...

	minion.DisplayNpcBlock()
}

func custom() {
	var sex, name, class string
	// sex from user as argument else random
	if len(mArgs) >= 1 { // is sex present?
		sex = mArgs[0]
	} else {
		sex = randimizer.StringRandimizer([]string{"Male", "Female"})
	}
	minion.SetSex(sex)

	// name from user as argument else random based on sex
	if len(mArgs) >= 2 { // is name present?
		name = mArgs[1]
	} else {
		if minion.Sex == "Female" {
			name = randimizer.StringRandimizer([]string{"Jane", "Mary", "Susan"})
		} else {
			name = randimizer.StringRandimizer([]string{"John", "Gary", "Richard"})
		}
	}
	minion.SetName(name)

	// class from user as argument else random
	if len(mArgs) >= 3 { // is class present?
		class = mArgs[2]
	} else {
		class = randimizer.StringRandimizer([]string{"Fighter", "Ranger", "Cleric"})
	}
	minion.SetClass(class)

	// abilities from user as argument else random based on class type
	if len(mArgs) >= 9 { // are abilities present?
		var abilities []int
		for i := 3; i <= 8; i++ {
			var temp int
			temp, _ = strconv.Atoi(mArgs[i])
			abilities = append(abilities, temp)
		}
		minion.SetStats(abilities)
	} else { // set random stats
		classType := randimizer.StringRandimizer([]string{"Melee", "Ranged",
			"Divine", "IArcane", "CArcane", "Skill"})
		minion.SetClassType(classType)
		minion.SetDefaultStats(minion.ClassType)
	}

	minion.DisplayNpcBlock()
}
