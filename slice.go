package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// slice is a reference to an array
	// slice punya 3 data penting -> pointer, length, capacity
	// pointer -> pointer ke data pertama di array yang mewakili slice
	// length -> panjang dari slice (bukan panjang array)
	// capacity -> kapasitas dari slice, jumlah maks dari pointer ke akhir array
	// [low:high] low -> index pertama atau index ke 4, high -> index terakhir - 1 jadi nya index ke 6
	// mirip python
	slice1 := months[4:7]
	slice2 := months[:7]
	slice3 := months[4:]
	slice4 := months[:]

	// pointer -> 4, length -> 3, capacity -> 8
	fmt.Println(slice1) // [Mei Juni Juli]
	fmt.Println(slice2) // [Januari Februari Maret April Mei Juni Juli]
	fmt.Println(slice3) // [Mei Juni Juli Agustus September Oktober November Desember]
	fmt.Println(slice4) // [Januari Februari Maret April Mei Juni Juli Agustus September Oktober November Desember]

	// fungsi yang kali aja berguna untuk slice
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3)) // capacity
	fmt.Println(cap(slice3)) // capacity

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	fmt.Println(daySlice1) // [Sabtu Minggu]

	daySlice1[0] = "Timpa sabtu" // mengubah data array days karena dia masih referensi itungannya
	fmt.Println(days)            // [Senin Selasa Rabu Kamis Jumat Timpa sabtu Minggu]

	// append slice, jika capacity sudah penuh maka akan membuat array baru
	daySlice2 := append(daySlice1, "Hari baru lagi")
	daySlice2[0] = "berubah" // tidak mengubah data array days karena dia sudah buat array baru
	fmt.Println(daySlice2)   // [berubah Minggu Hari baru lagi]
	fmt.Println(days)        // [Senin Selasa Rabu Kamis Jumat Timpa sabtu Minggu]

	newDays := [5]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat"}
	newDaysSlice := newDays[2:4]
	fmt.Println(newDaysSlice) // [Rabu Kamis]

	newDaysSliceAppend := append(newDaysSlice, "Hari baru") // tetep mengubah data array days karena capacity nya masih ada

	fmt.Println(newDaysSliceAppend) // [Rabu Kamis Hari baru]
	fmt.Println(newDays)            // [Senin Selasa Rabu Kamis Hari baru]

	// make slice, membuat slice baru dengan array kosong
	nameSlice := make([]string, 2, 5) // length 2, capacity 5
	nameSlice[0] = "Adi"
	nameSlice[1] = "Cahya"
	// nameSlice[2] = "Saputra" -> error karena length nya cuma 2

	// selama masih ada capacity maka tidak akan membuat array baru
	nameSlice = append(nameSlice, "Saputra")
	fmt.Println(nameSlice) // [Adi Cahya Saputra]

	// copy slice
	// dia bikin slice baru yang isi nya kosong namun length dan capacity nya sama dengan nameSlice
	copySlice := make([]string, len(nameSlice), cap(nameSlice))
	copy(copySlice, nameSlice)

	copySlice[0] = "Cahya" // tidak mengubah data array nameSlice karena dia sudah buat array baru

	fmt.Println(copySlice)
}
