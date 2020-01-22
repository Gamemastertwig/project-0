package fileinputs

import (
	"fmt"
	"testing"
)

func TestSliceFromByte(t *testing.T) {
	// create tempByte file
	tempByte := []byte("Monday\nTuesday\nWednesday\nThursday\nFriday")

	// test function
	temp := SliceFromByte(tempByte)

	// verify function
	if temp == nil {
		t.Errorf("SliceFromByte failed: Excpeted: [Monday Tuesday Wednesday Thursday Friday]")
	} else {
		temp2 := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
		for i := range temp {
			if temp2[i] != temp[i] {
				t.Errorf("SliceFromByte failed: Excpeted: %s", temp2[i])
			}
		}
	}
}

func ExampleSliceFromByte() {
	tempByte := []byte("Monday\nTuesday\nWednesday\nThursday\nFriday")
	lines := SliceFromByte(tempByte)

	fmt.Println(lines)
	// Output: [Monday Tuesday Wednesday Thursday Friday]
}
