package main

import "fmt"

func main() {
	var n, m, o int

	fmt.Println("Berapa kali anda ingin menghitung: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("Masukan alas segitiga: ")
		fmt.Scan(&m)
		fmt.Print("Masukan tinggi segitiga: ")
		fmt.Scan(&o)
		p := m * o / 2
		fmt.Println("Hasilnya adalah:", p)
	}
}
