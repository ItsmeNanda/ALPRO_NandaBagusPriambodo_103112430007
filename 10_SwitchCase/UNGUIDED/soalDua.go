package main

import (
	"fmt"
)

func main() {
	var kategori, jam int

	fmt.Println("1. Motor", "\n2. Mobil", "\n3. Truk")
	fmt.Print("Masukan kategori kendaraan: ")
	fmt.Scan(&kategori)
	fmt.Print("Masukan durasi parkir: ")
	fmt.Scan(&jam)

	switch kategori {
	case 1:
		harga := jam * 2000
		fmt.Println("Total harga parkir motor sebesar: ", harga)
	case 2:
		harga := jam * 5000
		fmt.Println("Total harga parkir mobil sebesar: ", harga)
	case 3:
		harga := jam * 8000
		fmt.Println("Total harga parkir Truk sebesar: ", harga)
	default:
		fmt.Println("Kode kategori tidak valid")
	}
}
