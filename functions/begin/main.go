// functions/begin/main.go
package main

// importing the fmt package
import (
	"fmt"
	"strconv"
)

// simple greet function
func greet() string {
	return "Hello!"
}

// greetWithName returns a greeting with the name
func greetWithName(name string) string {
	return "Hello, " + name + "!"
}

// greetWithNameAndAge returns a greeting with the name and age of the person
func greetWithNameAndAge(name string, age int) (greeting string) {
	greeting = "Hello, my name is " + name + " and I am " + strconv.Itoa(age) + "years old."
	return
}

// divide divides two numbers and returns the result
// if the second number is zero, it returns  error
//

func main() {
	// invoke greet function
	// fmt.Println(greet())

	// invoke greetWithName function
	// fmt.Println(greetWithName("Fortune"))

	// invoke divide function
	// fmt.Println(divide(10, 2))

	// invoke divide function with zero denominator to get an error
	// fmt.Println(divide(5, 0))
}
