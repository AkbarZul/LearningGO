package main

import "fmt"

func main() {
	const firstName string = "Akbar"
	const lastName = "Zulfikar"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	// multi variable
	const(
		name = "Zulfikar"
		age = 23
	)
	fmt.Println(name)
	fmt.Println(age)
}