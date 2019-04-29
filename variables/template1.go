// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

// Add imports
import "fmt"

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var kids int
	var name string

	// Display the value of those variables.
	fmt.Println("value is", name)
	fmt.Println("value is", kids)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	name = "PK"

	// Display the value of those variables.
	fmt.Println("Your name is ", name)

	// Perform a type conversion.
	x := 10
	y := int64(x)

	// Display the value of that variable.
	fmt.Println(y)
}
