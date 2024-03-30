package main

import (
	"fmt"
)

func main() {

	var map2 map[string]uint = make(map[string]uint, 1)
	map2 = map[string]uint{"adam": 12, "Peter": 1, "Raj": 23}
	fmt.Println(map2)

	var map1 map[string]uint = map[string]uint{"Ron": 12, "Ty": 12}
	fmt.Println(map1["Ron"])

	var age, ok = map1["Tiwar"]
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("Key not found", ok)
	}

	delete(map1, "Ron")
	fmt.Println(map1["Ron"])

	// Travesing the maps
	for k, v := range map2 {
		fmt.Println(k, v)
	}

}
