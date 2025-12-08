package main

import "fmt"

func main() {
	const firstName string = "Adi"
	const lastName = "Cahya"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	// Multiple
	const (
		a = 1
		b = 2
		c = 3
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
