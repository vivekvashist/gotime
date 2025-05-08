package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	text, err := os.ReadFile("main.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("file:%s\n", string(text))
}
