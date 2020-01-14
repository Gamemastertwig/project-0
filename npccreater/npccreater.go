// Package npccreater is a Go Package for assiting in creating NPCs for RPGs
package npccreater

import (
	"fmt"
	"strings"
)

// Npc struct hold variable unique to ever NPC
type Npc struct {
	Sex       string
	Name      string
	Class     string
	Abilities []int
}

// Minion is a global variable that is a instance of Npc
var Minion Npc

// DefaultAbilities is a global variable that is a 2 diminsional (double) slice
// to hold the default ability values for certain class types per D20PFSRD ruleset
var DefaultAbilities [][]int

// FillDefaultAbil fill the DefaultAbilities variable with default values from the
// D20PFSRD ruleset
func FillDefaultAbil() {
	DefaultAbilities = append(DefaultAbilities, []int{13, 11, 12, 9, 10, 8}) // Melee
	DefaultAbilities = append(DefaultAbilities, []int{11, 13, 12, 10, 9, 8}) // Ranged
	DefaultAbilities = append(DefaultAbilities, []int{10, 8, 12, 9, 13, 11}) // Divine
	DefaultAbilities = append(DefaultAbilities, []int{8, 12, 10, 13, 9, 11}) // IArcane (INT)
	DefaultAbilities = append(DefaultAbilities, []int{8, 12, 10, 11, 9, 13}) // IArcane (CHA)
	DefaultAbilities = append(DefaultAbilities, []int{10, 12, 11, 13, 8, 9}) // Skill
}

// SetRandomName takes a filename as a string and returns a random string from
// the file. If file does not exist a error string will be returned. Names in
// the file should be ordered in as one name per line.
/* func SetRandomName(filename string) n string, err string {

} */

// SetSex sets Minion.Sex to the string passed to it
func SetSex(sex string) {
	Minion.Sex = sex
}

// SetName sets Minion.Name to the string passed to it
func SetName(name string) {
	Minion.Name = name
}

// SetClass sets Minion.Class to the string passed to it
func SetClass(class string) {
	Minion.Class = class
}

// SetDefaultStats sets Minion.Abilities to hold the default abilities based on
// class type passed to it. Must equal one of the following strings or a non-nil
// error (string) will be returned. Melee, Ranged, Divine, IArcane, CArcane, or Skill.
func SetDefaultStats(classType string) {
	// fill defaults first
	FillDefaultAbil()
	// assign abilities based on class type provided
	switch strings.ToLower(classType) {
	case "melee":
		Minion.Abilities = DefaultAbilities[0]
	case "ranged":
		Minion.Abilities = DefaultAbilities[1]
	case "divine":
		Minion.Abilities = DefaultAbilities[2]
	case "iarcane":
		Minion.Abilities = DefaultAbilities[3]
	case "carcane":
		Minion.Abilities = DefaultAbilities[4]
	case "skill":
		Minion.Abilities = DefaultAbilities[5]
	default:
		fmt.Println("SetDefaultStats: No valid class type provided")
	}
}
