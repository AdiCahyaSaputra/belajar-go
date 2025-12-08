package main

import "fmt"

func main() {
	// tipe data nya gak bisa beda" :(
	// array tidak bisa bertambah kapasitas nya
	// mirip" java
	var numbersArray [4]int

	numbersArray[0] = 1
	numbersArray[1] = 2
	numbersArray[2] = 3
	numbersArray[3] = 4

	fmt.Println(numbersArray)

	// Idih mirip bet kek sebelah yak wkwkwk
	var values = [3]int{
		10,
		20,
		30,
	}

	fmt.Println(values)

	// Bisa auto detect length nya tapi pas di awal aja
	// gk bisa di tambah" lagi tetep
	var newValues = [...]int{
		10,
		20,
		30,
		40,
	}

	fmt.Println(newValues)
}
