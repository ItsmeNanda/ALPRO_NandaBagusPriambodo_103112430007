package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Print("Masukan bilangan awal: ")
	fmt.Scan(&a)
	fmt.Print("Masukan bilangan akhir: ")
	fmt.Scan(&b)

	for i := a; i <= b; i++ {

		c = c + i

	}
	fmt.Print(c)
}
