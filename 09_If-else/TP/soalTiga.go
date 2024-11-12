package main

import "fmt"

func main() {
	var umur int
	var kewarganegaraan string

	fmt.Print("Masukan umur anda: ")
	fmt.Scan(&umur)
	fmt.Print("Masukan kewarga negaraan anda: ")
	fmt.Scan(&kewarganegaraan)

	if umur >= 17 && kewarganegaraan == "indonesia" {

		fmt.Print("Anda bisa mengikuti pemilu")

	} else if umur < 17 {

		fmt.Print("Umur anda di bawah 17")

	} else if kewarganegaraan != "indoensia" {

		fmt.Print("Kewarganegaraan anda bukan indonesia")

	}
}
