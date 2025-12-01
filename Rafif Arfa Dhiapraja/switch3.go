package main

import "fmt"

func main() {
	var jeniskendaraan string
	var durasiparkir int
	fmt.Print("Masukkan jenis kendaraan (mobil/motor/truk): ")
	fmt.Scan(&jeniskendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasiparkir)

	switch  {
	case jeniskendaraan == "mobil" && durasiparkir <= 2:
		fmt.Println("tarif parkir: Rp 15000")
	case jeniskendaraan == "mobil" && durasiparkir > 2:
		fmt.Println("tarif parkir: Rp 20000")
	case jeniskendaraan == "motor" && durasiparkir <= 2:
		fmt.Println("tarif parkir: Rp 7000")
	case jeniskendaraan == "motor" && durasiparkir > 2:
		fmt.Println("tarif parkir: Rp 9000")
	case jeniskendaraan == "truk" && durasiparkir <= 2:
		fmt.Println("tarif parkir: Rp 25000")
	case jeniskendaraan == "truk" && durasiparkir > 2:
		fmt.Println("tarif parkir: Rp 35000")
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
		fmt.Println("tarif parkir: Rp 0") 
}

}