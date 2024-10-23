package main

import "fmt"

func main() {
	var a, c int

	fmt.Print("Masukan bilangan awal pangkat: ")
	fmt.Scan(&a)

	for i := 0; i <= a; i++ {

		c = i * i
		fmt.Print(c)

	}
}
