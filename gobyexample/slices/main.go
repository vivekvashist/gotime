package main

import (
	"fmt"
)

func main() {
	//	expressions := [][]string{
	//		{"2", "+", "3"},
	//		{"2", "-", "3"},
	//		{"2", "*", "3"},
	//		{"2", "/", "3"},
	//		{"2", "%", "3"},
	//		{"two", "+", "three"},
	//		{"5"},
	//	}
	//	fmt.Printf("%s\n", expressions)
	//	fmt.Printf("%v\n", expressions)
	//	fmt.Printf("%T\n", expressions)
	//	fmt.Printf("%#v\n", expressions)
	//	fmt.Printf("\n")
	//
	//	for i, expression := range expressions {
	//		fmt.Printf("%d: %v\n", i+1, expression)

	slice1 := make([]string, 5, 8)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"
	fmt.Printf("slice1: %v len: %d cap: %d\n", slice1, len(slice1), cap(slice1))
	slice2 := slice1[2:4]
	fmt.Printf("slice2: %v len: %d cap: %d\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("\n")
}
