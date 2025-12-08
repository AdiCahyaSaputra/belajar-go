package main

import "fmt"

func main() {
	var a uint32 = 10
	var b uint64 = uint64(a)

	var nilai32 int32 = 32768
	var nilai16 int16 = int16(nilai32)

	fmt.Println(a, b)

	// tiba" nilai16 nya -32768 (number overflow)
	// bakal di convert ke yang paling rendah nya + (nilai lebihan - 1) si int16 yaitu -32768
	// klo nilai lebihan - 1 nya masih kelebihan juga ya bakal muter lagi ke yang paling rendah
	fmt.Println(nilai32, nilai16)

	var name = "Adi"
	var aFromName = name[0] // dalam byte / uint
	var aVersiString = string(aFromName)

	fmt.Println(a, aVersiString)
}
