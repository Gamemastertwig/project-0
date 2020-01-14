// Project-O
// Character Generator
package main

import (
	"fmt"

	"github.com/Gamemastertwig/project-0/npc"
)

var minion npc.Npc

func init() {
	minion.SetSex("Male")
	minion.SetName("Fred Smith")
	minion.SetClass("Fighter")
	//minion.SetDefaultStats("guru") // need error control here...
	minion.SetStats([]int{12, 12, 12, 12, 12, 12})
}

// Main Function - This is where it all starts!
func main() {
	fmt.Println(minion)
}

// SetRandomName takes a filename as a string and returns a random string from
// the file. If file does not exist a error string will be returned. Names in
// the file should be ordered in as one name per line.
/* func SetRandomName(filename string) n string, err string {

} */
