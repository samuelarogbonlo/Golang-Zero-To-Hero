//Variables - Shorthand Method
//In Go, there are major types; int, string, bool, byte, float, complex.package main
//Shorthand method just displays other ways of variable declarations

package main

import "fmt"

func main() {

	name, email := "Anonymous is", "sbayo971@gmail.com" //shorthand
	size := 34.4

	var age = 24
	const anonprogrammer = true

	fmt.Println(name, age, anonprogrammer, email)
	fmt.Printf("%T\n", size) //it could be size, age, name or any other declaration

}
