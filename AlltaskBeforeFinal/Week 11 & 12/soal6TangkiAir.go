package main

import (
	"fmt"
)

func main() {
	var T, V int
	totalVolume := 0

	fmt.Print("Masukkan kapasitas tanki: ")
	fmt.Scan(&T)

	for {
		fmt.Print("Masukkan volume air dari ember: ")
		fmt.Scan(&V)

		totalVolume += V

		if totalVolume >= T {
			fmt.Println("true")
			break
		} else {
			fmt.Println("false")
		}
	}
}
