package main

import "fmt"

/*bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

var name = "stupid"
*/

func main() {
	name := "Zion Wang"
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
	// trying not declaring anything results in a variable that actually holds absolutely nothing
	var name3 string
	fmt.Println(name3)
	//automatic for a bool is false
}
