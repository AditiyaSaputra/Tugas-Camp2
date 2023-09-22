package main

import "fmt"

func main() {
	// Membuat slice kosong dengan panjang awal 5
	slice := make([]int, 5)

	// Menambahkan 3 data ke dalam slice
	slice[0] = 10
	slice[1] = 20
	slice[2] = 30
	slice[3] = 40
	slice[4] = 50

	// Cetak isi slice
	fmt.Println("Isi awal slice:", slice)

	// Menambahkan data ke dalam slice
	slice = append(slice, 60, 70, 80)

	// Cetak isi slice setelah penambahan data
	fmt.Println("Isi slice setelah penambahan data:", slice)
}