package main

import "fmt"

func main() {
	type NoKTP string
	type Umur int

	var ktpAdi NoKTP = "1234567890"
	var umurAdi Umur = 20

	fmt.Println(ktpAdi)
	fmt.Println(umurAdi)
}
