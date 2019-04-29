// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

import (
	"fmt"
)

// Add imports.

func main() {

	// Declare a nil slice of integers.
	var data []int

	// Append numbers to the slice.
	for i := 0; i < 10; i++ {
		data = append(data, i+10)
	}

	// Display each value in the slice.
	for _, datas := range data {
		fmt.Println("Value is ", datas)
	}
	fmt.Println()

	// Declare a slice of strings and populate the slice with names.
	animal := make([]string, 3)
	animal[0] = "Dog"
	animal[1] = "Cat"
	animal[2] = "Bird"

	//animal := []string{"Dog", "Cat", "Bird"}

	// Display each index position and slice value.
	for _, animals := range animal {
		fmt.Println(animals)
	}
	fmt.Println()

	// Take a slice of index 1 and 2 of the slice of strings.
	slice := animal[0:2]

	// Display each index position and slice values for the new slice.
	for i, name := range slice {
		fmt.Printf("Index: %d  Name: %s\n", i, name)
	}
}
