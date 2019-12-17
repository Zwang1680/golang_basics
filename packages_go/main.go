package main

import (
	"fmt"
	"math"

	"github.com/zwang1680/golang_basics/packages_go/strutil"
)

func main() {
	var num = 1.7
	fmt.Println("Math Floor Method on ", num, ": ", math.Floor(num))
	fmt.Println("Math Ceil Method on ", num, ": ", math.Ceil(num))
	num = 16
	fmt.Println("Math Sqrt Method on ", num, ": ", math.Sqrt(num))
	fmt.Println(strutil.Reverse("thats pretty crazy"))
}
