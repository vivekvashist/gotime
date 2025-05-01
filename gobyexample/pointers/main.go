// TODO: https://stackoverflow.com/questions/4938612/how-do-i-print-the-pointer-value-of-a-go-object-what-does-the-pointer-value-mea
// why use pointers?
// its efficient to store a variable once/one place and access it from multiple places
// allows you to share and update the variable from multiple places
// its more efficient than copying the variable multiple times
// and if you need to manipulate a value accross function calls you need pointers
// & -> address operator
// * -> indirection operator (dereferencing)
// *int -> a pointer type that represents a pointer
// nil -> is an untyped identifier that represents the lack of value of certain types
package main

import "fmt"

func main() {
	i := 42
	j := 2701
	fmt.Printf("i: %d\n", i)   // get the value of i
	fmt.Printf("j: %d\n", j)   // get the value of j
	fmt.Printf("&i: %p\n", &i) // address of i - read it as address of i
	fmt.Printf("&j: %p\n", &j) // address of j - read it as address of j

	// * has two roles
	// if * is in front of a type name e.g *int is a pointer type and int is the base
	// if * is in front of a variable which is a pointer type e.g *p the * acts as an operator
	// and returns what p is pointing to also know as dereferencing
	// value of p *int is the address of i
	// and *p is the value at i

	p := &i
	fmt.Printf("p:  %p\n", p)
	fmt.Printf("*p:  %d\n", *p)
	fmt.Printf("i: %d\n", i)
	*p = 21 // changing *p to 21 changes the value of i
	fmt.Printf("i: %d\n", i)
	fmt.Printf("\n")

	fmt.Printf("j: %d\n", j)
	p = &j
	*p = *p / 37 // modifying of value j
	fmt.Printf("j: %d\n", j)
}
