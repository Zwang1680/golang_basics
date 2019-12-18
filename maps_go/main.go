package main

import "fmt"

func main() {
	//definition
	emails := make(map[string]string)

	emails["bob"] = "bob@gmail.com"
	emails["rick"] = "rick@gmail.com"
	emails["tom"] = "tom@gmail.com"
	emails["sam"] = "sam@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["bob"])
	//can use len(emails) to get the length
	delete(emails, "bob")
	fmt.Println(emails)

	studentId := map[int]string{23: "bob", 44: "sam"}
	fmt.Println(studentId)
}
