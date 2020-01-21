package randimizer

import (
	"fmt"
	"testing"
)

func TestStringRandimizer(t *testing.T) {
	present := false
	testStringSlice := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	temp := StringRandimizer(testStringSlice)
	for i := range testStringSlice {
		if temp == testStringSlice[i] {
			present = true
			break
		}
	}
	if !present {
		t.Errorf("TestStringRandimizer failed: unable to locate return string in provided slice")
	}
}

func ExampleStringRandimizer() {
	testStringSlice := []string{"Morning", "Evening", "Night"}
	temp := StringRandimizer(testStringSlice)
	fmt.Println(temp)
	// Output: Morning || Evening || Night
}
