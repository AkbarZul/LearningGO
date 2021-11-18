package main

import "fmt"

func main() {
	var name string

	name = "Akbar Zulfikar"
	fmt.Println(name)

	name = "Zulfikar"
	fmt.Println(name)

	var friendName = "Budi"
	fmt.Println(friendName)

	var age = 23
	fmt.Println(age)

	namaAsli := "Akbar dong"
	fmt.Println(namaAsli)

	// Multi Variable
	var (
		firstName = "Akbar"
		lastName = "Zulfikar"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}