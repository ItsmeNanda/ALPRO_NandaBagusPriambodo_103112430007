package main

import "fmt"

func main() {
	var n int

	fmt.Print("masukkan bilangan: ")
	fmt.Scan(&n)

	hasil := 1

	for i := 2; i <= n; i++ {
		hasil = hasil * i
	}
	fmt.Println(hasil)
}
