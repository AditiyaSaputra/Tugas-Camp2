package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Tukar elemen jika elemen saat ini lebih kecil dari elemen berikutnya
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	// Array bilangan yang akan diurutkan
	arr := []int{64, 34, 25, 12, 22, 11, 17, 1, 90}

	// Panggil fungsi bubbleSort untuk mengurutkan dari yang terbesar ke yang terkecil
	bubbleSort(arr)

	// Cetak hasil pengurutan
	fmt.Println("Bilangan setelah diurutkan (dari terbesar ke terkecil):", arr)
}
