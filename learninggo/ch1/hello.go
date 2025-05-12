package main // package declaration

import "fmt" // packages referenced by this file i.e fumpt package from stdlib

func main() { // all go programs start from the main function in the main package
	fmt.Printf("Hello, World!\n")
}

// go compiler -> go build
// go code formatter -> go fmt
// go dependency manager -> go mod
// scan for common coding mistakes -> go vet
// go mod init -> markt a dir as Go module
// go mod init helloworld
// go project is called a module
// a module is not just source code - it is also exact specification of the dependencies
// of the code withing the module
// every module has a go.mod file in its root dir
// go mod init create go.mod file
// go mod file declares the name of the module, min supported version (similar to pythons requirements.txt)
// don't edit go.mod file directly, use go get and then go mod tidy
// within a go module code is organized into one or more packages
// main package in a Go module contains the code that starts the go program
// go imports only whole packages
