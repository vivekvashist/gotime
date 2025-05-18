package main

import (
	"fmt"
)

func main() {
	do()
}

func do() {
	defer fmt.Printf("First\n")
	defer fmt.Printf("Second\n")
	defer fmt.Printf("Third\n")
}

// program outputs defer output is backwards i.e last goes first
/*
Third
Second
First
*/
