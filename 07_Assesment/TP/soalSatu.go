package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Masukkan jumlah bilangan untk pangkat: ")
	fmt.Scan(&a)

	for i := 1; i <= a; i++ {
		b = i * i
		fmt.Print(b)
	}
}
