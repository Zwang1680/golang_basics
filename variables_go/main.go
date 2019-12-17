package main

import "fmt"

var name = "stupid"

func main() {
	name = "Zion Wang"
	var age = 20
	const test = true
	dummy := "idiot"
	//automatically assigns as a float64
	size := 1.4
	var flt float64 = 1.65
	//asigning multiple variables
	name1, name2 := "brad brad", "charles charles"
	fmt.Println(name, age, dummy)
	fmt.Printf("%T\n", name)
	fmt.Println(size, flt, name1, name2)
}
