package main

import "fmt"

func main() {
	// Init Statement dan Post Statement
	for counter := 0; counter < 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	idx := 0

	// Tanpa Init dan Post statement
	for idx < 3 {
		fmt.Println("Perulangan ke", idx)
		idx++
	}

	// Slices string
	names := []string{"Adi", "Cahya", "Saputra"}

	// Pake versi normal
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// For Range kayak di python
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	// Gak butuh variable index
	for _, name := range names {
		fmt.Println(name)
	}
}
