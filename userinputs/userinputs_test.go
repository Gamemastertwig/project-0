package userinputs

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestRequestAnswer(t *testing.T) {
	content := []byte("Brandon")
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(in.Name()) // clean up
	defer in.Close()

	if _, err := in.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := in.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = in
	if test := RequestAnswer("What is your name?"); test != string(content) {
		t.Errorf("TestRequestAnswer failed: Excpeted: %s", content)
	}
}

func TestMultiChoiceAnswer(t *testing.T) {
	content := []byte("2")
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(in.Name()) // clean up
	defer in.Close()

	if _, err := in.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := in.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = in
	if test := MultiChoiceAnswer("What is your class?", []string{"Thief", "Fighter", "Mage"}); test != "Fighter" {
		t.Errorf("TestRequestAnswer failed: Excpeted: Fighter")
	}
}

func ExampleRequestAnswer() {
	RequestAnswer("What is your name?")
	// Output: What is your name?
	// Input: Brandon
	// Output: Brandon
}

func ExampleMultiChoiceAnswer() {
	MultiChoiceAnswer("What is your class?", []string{"Thief", "Fighter", "Mage"})
	// Output: What is your Class?
	// Output: 1 Thief
	// Output: 2 Fighter
	// Output: 3 Mage

}
