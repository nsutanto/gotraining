// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// Add imports

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var age int
	var name string

	// Display the value of those variables.
	fmt.Printf("%T [%v]\n", age, age)
	fmt.Printf("%T [%v]\n", name, name)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	spouse := "Joan"

	// Display the value of those variables.
	fmt.Printf("%T [%v]\n", spouse, spouse)

	// Perform a type conversion.
	var y float32
	y = 3.14
	x := float64(y)

	// Display the value of that variable.
	fmt.Printf("%T [%v]\n", x, x)
}
