package main

import (
	"fmt"
)

func main() {

	var f int16
	var rumus int16

	fmt.Println("Masukan sisi kubus")
	fmt.Scanln(&f)
	rumus = f * f * f
	fmt.Println("Jadi volume kubus nya adalah", rumus)
}
