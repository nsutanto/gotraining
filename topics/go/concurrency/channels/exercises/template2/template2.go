// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that uses a fan out pattern to generate 100 random numbers
// concurrently. Have each goroutine generate a single random number and return
// that number to the main goroutine over a buffered channel. Set the size of
// the buffered channel so no send ever blocks. Don't allocate more capacity
// than you need. Have the main goroutine store each random number it receives
// in a slice. Print the slice values then terminate the program.
package main

import (
	"time"
	"math/rand"
	"fmt"
)

// Add imports.

// Declare constant for number of goroutines.
const (
	goroutines = 100
)

func init() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Create the buffered channel with room for
	// each goroutine to be created.
	values := make(chan int, goroutines)

	// Iterate and launch each goroutine.
	for gr := 0; gr < goroutines; gr++ {

		// Create an anonymous function for each goroutine that
		// generates a random number and sends it on the channel.
		go func() {
			values <- rand.Intn(50)
		}()
	}

	// Create a variable to be used to track received messages.
	// Set the value to the number of goroutines created.
	wait := goroutines

	// Iterate receiving each value until they are all received.
	// Store them in a slice of ints.
	var numbers []int
	for wait > 0 {
		numbers = append(numbers, <-values)
		wait--
	}

	// Print the values in our slice.
	fmt.Println(numbers)
}
