package main

import (
	"fmt"
)

func main() {
	var kendaraan string
	var durasi, tarif int
	fmt.Scan(&kendaraan, &durasi)
	switch kendaraan {
	case "motor":
		tarif = 2000
	case "mobil":
		tarif = 5000
	case "truk":
		tarif = 8000
	default:
		fmt.Println("invalid kendaraan")
		return
	}
	if durasi < 1 {
		durasi = 1
	}
	biaya := tarif * durasi

	fmt.Printf("Rp %d\n", biaya)
}
