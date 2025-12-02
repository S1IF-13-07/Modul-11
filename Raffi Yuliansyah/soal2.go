package main

import "fmt"

func main() {
	var kendaraan string
	var total, waktu int

	fmt.Scan(&kendaraan)
	fmt.Scan(&waktu)

	switch kendaraan {
	case "motor", "Motor":
		total = waktu * 2000

	case "mobil", "Mobil":
		total = waktu * 5000

	case "truk", "Truk":
		total = waktu * 8000

	default:
		fmt.Print("Invalid Input")
	}
	fmt.Print("Tarif Parkir ", waktu, " jam : Rp. ", total)
}