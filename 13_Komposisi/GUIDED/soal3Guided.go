package main

import "fmt"

func main() {
	var inp int

	fmt.Print("Masukan bilangan bulat positif: ")
	fmt.Scan(&inp)

	for i := 1; i <= inp; i++ {
		if inp%i == 0 {
			fmt.Print(i, " ")
		}
	}
}
