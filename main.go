package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4}

	fmt.Println(reverseArray(array))
}

func reverseArray(array []int) []int {
	var newArray []int

	for i := len(array) - 1; i >= 0; i-- {
		newArray = append(newArray, array[i])
	}
	return newArray
}
