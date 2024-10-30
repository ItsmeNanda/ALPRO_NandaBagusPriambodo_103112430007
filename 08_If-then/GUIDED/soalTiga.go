package main

import "fmt"

func main() {
	var a int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&a)

	if a < 0 && a%2 == 0 {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}

}
