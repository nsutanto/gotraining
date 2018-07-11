// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

import "fmt"

// Add imports.

func main() {

	// Declare a nil slice of integers.
	var mySlices []int

	// Appends numbers to the slice.
	mySlices = append(mySlices, 1)
	mySlices = append(mySlices, 2)

	// Display each value in the slice.
	for _, mySlice := range mySlices {
		fmt.Println(mySlice)
	}

	// Declare a slice of strings and populate the slice with names.
	names := []string{"Bill", "Joan", "Jim", "Cathy", "Beth"}

	// Display each index position and slice value.
	for i, name := range names {
		fmt.Printf("Index: %d  Name: %s\n", i, name)
	}

	// Take a slice of index 1 and 2 of the slice of strings.
	slice := names[1:3]

	// Display each index position and slice values for the new slice.
	for i, name := range slice {
		fmt.Printf("Index: %d  Name: %s\n", i, name)
	}
}
