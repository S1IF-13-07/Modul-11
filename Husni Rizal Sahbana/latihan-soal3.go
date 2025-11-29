package main

import "fmt"

func main() {
	
	var kendaraan string
	var durasi_parkir int
	var tarif int

	fmt.Print("Jenis Kendaraan : ")
	fmt.Scan(&kendaraan)

	fmt.Print("Durasi Parkir : ")
	fmt.Scan(&durasi_parkir)

	switch {
	case kendaraan == "Motor" && durasi_parkir >= 1 && durasi_parkir <= 2:
		tarif = 7000
	case kendaraan == "Motor" && durasi_parkir > 2:
		tarif = 9000

	case kendaraan == "Mobil" && durasi_parkir >= 1 && durasi_parkir <= 2:
		tarif = 15000
	case kendaraan == "Mobil" && durasi_parkir > 2:
		tarif = 20000

	case kendaraan == "Truk" && durasi_parkir >= 1 && durasi_parkir <= 2:
		tarif = 25000
	case kendaraan == "Truk" && durasi_parkir > 2:
		tarif = 35000
	default:
		fmt.Println("Tidak Valid")
	}
	fmt.Printf("Tarif Parkir: Rp %d\n", tarif)
}
