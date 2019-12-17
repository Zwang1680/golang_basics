package main

import "fmt"

//func functionName(variables) (functionReturn can be multiple variables) {
func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func split(num1 int) (x, y int) {
	x = num1 * 2 / 4
	y = num1 - x
	return
}
func main() {
	fmt.Println(greeting("Zion"))
	fmt.Println(getSum(3, 4))
	fmt.Println(split(34))
}
