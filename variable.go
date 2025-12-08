package main

import "fmt"

func main() {
	// kayak let di js
	// tipe data nya bisa auto detect kayak ts
	var name string

	name = "Adi"
	fmt.Println(name)

	name = "Cahya"
	fmt.Println(name)

	// var juga gak wajib, tapi harus langsung assign pake :=
	age := 20
	fmt.Println(age)

	// multiple var
	var (
		firstName = "Adi"
		lastName  = "Cahya"
	)

	fmt.Println(firstName, lastName)
}
