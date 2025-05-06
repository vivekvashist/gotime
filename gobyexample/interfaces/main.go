package main

import (
	"fmt"
	"reflect"
)

type Cat struct{}

func (c Cat) Say() string { return "meow" }

type Dog struct{}

func (d Dog) Say() string { return "woof" }

type Sayer interface {
	Say() string
}

func main() {
	c := Cat{}
	//	fmt.Printf("Cat says: %s\n", c.Say())
	d := Dog{}
	//	fmt.Printf("Dog says: %s\n", d.Say())

	animals := []Sayer{c, d}
	for _, a := range animals {
		fmt.Printf("%s says: %s\n", reflect.TypeOf(a).Name(), a.Say())
	}

	fmt.Printf("\n\n")
}
