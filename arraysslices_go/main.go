package main

import "fmt"

func main() {
	//arrays in go must be a fixed length and you must name the type
	//slices are arrays without a fixed type
	var arr1 [2]string
	//assigning values, indexes go from 0 to max-1
	arr1[0] = "this"
	arr1[1] = "duh"
	// can print just the array and will print in order
	fmt.Println(arr1)
	//declaring and assigning at the same time
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	//assigning nothing to the array and attempting to print
	var arr3 [3]string
	fmt.Println(arr3)
	// outputs an actually empty string
	//slices kind of like arrays in c++
	arr4 := []string{"how", "do", "stuff"}
	arr4 = append(arr4, "stupid")
	arr4 = append(arr4, "duh", "banana")
	//len returns the length of an array or slice
	fmt.Println(len(arr4), arr4)
	//taking a specific array arr4[1:2] starts at 1 and stops at 2 returning an array arr4[1] just gives 1 and returns as a string not
	// as an array with 1 string
	fmt.Println(arr4[1:3], ", ", arr4[1])
}
