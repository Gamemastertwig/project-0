// Project-O (NPCGEN)
/*
	Project-O (NPCGEN) is a CLI application which will assist a player
	in creating a non-player role-playing characters for use in table
	top RPGs.

	Revature: Brandon Locker (GameMasterTwig)
*/
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
	"github.com/Gamemastertwig/project-0/webface"
)

var minion npc.Npc    // Npc
var r bool            // random flag
var m bool            // make flag
var w bool            // web/http flag
var h bool            // help flag
var mArgs []string    // non-flag arguments from make flag
var wg sync.WaitGroup // wait group for go routines

func init() {
	// sets flag options
	flag.BoolVar(&r, "random", false, "Generates a fully random NPC using set defaults") // random
	flag.BoolVar(&m, "make", false, "Generates a NPC using user defined variables")      // make
	flag.BoolVar(&w, "http", false, "Starts the HTTP web version")                       // http
	flag.BoolVar(&h, "help", false, "Display help menu and exits application")           // help1
	flag.BoolVar(&h, "h", false, "Display help menu and exits application")              // help2
	flag.Parse()
	if m || r || w {
		mArgs = flag.Args()
	}
}

// Main Function - This is where it all starts!
func main() {
	// -h, -help
	switch {
	case h: // -h, -help
		help()
		os.Exit(0)
	case r: // -random
		if len(mArgs) >= 1 {
			count, err := strconv.Atoi(mArgs[0])
			if err != nil {
				fmt.Println("Value after -random is not an integer!")
			} else {
				wg.Add(count) // add the nuber of go routines to wait group
				for i := 0; i < count; i++ {
					go random() // run all on seperate go routines
				}
			}
		} else {
			wg.Add(1)
			go random()
		}
	case m: // -make
		custom()
	case w: // -http
		web()
	default: // default operations
		userRequest()
	}
	//wait for any running go routines
	wg.Wait()
}

func web() {
	if len(mArgs) < 1 {
		fmt.Println("Please provide port, Example: -http :8080")
		os.Exit(0)
	} else if len(mArgs) > 1 {
		fmt.Println("Too many argments provided, Example: -http :8080")
		os.Exit(0)
	}
	//fmt.Println(mArgs)
	webface.StartHTTP(mArgs[0])
}

// userRequest (default) asked the user a series of questions to create a NPC to their liking
func userRequest() {
	a := true
	var str, dex, con, itl, wis, cha int
	var err error

	minion.SetGender(userinputs.RequestAnswer("What is your Gender?"))
	minion.SetName(userinputs.RequestAnswer("What is your Name?"))
	minion.SetClass(userinputs.RequestAnswer("What is Class?"))
	for a {
		str, err = strconv.Atoi(userinputs.RequestAnswer("What is your Strength?"))
		if err != nil {
			fmt.Printf("Invalid: non-numeric value provided for Strength\n")
		} else if str < 0 {
			fmt.Printf("Invalid: negative value provided for Strength\n")
		} else {
			a = false
		}
	}
	a = true
	for a {
		dex, err = strconv.Atoi(userinputs.RequestAnswer("What is your Dexterity?"))
		if err != nil {
			fmt.Printf("Invalid: non-numeric value provided for Dexterity\n")
		} else if dex < 0 {
			fmt.Printf("Invalid: negative value provided for Dexterity\n")
		} else {
			a = false
		}
	}
	a = true
	for a {
		con, err = strconv.Atoi(userinputs.RequestAnswer("What is your Consitution?"))
		if err != nil {
			fmt.Printf("Invalid: non-numeric value provided for Consitution\n")
		} else if con < 0 {
			fmt.Printf("Invalid: negative value provided for Consitution\n")
		} else {
			a = false
		}
	}
	a = true
	for a {
		itl, err = strconv.Atoi(userinputs.RequestAnswer("What is your Inteligence?"))
		if err != nil {
			fmt.Printf("Invalid: non-numeric value provided for Inteligence\n")
		} else if itl < 0 {
			fmt.Printf("Invalid: negative value provided for Inteligence\n")
		} else {
			a = false
		}
	}
	a = true
	for a {
		wis, err = strconv.Atoi(userinputs.RequestAnswer("What is your Wisdom?"))
		if err != nil {
			fmt.Printf("Invalid: non-numeric value provided for Wisdom\n")
		} else if wis < 0 {
			fmt.Printf("Invalid: negative value provided for Wisdom\n")
		} else {
			a = false
		}
	}
	a = true
	for a {
		cha, err = strconv.Atoi(userinputs.RequestAnswer("What is your Charisma?"))
		if err != nil {
			fmt.Printf("Invalid: non-numeric value provided for Charisma\n")
		} else if cha < 0 {
			fmt.Printf("Invalid: negative value provided for Charisma\n")
		} else {
			a = false
		}
	}
	minion.SetStats([]int{str, dex, con, itl, wis, cha})
	// display NPC
	minion.DisplayNpcBlock()
}

// help (-h, -help) displays a help menu to the user to help with use
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

// random (-random) randomly creates NPC based on files if files not present uses
// hard coded defaults
func random() {
	defer wg.Done()

	var gender, name, class, classType string
	gender = randimizer.StringRandimizer([]string{"Male", "Female"})
	minion.SetGender(gender)
	// set random name based on gender
	if minion.Gender == "Female" {
		names, err := fileinputs.ReadFile("gName.txt") // use gName file
		if err != nil {                                // if file not present use defaults
			name = randimizer.StringRandimizer([]string{"Jane", "Mary", "Susan"})
		} else {
			name = randimizer.StringRandimizer(fileinputs.SliceFromByte(names))
		}
	} else {
		names, err := fileinputs.ReadFile("bName.txt") // use bName file
		if err != nil {                                // if file not present use defaults
			name = randimizer.StringRandimizer([]string{"John", "Gary", "Richard"})
		} else {
			name = randimizer.StringRandimizer(fileinputs.SliceFromByte(names))
		}
	}
	minion.SetName(name)
	// set random class
	classes, err := fileinputs.ReadFile("class.txt")
	if err != nil {
		class = randimizer.StringRandimizer([]string{"Fighter", "Ranger", "Cleric"})
	} else {
		class = randimizer.StringRandimizer(fileinputs.SliceFromByte(classes))
	}
	minion.SetClass(class)
	// set random class type NOTE: need to work out a way to link this to classes
	classType = randimizer.StringRandimizer([]string{"Melee", "Ranged",
		"Divine", "IArcane", "CArcane", "Skill"})
	minion.SetDefaultStats(classType) // need error control here...
	// display NPC
	minion.DisplayNpcBlock()
}

// custom sets up a NPC based inputs passed to the application using arguments else
// else it will generate a randon value based on what is missing
func custom() {
	var gender, name, class string
	// gender from user as argument else random
	if len(mArgs) >= 1 { // is gender present?
		gender = mArgs[0]
	} else {
		gender = randimizer.StringRandimizer([]string{"Male", "Female"})
	}
	minion.SetGender(gender)
	// name from user as argument else random based on gender
	if len(mArgs) >= 2 { // is name present?
		name = mArgs[1]
	} else {
		if minion.Gender == "Female" {
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
	// display NPC
	minion.DisplayNpcBlock()
}
