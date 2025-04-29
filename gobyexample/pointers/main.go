// https://stackoverflow.com/questions/4938612/how-do-i-print-the-pointer-value-of-a-go-object-what-does-the-pointer-value-mea
package main

import (
	"fmt"
)

func main() {
	i := int(42)
	fmt.Printf("1. main -- i %T: &i=%p i=%v\n", i, &i, i)
	p := &i
	fmt.Printf("2. main -- p %T: &p=%p p=&i=%p *p=i=%v\n", p, &p, p, *p)

}
