package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name        string
	Age         int
	Active      bool
	lastLoginAt string
}

func main() {
	data := User{Name: "Prateek", Age: 12, Active: true, lastLoginAt: "22"}
	data1, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data1))

}

func history() {
	u, err := json.Marshal(User{Name: "Bob", Age: 10, Active: true, lastLoginAt: "today"})

	if err != nil {
		panic(err)
	}
	fmt.Println(string(u))
}
