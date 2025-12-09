package main

import "fmt"

func main() {
	// map tidak ada batasan key
	person := map[string]string{
		"name":    "Adics",
		"address": "Indonesia",
	}

	fmt.Println(person["name"])
	fmt.Println(person["gaada"]) // empty string tergantung tipe data di awal

	person["age"] = "20"
	fmt.Println(person["age"])

	delete(person, "age") // menghapus key

	fmt.Println(person)
}
