//Variables
//In Go, there are major types; int, string, bool, byte, float, complex.package main

package main

import "fmt"

func main() {
	var name = "Anonymous is"
	var age = 24
	const anonprogrammer = true

	fmt.Println(name, age, anonprogrammer)
	fmt.Printf("%T\n", anonprogrammer)
}
