package main

import "fmt"

func main() {

	var angka, jumlah int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	for i := 1; i < angka; i++ {
		if angka%i == 0 {
			jumlah += i
		}
	}

	if jumlah == angka {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}

}
