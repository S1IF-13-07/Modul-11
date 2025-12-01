package main

import "fmt"

func main() {
	var kendaraan string
	fmt.Print("Masukkan jenis kendaraan (mobil/motor/truk): ")
	fmt.Scan(&kendaraan)
	var durasi int
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	total := 0

	switch kendaraan {
	case "mobil":
		total = durasi * 5000
		fmt.Println("Rp ", total)
	case "motor":
		total = durasi * 2000
		fmt.Println("Rp ", total)
	case "truk":
		total = durasi * 8000
		fmt.Println("Rp", total)
	default:
		fmt.Println("Jenis kendaraan tidak dikenali")
	}
}