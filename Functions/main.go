package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(printMe())
	divide, mod, err := ValueDivision(4, 0)
	if err == nil {
		fmt.Printf("Divide %v, mod %v", divide, mod)
	} else {
		fmt.Println(err)
	}

}

func printMe() string {
	return "printME"
}

func ValueDivision(x int, y int) (int, int, error) {
	var err error
	if y == 0 {
		err = errors.New("y cant be zero")
		return 0, 0, err
	}

	return x / y, x % y, nil
}
