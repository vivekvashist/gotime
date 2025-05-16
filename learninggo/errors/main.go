package main

import (
	"fmt"
	"time"
)

type MyError struct {
	str  string
	when time.Time
}

func (e MyError) Error() string {
	return fmt.Sprintf("%s at %s\n", e.str, e.when)
}

func main() {
	err := MyError{}
	err.str = "Network is down"
	err.when = time.Now()
	fmt.Printf("%s\n", err)
}
