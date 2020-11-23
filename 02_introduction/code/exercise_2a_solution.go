// every go file needs a package name (at least main)
package main

//
import "fmt"
// Printing: 
// 1) fmt.Print() prints output to console
// 2) fmt.Println() the same as above but adds line after
// 3) fmt.Printf() for interpolating variables (%s for string, %t for boolean) like in Part Two 


// // Note: only uncomment one main() function at a time when testing 
//go run exercise_2a_solution.go that executes main function
//You an have only one func main.

// // Part One
// func main() {
// 	fmt.Print("Hello World")
// }

// // Part Two
func main() {
	fmt.Printf("Hi! my name is %s. I have lived in %s for %d. They say the weather is amazing, which is %t.\n",
  "Kasia", "Krakow", 5, true)
}
