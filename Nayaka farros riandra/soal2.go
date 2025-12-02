package main

import "fmt"

func main(){
	var kendaran string
	var durasi int
	tarifparkir := 0
	fmt.Print("masukkan jenis kendaraan: ")
	fmt.Scan(&kendaran)
	fmt.Print("masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	switch kendaran {
	case "motor":
		if durasi == 1{
			fmt.Print("Rp. 2.000")
		} else if durasi >= 2{
			tarifparkir = 2000 * durasi
			fmt.Print("Rp.", tarifparkir)
		}
	case "mobil":
		if durasi == 1{
			fmt.Print("Rp. 5.000")
		} else if durasi >= 2{
			tarifparkir = 5000 * durasi
			fmt.Print("Rp.", tarifparkir)
		}
	case "truk":
		if durasi == 1{
			fmt.Print("Rp. 8.000")
		} else if durasi >= 2{
			tarifparkir = 8000 * durasi
			fmt.Print("Rp.", tarifparkir)
		}
	}
}