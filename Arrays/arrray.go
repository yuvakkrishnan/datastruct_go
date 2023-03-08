package main

import "fmt"

func main() {
	// initialize an array
	arr := [5]int{1, 2, 3, 4, 5}

	// retrieve the element at the given index
	fmt.Println(arr[2]) // 3

	// update the element at the given index with the new value
	arr[3] = 7
	fmt.Println(arr) // [1 2 3 7 5]

	// get the length of the array
	fmt.Println(len(arr)) // 5

	// create a new array and copy elements from the source array
	dest := [5]int{}
	copy(dest[:], arr[:])
	fmt.Println(dest) // [1 2 3 7 5]

	// create a new array containing elements from the start index up to, but not including, the end index
	slice1 := arr[1:4]
	fmt.Println(slice1) // [2 3 7]

	// create a new array containing elements from the start index to the end of the array
	slice2 := arr[2:]
	fmt.Println(slice2) // [3 7 5]

	// create a new array containing elements from the beginning of the array up to, but not including, the end index
	slice3 := arr[:3]
	fmt.Println(slice3) // [1 2 3]
}
