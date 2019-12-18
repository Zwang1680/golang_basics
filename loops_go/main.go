package main

import "fmt"

func main() {
	//simple for loop, only have for loops in go
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	//another method
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	//basically a while loop using a for instead of a while
	//you can declare a for {} with no declarations but it is a infinite loop
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d \n", i)
	}

	//fizzbuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
