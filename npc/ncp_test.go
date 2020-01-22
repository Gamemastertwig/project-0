package npc

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var n Npc

func TestFillDefaultAbil(t *testing.T) {
	testSlice := [][]int{
		{13, 11, 12, 9, 10, 8},
		{11, 13, 12, 10, 9, 8},
		{10, 8, 12, 9, 13, 11},
		{8, 12, 10, 13, 9, 11},
		{8, 12, 10, 11, 9, 13},
		{10, 12, 11, 13, 8, 9},
	}
	FillDefaultAbil()
	for i := range testSlice {
		for j := range testSlice[i] {
			if testSlice[i][j] != DefaultAbilities[i][j] {
				t.Errorf("FillDefaultAbil failed: Excpeted: %d", testSlice[i][j])
			}
		}
	}
}

func TestSetGender(t *testing.T) {
	gender := "Alien"
	n.SetGender(gender)
	if n.Gender != gender {
		t.Errorf("SetGender failed: Excpeted: %s", gender)
	}
}

func TestSetName(t *testing.T) {
	name := "Fred"
	n.SetName(name)
	if n.Name != name {
		t.Errorf("SetName failed: Excpeted: %s", name)
	}
}

func TestSetClass(t *testing.T) {
	class := "Teacher"
	n.SetClass(class)
	if n.Class != class {
		t.Errorf("SetClass failed: Excpeted: %s", class)
	}
}

func TestSetClassType(t *testing.T) {
	classType := "Skill"
	n.SetClassType(classType)
	if n.ClassType != classType {
		t.Errorf("SetClassType failed: Excpeted: %s", classType)
	}
}

func TestSetStats(t *testing.T) {
	stats := []int{10, 12, 11, 13, 8, 9}
	n.SetStats(stats)
	for i := range stats {
		if n.Abilities[i] != stats[i] {
			t.Errorf("SetStats failed: Excpeted: %d", stats[i])
		}
	}
}

func TestSetDefaultStats(t *testing.T) {
	classType := "Skill"
	var stats []int
	n.SetDefaultStats(classType)
	// once SetDefaultStats is called so should DefaultAbilities ([][]int)
	// since "Skill" is chosen it should match DefaultAbilities[5]
	for i := range DefaultAbilities[5] {
		stats = append(stats, DefaultAbilities[5][i])
	}
	// test n.abilities
	for i := range stats {
		if n.Abilities[i] != stats[i] {
			t.Errorf("SetDefaultStats failed: Excpeted: %d", stats[i])
		}
	}
}

func TestDisplayNpcBlock(t *testing.T) {
	// load NPC
	n.SetGender("Alien")
	n.SetName("Fred")
	n.SetClass("Teacher")
	n.SetClassType("Skill")
	n.SetStats([]int{10, 12, 11, 13, 8, 9})

	// back up Stdout
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	// create writer to fake Stdout
	read, out, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	// set writer to Stdout
	os.Stdout = out
	// run function to test, and capture Stdout
	n.DisplayNpcBlock()
	// need to close out here, otherwise ReadAll never gets "EOF".
	err = out.Close()
	if err != nil {
		log.Fatal(err)
	}
	// read from captured fake Stdout
	content, err := ioutil.ReadAll(read)
	if err != nil {
		log.Fatal(err)
	}
	// close reader
	err = read.Close()
	if err != nil {
		log.Fatal(err)
	}
	// exspected output
	testString := "NPC:\n Gender: Alien\n Name: Fred\n Class: Teacher\n" +
		" Abilities:\n  STR: 10 DEX: 12 CON: 11 INT: 13 WIS: 8 CHA: 9\n"
	// compair output to exspected
	if testString != string(content) {
		t.Errorf("DisplayNpcBlock failed: Excpeted: %s", testString)
	}
}
