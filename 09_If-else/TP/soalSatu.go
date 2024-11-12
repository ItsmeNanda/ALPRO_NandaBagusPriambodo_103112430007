package main

import "fmt"

func main() {
	var a int

	fmt.Print("Masukan nilai dari 1 - 100: ")
	fmt.Scan(&a)

	if a > 90 {
		fmt.Print("Nilai anda A")
	} else if a >= 80 && a <= 90 {
		fmt.Print("Nilai anda AB")
	} else if a >= 70 && a < 80 {
		fmt.Print("Nilai anda B")
	} else {
		fmt.Print("Nilai anda C")
	}
}
