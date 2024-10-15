package main

import "fmt"

func main() {
	var n int

	fmt.Println("Masukan bilangan yang ingin di loop: ")
	fmt.Scanln(&n)
	/// ini adalah sebuah loop dimana kita memanggil for i yang akan di mulai dari angka 1
	for i := 1; i <= n; i++ { // disini terdapat n yaitu max loop yang di masukkan oleh user
		fmt.Println(i)
	}
}
