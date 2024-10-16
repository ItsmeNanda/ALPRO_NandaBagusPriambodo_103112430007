package main

import "fmt"

func main() {
	var n, m, rumus int

	fmt.Print("Masukan bilangan pertama: ")
	fmt.Scanln(&n)
	fmt.Print("Masukan bilangan terakhir: ")
	fmt.Scanln(&m)
	rumus = 1
	for i := 0; i < m; i++ {
		rumus = rumus * n
	}

	fmt.Print("hasilnya adalah: ", rumus)
}
