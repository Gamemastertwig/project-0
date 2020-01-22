package randimizer

import (
	"fmt"
	"testing"
)

func TestStringRandimizer(t *testing.T) {
	present := false
	testStringSlice := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	temp := StringRandimizer(testStringSlice)
	// test that return was in the original slice
	for i := range testStringSlice {
		if temp == testStringSlice[i] {
			present = true
			break
		}
	}
	// if not fail
	if !present {
		t.Errorf("TestStringRandimizer failed: unable to locate return string in provided slice")
	}
}

// output can be Morning, Evening, or Night
func ExampleStringRandimizer() {
	testStringSlice := []string{"Morning", "Evening", "Night"}
	temp := StringRandimizer(testStringSlice)
	fmt.Println(temp)
}
