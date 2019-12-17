package main

import "fmt"

//func functionName(variables) functionReturn {
func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}
func main() {
	fmt.Println(greeting("Zion"))
	fmt.Println(getSum(3, 4))
}
