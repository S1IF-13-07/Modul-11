package main

import "fmt"

func main() {
	var kendaraan string
	var durasi, tarif int
	fmt.Scan(&kendaraan)
	fmt.Scan(&durasi)

	switch kendaraan {
	case "mobil", "Mobil":
		if durasi > 2 {
			tarif = 20000
		} else {
			tarif = 15000
		}
	case "motor", "Motor":
		if durasi > 2 {
			tarif = 9000
		} else {
			tarif = 7000
		}
	case "truk", "Truk":
		if durasi > 2 {
			tarif = 35000
		} else {
			tarif = 25000
		}
	default:
		fmt.Println("Invalid")
		tarif = 0
	}
	fmt.Println("Tarif = ", tarif)
}
