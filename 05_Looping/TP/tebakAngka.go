package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	angkaRahasia := rand.Intn(100) + 1

	fmt.Println("Tebak angka antara 1 sampai 100")

	var tebakan int
	var kesempatan int = 5

	for kesempatan > 0 {
		fmt.Printf("Kesempatan ke-%d: ", kesempatan)
		fmt.Scanln(&tebakan)

		if tebakan == angkaRahasia {
			fmt.Println("Selamat, Anda berhasil menebak!")
			break
		} else if tebakan < angkaRahasia {
			fmt.Println("Terlalu kecil!")
		} else {
			fmt.Println("Terlalu besar!")
		}

		kesempatan--
	}

	if kesempatan == 0 {
		fmt.Printf("Kesempatan habis. Angka rahasianya adalah: %d\n", angkaRahasia)
	}
}
