// Package fileinputs is a collections of functions for requesting
// information from files
package fileinputs

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// SliceFromTXTFile reads the text file with the filename (string) and
// returns a slice (string) holding all data from text file seperated by
// a newline ("\n")
func SliceFromTXTFile(filename string) []string {
	// read ALL bytes from file into data ([]byte)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
	}
	// combinds content of file into one string
	allDataString := string(data)
	// seperates string into slices (string) per line
	test := strings.Split(allDataString, "\n")
	// returns slices (string)
	return test
}
