package main

import "fmt"

func main() {
	var a [5]int
	fmt.Printf("emp:%v\n", a)

	a[4] = 200
	fmt.Printf("emp: %v\n", a)
	fmt.Printf("emp[2]: %v\n", a[1])
	a[1] = 100
	fmt.Printf("emp: %v\n", a)
	fmt.Printf("emp[2]: %v\n", a[1])
	// a[10] = 200 // index  out of bounds
	fmt.Printf("len a: %d\n", len(a))
	aa := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("%v\n", aa)

	var twoD [5][5]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Printf("2d: %v\n", twoD)

}
