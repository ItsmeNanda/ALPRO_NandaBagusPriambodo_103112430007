package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Masukkan jumlah barang yang di beli: ")
	fmt.Scan(&a)

	for i := 1; i <= a; i++ {
		if i <= 5 {
			b += 10
		} else {
			b += 15
		}
	}

	fmt.Print("Total poin Yang di Dapat: ", b)
}
