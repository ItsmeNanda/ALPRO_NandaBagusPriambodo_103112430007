package main

import "fmt"

func main() {
	var a int

	fmt.Print("Masukkan jumlah motor: ")
	fmt.Scan(&a)

	motor := a / 2

	if a%2 == 0 {
		fmt.Print(motor)
	} else {

		fmt.Print(motor + 1)
	}

}
