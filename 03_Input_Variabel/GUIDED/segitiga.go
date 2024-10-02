package main

import (
	"fmt"
)

func main() {

	var (
		f     int16
		g     int16
		rumus int16
	)

	fmt.Println("Masukan alas segitiga")
	fmt.Scanln(&f)
	fmt.Println("Masukan tinggi segitiga")
	fmt.Scanln(&g)
	rumus = (f / 2) * g
	fmt.Println("Jadi volume kubus nya adalah", rumus)
}
