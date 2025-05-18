package main

import (
	"fmt"
)

var (
	even  int
	odd   int
	zeros int
	total int
)

func main() {
	numbers := []int{1, 3, 0, 0, 0, 4, 5, 13, 20, 18, 10, 11, 22, 100}
	for i, n := range numbers {
		fmt.Printf("index: %d number: %d\n", i+1, n)
		total += 1

		switch {
		case n == 0:
			zeros += 1
		case n%2 == 0:
			even += 1
		default:
			odd += 1
		}

	}
	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("Lenght: %d, Evens: %d , Odds: %d, Zeros: %d, Total: %d\n", len(numbers), even, odd, zeros, total)
	fmt.Printf("\n")
}
