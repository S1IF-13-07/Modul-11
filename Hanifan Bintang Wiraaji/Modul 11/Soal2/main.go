package main

import "fmt"

func main() {
	var m, tarif int
	var n string
	fmt.Print("Masukan kendaraan: ")
	fmt.Scanln(&n)
	fmt.Print("Berapa jam parkir: ")
	fmt.Scanln(&m)
	switch {
		case n == "Motor" || n == "motor" :
			tarif = m * 2000
		case n == "Mobil" || n == "mobil" :
			tarif = m * 5000
		case n == "Truk" || n == "truk" :
			tarif = m * 8000
		default :
			tarif = 0
	}
	fmt.Printf("Rp. %v", tarif)
}