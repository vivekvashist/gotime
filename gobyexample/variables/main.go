package main

import (
	"fmt"
)

func main() {
	var a int
	var s string
	var f float64
	var b bool
	fmt.Printf("var a int \t %[1]T %[1]v\n", a)
	fmt.Printf("var s int \t %T [%v]\n", s, s)
	fmt.Printf("var f int \t %T [%v]\n", f, f)
	fmt.Printf("var b int \t %T [%v]\n", b, b)
	fmt.Printf("\n\n")
}
