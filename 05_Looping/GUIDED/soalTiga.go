package main

import "fmt"

func main() {
	var n, m, rumus int

	fmt.Println("Masukan bilangan pertama: ")
	fmt.Scanln(&n)
	fmt.Println("Masukan bilangan terakhir: ")
	fmt.Scanln(&m)
	for i := 1; i <= m; i++ {
		rumus = rumus + n
	}
	fmt.Print("Hasilnya adalah: ", rumus)
}
