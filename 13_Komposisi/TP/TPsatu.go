package main

import "fmt"

func main() {
	var (
		angka int
		prima bool
	)

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	for i := 2; i <= angka; i++ {
		prima = true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				prima = false
				break
			}
		}
		if prima {
			fmt.Print(i, " ")
		}
	}
}
