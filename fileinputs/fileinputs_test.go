package fileinputs

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestSliceFromTXTFile(t *testing.T) {
	// create temp file
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// fill it with content (weekdays)
	if _, err := file.WriteString("Monday\nTuesday\nWednesday\nThursday\nFriday"); err != nil {
		log.Fatal(err)
	}

	// close file
	file.Close()

	// test function
	temp, err := SliceFromTXTFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// verify function
	if temp == nil {
		t.Errorf("SliceFromTXTFile failed: Excpeted: [Monday Tuesday Wednesday Thursday Friday]")
	} else {
		temp2 := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
		for i := range temp {
			if temp2[i] != temp[i] {
				t.Errorf("SliceFromTXTFile failed: Excpeted: %s", temp2[i])
			}
		}
	}

	// remove temp file
	err = os.Remove("test.txt") // cleanup
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleSliceFromTXTFile() {
	// test.txt
	//  Monday
	//  Tuesday
	//  Wednesday
	//  Thursday
	//  Friday

	lines, err := SliceFromTXTFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)

	// Output: [Monday Tuesday Wednesday Thursday Friday]
}
