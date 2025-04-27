package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Printf("%s\n", s)

	const n = 500000000
	const d = 3e20 / n
	fmt.Printf("%f\n", d)
	fmt.Printf("%d\n", int64(d))
	fmt.Printf("%f\n", math.Sin(n))
}
