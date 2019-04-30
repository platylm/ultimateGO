// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

// Add imports.
import (
	"fmt"
)

func main() {

	// Declare and make a map of integer type values.
	number := make(map[string]int)

	// Initialize some data into the map.
	number["a"] = 1
	number["b"] = 2
	number["c"] = 3

	// Display each key/value pair.
	for key := range number {
		fmt.Println(key, number[key])
		// for key,value := range number{
		//fmt.Println(key, value)
		//}
	}
}
