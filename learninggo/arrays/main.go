package main

import "fmt"

func main() {
	//	var a [5]int
	//	fmt.Printf("emp:%v\n", a)
	//
	//	a[4] = 200
	//	fmt.Printf("emp: %v\n", a)
	//	fmt.Printf("emp[2]: %v\n", a[1])
	//	a[1] = 100
	//	fmt.Printf("emp: %v\n", a)
	//	fmt.Printf("emp[2]: %v\n", a[1])
	//	// a[10] = 200 // index  out of bounds
	//	fmt.Printf("len a: %d\n", len(a))
	//	aa := [5]int{10, 20, 30, 40, 50}
	//	fmt.Printf("%v\n", aa)
	//
	//	var twoD [5][5]int
	//	for i := 0; i < 2; i++ {
	//		for j := 0; j < 3; j++ {
	//			twoD[i][j] = i + j
	//		}
	//	}
	//	fmt.Printf("2d: %v\n", twoD)

	//numbers := [4]int{10, 20, 30, 40} // literal construct
	//fmt.Printf("numbers: %v\n", numbers)
	//fmt.Printf("numbers: %T\n", numbers)

	var a [3]int
	a[0] = 2
	a[1] = 4
	a[2] = 6
	fmt.Printf("array: %v, %d, %d, %d\n", a, a[0], a[1], a[2])
	fmt.Printf("a: %T\n", a)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a: %d\n", a) // works!
	fmt.Printf("\n\n")

	var aa [4]int
	fmt.Printf("a: %v aa: %v\n", a, aa)
	// var aaa [4]int = a //./main.go:44:19: cannot use a (variable of type [3]int) as [4]int value in variable declaration
	fmt.Printf("\n\n")

}
