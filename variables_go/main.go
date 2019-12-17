package main

import "fmt"

var name = "stupid"

func main() {
	name = "Zion Wang"
	var age = 20
	const test = true
	dummy := "idiot"
	fmt.Println(name, age, dummy)
	fmt.Printf("%T\n", name)
}
