package main

import "fmt"

func main() {
	name := "Adi"

	// Short statement if
	if length := len(name); length > 2 {
		fmt.Println("Nama terlalu panjang")
	}

	// Short statement switch
	switch length := len(name); length > 2 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}
}
