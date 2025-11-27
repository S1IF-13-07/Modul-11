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

	var tarifPerJam int

	switch jenis {
	case "motor":
		tarifPerJam = 2000
	case "mobil":
		tarifPerJam = 5000
	case "truk":
		tarifPerJam = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid.")
		return
	}

	total := tarifPerJam * durasi
	fmt.Printf("Total biaya parkir: Rp %d\n", total)
}
