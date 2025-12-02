package main

import "fmt"

func main(){
	var jeniskendaraan string
	var durasi int
	fmt.Print("Masukkan jenis kendaraan : ")
	fmt.Scan(&jeniskendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	switch jeniskendaraan {
	case "motor":
		if durasi <= 2 {
			fmt.Print("Rp. 7.000")
		} else {
			fmt.Print("Rp. 9.000")
		}
	case "mobil":
		if durasi <= 2 {
			fmt.Print("Rp. 15.000")
		} else {
			fmt.Print("Rp. 20.000")
		}
	case "truk":
		if durasi <= 2 {
			fmt.Print("Rp. 25.000")
		} else {
			fmt.Print("Rp. 35.000")
		}
	default:
		fmt.Print("Jenis kendaraan tidak dikenali")
	}
}