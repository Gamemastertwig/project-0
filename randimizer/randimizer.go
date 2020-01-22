// Package randimizer is a collections of functions for randimizing groups of data
package randimizer

import (
	"math/rand"
	"time"
)

// StringRandimizer is function that takes a string slice and returns a
// random string from it.
func StringRandimizer(inputs []string) string {
	var l, n int
	// seed rand based on time
	rand.Seed(time.Now().UnixNano())
	// get a random int from 0 to length of string slice -1
	l = len(inputs)
	n = rand.Intn(l)
	// find string at index of random number
	// return that string
	return inputs[n]
}
