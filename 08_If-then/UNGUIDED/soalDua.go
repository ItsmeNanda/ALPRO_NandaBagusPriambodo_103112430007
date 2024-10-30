package main

import "fmt"

func main() {
	var a int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&a)

	if a%2 == 0 && a < 0 {
		fmt.Print("genap negatif")
	} else {

		fmt.Print("bukan")
	}

}
