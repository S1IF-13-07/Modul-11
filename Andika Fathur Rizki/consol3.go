package main

import "fmt"

func main() {
	var kendaraan string
	var durasi int
	var tarif int
	fmt.Print("Masukan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)
	switch {
	case kendaraan == "Motor" && durasi >= 1 && durasi <= 2:
		tarif = 7000
	case kendaraan == "Motor" && durasi > 2:
		tarif = 9000
	case kendaraan == "Mobil" && durasi >= 1 && durasi <= 2:
		tarif = 15000
	case kendaraan == "Mobil" && durasi < 2:
		tarif = 20000
	case kendaraan == "Truk" && durasi >= 1 && durasi <= 2:
		tarif = 250000
	case kendaraan == "Truk" && durasi < 2:
		tarif = 350000
	default:
		fmt.Println("Jenis kendataan atau durasi parkir tidak valid")
	}
	fmt.Printf("Tarif parkir: RP %d\n", tarif)
}
