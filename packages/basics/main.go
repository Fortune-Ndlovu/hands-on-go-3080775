// packages/basics/main.go
package main

// import the fmt package from the standard library
import (
	"fmt"
	"time"
)

// import the packages from the standard library
func main() {

	// use the fmt package to print the string "Hello Fortune"
	fmt.Println("Hello Fortune, you are more than welcome")

	// use the time package to print the current weekday
	fmt.Printf("Today is %s", time.Now().Weekday())
}
