package main

import "fmt"

func main() {
	var a, b int
	var faktorA, faktorB bool

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&b)

	if b%a == 0 {
		faktorA = true
	}
	if a%b == 0 {
		faktorB = true
	}

	fmt.Println(faktorA)
	fmt.Println(faktorB)

}
