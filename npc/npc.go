// Package npc is a Go Package the sudo-emulates a NPC classType, allows creations of NPCs
// and includes methods for setting there values
package npc

import (
	"fmt"
	"strings"
)

// Npc struct hold variable unique to ever NPC
type Npc struct {
	Gender    string
	Name      string
	Class     string
	ClassType string
	Abilities []int
}

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

// SetGender sets Minion.Gender to the string passed to it
func (m *Npc) SetGender(gender string) {
	m.Gender = gender
}

// SetName sets Minion.Name to the string passed to it
func (m *Npc) SetName(name string) {
	m.Name = name
}

// SetClass sets Minion.Class to the string passed to it
func (m *Npc) SetClass(class string) {
	m.Class = class
}

// SetClassType sets Minion.ClassType to the string passed to it
func (m *Npc) SetClassType(classType string) {
	m.ClassType = classType
}

// SetStats sets Minion.Abilities to the slice of intigers passed to it.
// NOTE: most RPGs have 6 stats...
// (Strength (STR), Dexterity (DEX), Constitution (CON), Intelligence (INT),
// Wisdom (WIS), and Charisma (CHA)). The slice allows for additional stats
// to be added based on play style.
func (m *Npc) SetStats(stats []int) {
	m.Abilities = stats
}

// SetDefaultStats is a member function of the NPC struct that sets Abilities to
// hold the default abilities based on class type passed to it.
// It calls FillDefaultAbil() and the string passed to it must equal one of the
// following strings. Melee, Ranged, Divine, IArcane, CArcane, or Skill.
func (m *Npc) SetDefaultStats(classType string) {
	// fill defaults first
	FillDefaultAbil()
	// assign abilities based on class type provided
	switch strings.ToLower(classType) {
	case "melee":
		m.Abilities = DefaultAbilities[0]
	case "ranged":
		m.Abilities = DefaultAbilities[1]
	case "divine":
		m.Abilities = DefaultAbilities[2]
	case "iarcane":
		m.Abilities = DefaultAbilities[3]
	case "carcane":
		m.Abilities = DefaultAbilities[4]
	case "skill":
		m.Abilities = DefaultAbilities[5]
	default:
		m.Abilities = []int{10, 10, 10, 10, 10, 10}
	}
}

// DisplayNpcBlock writes the NPC to stdout in a readable format
func (m *Npc) DisplayNpcBlock() {
	fmt.Printf("NPC:\n Gender: %s\n Name: %s\n", m.Gender, m.Name)
	fmt.Printf(" Class: %s\n", m.Class)
	if len(m.Abilities) == 6 {
		fmt.Printf(" Abilities:\n  STR: %d DEX: %d CON: %d INT: %d WIS: %d CHA: %d\n",
			m.Abilities[0], m.Abilities[1], m.Abilities[2], m.Abilities[3],
			m.Abilities[4], m.Abilities[5])
	}
}
