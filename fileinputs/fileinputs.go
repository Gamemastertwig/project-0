// Package fileinputs is a collections of functions for requesting
// information from files
package fileinputs

import (
	"io/ioutil"
	"strings"
)

// SliceFromByte reads the data ([]byte) and returns a slice (string) holding
// all values from data ([]byte) that is seperated by a newline ("\n")
func SliceFromByte(data []byte) []string {
	// combinds content of data into one string
	allDataString := string(data)
	// seperates string into slices (string) per line
	test := strings.Split(allDataString, "\n")
	// returns slices (string)
	return test
}

// ReadFile is a wrapper for ioutil.ReadFile
func ReadFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	return data, err
}
