package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 4
	fmt.Println(arr[0:2])
	fmt.Println(&arr[1])

	var arr1 [3]int = [3]int{1, 3, 4}
	fmt.Println(arr1)

	arr2 := [3]int{5, 67, 8}
	fmt.Println(arr2)

	var arr3 []int = []int{2, 41, 3, 4}
	arr3 = append(arr3, 9)
	fmt.Println(arr3)

	var arr4 []int = []int{6, 7, 89, 000, 12}
	arr3 = append(arr3, arr4...)
	fmt.Println(arr3)

}
