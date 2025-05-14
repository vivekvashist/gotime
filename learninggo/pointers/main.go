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
	//i := 42
	//j := 2701
	//fmt.Printf("i: %d\n", i)   // get the value of i
	//fmt.Printf("j: %d\n", j)   // get the value of j
	//fmt.Printf("&i: %p\n", &i) // address of i - read it as address of i
	//fmt.Printf("&j: %p\n", &j) // address of j - read it as address of j

	// * has two roles
	// if * is in front of a type name e.g *int is a pointer type and int is the base
	// if * is in front of a variable which is a pointer type e.g *p the * acts as an operator
	// and returns what p is pointing to also know as dereferencing
	// value of p *int is the address of i
	// and *p is the value at i

	//p := &i
	//fmt.Printf("p:  %p\n", p)
	//fmt.Printf("Type of p:  %T\n", p)
	//fmt.Printf("*p:  %d\n", *p)
	//fmt.Printf("i: %d\n", i)
	//*p = 21 // changing *p to 21 changes the value of i
	//fmt.Printf("i: %d\n", i)
	//fmt.Printf("\n")

	//fmt.Printf("j: %d\n", j)
	//p = &j
	//*p = *p / 37 // modifying of value j
	//fmt.Printf("j: %d\n", j)

	//age := 30
	//newAge := &age
	//fmt.Printf("newAge = %v\n", newAge)
	//fmt.Printf("&newAge = %p\n", &newAge)
	//fmt.Printf("*newAge = %v\n", *newAge)
	a := 10
	b := &a
	c := a
	fmt.Printf("a := 10, b := &a , c := a\n")
	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %p\n", b)
	fmt.Printf("*b: %d\n", *b)
	fmt.Printf("c: %d\n", c)
	fmt.Printf("\n")
	a = 20
	fmt.Printf("a := 20, b := &a , c := a\n")
	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %p\n", b)
	fmt.Printf("*b: %d\n", *b)
	fmt.Printf("c: %d\n", c)
	fmt.Printf("\n")
	*b = 30 // dereference the pointer to set the value
	fmt.Printf("a := 20, b := &a , *b = 30, c := a\n")
	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %p\n", b)
	fmt.Printf("*b: %d\n", *b)
	fmt.Printf("c: %d\n", c)
	fmt.Printf("\n")
	c = 40
	fmt.Printf("a := 20, b := &a , *b = 30, c := 40\n")
	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %p\n", b)
	fmt.Printf("*b: %d\n", *b)
	fmt.Printf("c: %d\n", c)
	fmt.Printf("\n")

	// another way to create a pointer
	// zero value of a pointer is nil
	// nil is not a memory location but the absence of one
	var bb *int // bb is type pointer to int
	fmt.Printf("%d\n", bb)
	//	fmt.Printf("%p\n", *bb) // nil dereference - panic: runtime error: invalid memory address or nil pointer dereference

	// use builtin new() func to create a pointerr
	bbb := new(int)
	fmt.Printf("%v\n", bbb) // 0xc000....
	fmt.Printf("%p\n", bbb) // 0xc000....
	fmt.Printf("%d\n", bbb) // 8242......
	fmt.Printf("%T\n", bbb) // *int

	fmt.Printf("%v\n", *bbb) // 0 i.e no panic

	// we can pass pointers to a function
	// go is call by value to assign values to paramters when a func is called
	// but for a pointer the value is the memory location where the value is stored
	// this gives an advantage to change the values in funcs that call them

	aaa := 20
	fmt.Printf("BEFORE: aaa: %d\n", aaa)
	fmt.Printf("&aaa %p\n", &aaa)
	setTo10(&aaa) // passing the pointer to aaa
	fmt.Printf("AFTER: aaa: %d\n", aaa)

	// go is still call by value trying to make aaa point to a different memory location and set to 10
	// that will not change the value in main
	// we are no longer refer to same memory location in both functions

	aaaa := 20
	fmt.Printf("aaaa: %d\n", aaaa) // aaaa = 20
	newsetTo10(&aaaa)
	fmt.Printf("aaaa: %d\n", aaaa) // aaaa = 20 still same

}
func setTo10(aaa *int) {
	*aaa = 10
}
func newsetTo10(aaaa *int) {
	aaaa = new(int) // assigning aaaa to brand new pointer i.e aaaa is a brand new pointer
	*aaaa = 10
}
