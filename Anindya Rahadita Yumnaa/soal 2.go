package main

import "fmt"

func main() {
	var jenis string
	var durasi int

	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&jenis)

	fmt.Print("Masukkan durasi parkir (jam): ")
	fmt.Scan(&durasi)

	if durasi < 1 {
		durasi = 1
	}

	var tarif int

	if jenis == "motor" {
		tarif = 2000
	} else if jenis == "mobil" {
		tarif = 5000
	} else if jenis == "truk" {
		tarif = 8000
	} else {
		fmt.Println("Jenis kendaraan tidak valid.")
		return
	}

	total := tarif * durasi
	fmt.Printf("Total biaya parkir: Rp %d\n", total)
}