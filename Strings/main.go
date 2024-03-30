package main

import "fmt"

func main() {
	var myst = []rune("Prateek")
	fmt.Println(myst[0])

	for i, _ := range myst {
		fmt.Printf("index=%v and value=%v", i, myst[i])
	}

}
