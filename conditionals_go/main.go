package main

import "fmt"

func main() {
	x := 7
	y := 10
	//no need for perenthesis
	if x <= y {
		//use %d as placeholders for the variables passed in later on this is only in Printf
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}
	color := "red"
	switch color {
	case "red":
		fmt.Println("color is red")
	case "green":
		fmt.Println("color is green")
	default:
		fmt.Println("color is not green or red")
	}
}
