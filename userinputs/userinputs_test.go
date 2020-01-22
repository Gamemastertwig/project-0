package userinputs

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestRequestAnswer(t *testing.T) {
	content := []byte("Brandon")
	// create a temp file writer to store "fake" input from user
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	// defer closing and removing
	defer os.Remove(in.Name()) // clean up
	defer in.Close()
	// write "fake" input to file
	if _, err := in.Write(content); err != nil {
		log.Fatal(err)
	}
	// return to top of file
	if _, err := in.Seek(0, 0); err != nil {
		log.Fatal(err)
	}
	// create "fake" Stdin
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin
	// set Stdin to file writer
	os.Stdin = in
	if test := RequestAnswer("What is your name?"); test != string(content) {
		t.Errorf("TestRequestAnswer failed: Excpeted: %s", content)
	}
}

func TestMultiChoiceAnswer(t *testing.T) {
	content := []byte("2")
	// create a temp file writer to store "fake" input from user
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	// defer closing and removing
	defer os.Remove(in.Name()) // clean up
	defer in.Close()
	// write "fake" input to file
	if _, err := in.Write(content); err != nil {
		log.Fatal(err)
	}
	// return to top of file
	if _, err := in.Seek(0, 0); err != nil {
		log.Fatal(err)
	}
	// create "fake" Stdin
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin
	// set Stdin to file writer
	os.Stdin = in
	if test := MultiChoiceAnswer("What is your class?", []string{"Thief", "Fighter", "Mage"}); test != "Fighter" {
		t.Errorf("TestRequestAnswer failed: Excpeted: Fighter")
	}
}

func ExampleRequestAnswer() {
	RequestAnswer("What is your name?")
	// Output: What is your name?
}

func ExampleMultiChoiceAnswer() {
	MultiChoiceAnswer("What is your class?", []string{"Thief", "Fighter", "Mage"})
	// Output: What is your Class?
	// Output: 1 Thief
	// Output: 2 Fighter
	// Output: 3 Mage
}
