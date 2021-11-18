package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPAkbar NoKTP = "123445454353434"
	var marriedStatus Married = false
	fmt.Println(noKTPAkbar)
	fmt.Println(marriedStatus)
}