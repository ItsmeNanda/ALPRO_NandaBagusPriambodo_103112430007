package main

import "fmt"

func main() {
	var huruf string

	fmt.Print("Masukan Huruf: ")
	fmt.Scan(&huruf)

	if huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" {
		fmt.Print("Ini adalah huruf vokal")
	} else {
		fmt.Print("Ini adalah huruf konsonan")
	}
}
