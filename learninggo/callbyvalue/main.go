package main

import (
	"fmt"
)

func doubleFail(a int, arr [2]int, s string) {
	a = a * 2
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
	s = s + s
	fmt.Printf("in doubleFail:%d, %v, %s\n", a, arr, s)
}

func main() {
	a := 1
	arr := [2]int{2, 4}
	s := "hello"
	doubleFail(a, arr, s)
	fmt.Printf("in main: %d, %v, %s\n", a, arr, s)
	fmt.Printf("\n\n")
}
