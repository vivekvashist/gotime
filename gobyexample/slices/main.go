package main

import (
	"fmt"
)

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	fmt.Printf("%s\n", expressions)
	fmt.Printf("%v\n", expressions)
	fmt.Printf("%T\n", expressions)
	fmt.Printf("%#v\n", expressions)
	fmt.Printf("\n")

	for i, expression := range expressions {
		fmt.Printf("%d: %v\n", i+1, expression)
	}

}
