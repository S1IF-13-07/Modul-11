package main 
import "fmt"

func main() {
	var durasiParkir int
	var jenisKendaraan string
	var biayaParkir int
	fmt.Print("Masukkan jenis kendaraan (mobil/motor/truk): ")
	fmt.Scan(&jenisKendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasiParkir)
	switch jenisKendaraan {
	case "motor":
		biayaParkir = 2000
		if durasiParkir <= 1 {
			fmt.Println("Biaya parkir: Rp", biayaParkir * durasiParkir)
		} else {
			fmt.Println("Biaya parkir: Rp", biayaParkir * durasiParkir)
		}
	case "mobil":
		biayaParkir = 5000
		if durasiParkir <= 1 {
			fmt.Println("Biaya parkir: Rp", biayaParkir * durasiParkir)
		} else {
			fmt.Println("Biaya parkir: Rp", biayaParkir * durasiParkir)
		}
	case "truk":
		biayaParkir = 8000
		if durasiParkir <= 1 {
			fmt.Println("Biaya parkir: Rp", biayaParkir * durasiParkir)
		} else {
			fmt.Println("Biaya parkir: Rp", biayaParkir * durasiParkir)
		}
	}
}