package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Printf("7 is even\n")
	} else {
		fmt.Printf("7 is odd\n")
	}

	if 8%4 == 0 {
		fmt.Printf("8 is divisible by 4\n")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Printf("either 8 or 7 are even\n")
	}

	num := 9
	if num < 0 {
		fmt.Printf("%d is negative\n", num)
	} else if num < 10 {
		fmt.Printf("%d has 1 digit\n", num)
	} else {
		fmt.Printf("%d has multiple digits\n", num)
	}
}
