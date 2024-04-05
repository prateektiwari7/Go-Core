package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Println("The value of p", *p)
	fmt.Println("The value of i", i)
	p = &i
	*p = 1
	fmt.Println("The New value of p", *p)
	fmt.Println("The New value of i", i)
}
