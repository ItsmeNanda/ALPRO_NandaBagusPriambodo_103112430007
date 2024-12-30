package main

import (
	"fmt"
)

func main() {
	var gol1, gol2, gol3, gol4 int

	fmt.Print("Masukkan jumlah gol tim (4 tim): ")
	fmt.Scan(&gol1, &gol2, &gol3, &gol4)

	maxGol := gol1
	if gol2 > maxGol {
		maxGol = gol2
	}
	if gol3 > maxGol {
		maxGol = gol3
	}
	if gol4 > maxGol {
		maxGol = gol4
	}

	minGol := gol1
	if gol2 < minGol {
		minGol = gol2
	}
	if gol3 < minGol {
		minGol = gol3
	}
	if gol4 < minGol {
		minGol = gol4
	}

	fmt.Print(maxGol, minGol)
}
