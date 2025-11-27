package main

import "fmt"
func main() {
	var kendaraan string
	var lamaParkir, biaya int	
	fmt.Print("Masukkan jenis kendaraan (mobil/motor/truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan lama parkir (dalam jam): ")
	fmt.Scan(&lamaParkir)
	switch kendaraan {
	case "motor": 
		if lamaParkir <=2 {
			biaya = 7000
			fmt.Println("Biaya parkir: Rp", biaya)
		} else {
			biaya = 9000
			fmt.Println("Biaya parkir: Rp", biaya)
		}
	case "mobil":
		if lamaParkir <=2 {
			biaya = 15000
			fmt.Println("Biaya parkir: Rp", biaya)
		} else {
			biaya = 20000
			fmt.Println("Biaya parkir: Rp", biaya)
		}
	case "truk":
		if lamaParkir <=2 {
			biaya = 25000
			fmt.Println("Biaya parkir: Rp", biaya)
		} else {
			biaya = 35000
			fmt.Println("Biaya parkir: Rp", biaya)
		}
	default:
		fmt.Println("Jenis kendaraan atau durasi tidak valid.")
		fmt.Println("Biaya parkir: Rp 0")
	}
}