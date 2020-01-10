// NPCCreater is a Go Package for assiting in creating NPCs for RPGs
package NPCCreater

// RequestName takes a string (question), displays it to the user and returns
// the response in the form a string
func RequestName(question string) n string {
	// var n string
	fmt.Println(question) //output question
	fmt.Scanln(&n) //get answer
}

// RandomName takes a filename as a string and returns a random string from
// the file. If file does not exist a error string will be returned. Names in
// the file should be ordered in as one name per line.
func RandomName(filename string) n string err string {
  
}

// RequestClass takes a string (question) and an array of string (answers)
// displays both to user and compaires the answer to answers, it will reask 
// the question if a valid answer is not found. To be more flexable if every
// every answer in the answers array should start with a unique character.
func RequestClass(question string, answers []string) c string {
	var answer string
	a := false
	for !a {
		// display question
		fmt.Println(question)
		// display answers array
		for i int = 0; i < len(answers); i ++ {
			fmt.Println(answers[i])
		}
		fmt.Scan(&answer)

		// compare answer to answers  
	}
	/*
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
		}*/
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