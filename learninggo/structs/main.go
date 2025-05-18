package main

import (
	"fmt"
)

type User struct {
	Uid      string
	Gid      string
	Username string
	Name     string
	HomeDir  string
}

func main() {
	u := User{
		Uid:      "1000",
		Gid:      "1000",
		Username: "test",
		Name:     "Test Testing",
		HomeDir:  "/home/testt",
	}
	fmt.Printf("User: %v\n", u)  //User: {1000 1000 test Test Testing /home/testt}
	fmt.Printf("User: %+v\n", u) // User: {Uid:1000 Gid:1000 Username:test Name:Test Testing HomeDir:/home/testt}
	fmt.Printf("User: %#v\n", u) // User: main.User{Uid:"1000", Gid:"1000", Username:"test", Name:"Test Testing", HomeDir:"/home/testt"}
	fmt.Printf("\n")

}
