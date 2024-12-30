package main

import (
	"fmt"
)

func main() {
	var h1, m1, h2, m2 int

	fmt.Print("Masukkan jam dan menit mulai parkir: ")
	fmt.Scan(&h1, &m1)
	fmt.Print("Masukkan jam dan menit selesai parkir: ")
	fmt.Scan(&h2, &m2)

	mulai := h1*60 + m1
	selesai := h2*60 + m2

	durasi := selesai - mulai

	if durasi < 0 {
		durasi += 12 * 60
	}

	jam := durasi / 60
	menit := durasi % 60

	fmt.Printf("%d jam %d menit\n", jam, menit)
}
