package main

import "fmt"

func main() {
	var n, m int

	fmt.Println("Masukan bilangan pertama: ")
	fmt.Scanln(&n)
	fmt.Println("Masukan bilangan terakhir: ")
	fmt.Scanln(&m)
	for i := n; i <= m; i++ {
		fmt.Println(i)
	}
}
