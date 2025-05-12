package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 3 {
		fmt.Printf("%d\n", i)
		i += 1
	}

	fmt.Printf("\n")

	for j := 0; j <= 3; j++ {
		fmt.Printf("%d\n", j)
	}

	fmt.Printf("\n")

	for i := range 3 {
		fmt.Printf("%d\n", i)
	}

	fmt.Printf("\n")

	for {
		fmt.Printf("loop\n")
		break
	}

	fmt.Printf("\n")

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Printf("%d\n", n)
	}

	fmt.Printf("\n")

	for i := range 5 {
		i++
		fmt.Printf("pinging 10.1.1.%d\n", i)
	}
	fmt.Printf("\n")

	for n := range 10 {
		if n == 6 {
			break
		}
		fmt.Printf("%d\n", n)
	}

}
