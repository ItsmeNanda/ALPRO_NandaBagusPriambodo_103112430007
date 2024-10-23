package main

import "fmt"

func main() {
	var a, dinar, dirham, fals, qirat int

	fmt.Print("Masukan uang qirat: ")
	fmt.Scan(&a)

	dinar = a / 600
	dirham = a % 600 / 60
	fals = (a % 600) % 60 / 6
	qirat = ((a % 600) % 60) % 6 / 1

	fmt.Println("Dalam bentuk dinar: ", dinar)
	fmt.Println("Dalam bentuk dirham: ", dirham)
	fmt.Println("Dalam bentuk fals: ", fals)
	fmt.Println("Dalam bentuk qirat: ", qirat)
}

// 1 dinar = 10 dirham
// 1 dirham = 10 fals
// 1 fals = 6 qirat
