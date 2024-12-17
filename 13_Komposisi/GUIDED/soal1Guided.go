package main

import "fmt"

func main() {
	var inp int

	fmt.Print("Masukkan angka positif: ")
	fmt.Scan(&inp)

	for i := 0; i <= inp; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
