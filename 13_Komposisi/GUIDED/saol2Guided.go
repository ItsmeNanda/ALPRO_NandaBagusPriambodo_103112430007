package main

import (
	"fmt"
)

func main() {
	var n, max, min, number int

	fmt.Print("Masukkan jumlah bilangan yang ingin diinput: ")
	fmt.Scan(&n)

	fmt.Print("Masukkan bilangan 1: ")
	fmt.Scan(&number)
	max = number
	min = number

	for i := 2; i <= n; i++ {
		fmt.Printf("Masukkan bilangan %d: ", i)
		fmt.Scan(&number)

		if number > max {
			max = number
		}
		if number < min {
			min = number
		}
	}

	fmt.Printf("Bilangan terbesar: %d\n", max)
	fmt.Printf("Bilangan terkecil: %d\n", min)
}
