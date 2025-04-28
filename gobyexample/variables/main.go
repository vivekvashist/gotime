package main

import (
	"fmt"
)

func main() {
	var a = "initial"
	fmt.Printf("%s\n", a)

	var b, c int = 1, 2
	fmt.Printf("%d, %d\n", b, c)

	var d = true
	fmt.Printf("%t\n", d)

	var e int
	fmt.Printf("%d\n", e)

	f := "apple" // := shorthand for declaring and initializing a variable e.g var f string = "apple", This syntax is only available inside functions
	fmt.Printf("%s\n", f)

}
