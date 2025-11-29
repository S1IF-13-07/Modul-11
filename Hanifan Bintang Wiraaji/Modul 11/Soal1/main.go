package main

import "fmt"

func main() {
	var ph float64
	var label string
	fmt.Scan(&ph)
	switch {
		case ph >= 6.5 && ph <= 8.6 :
			label = "Air layak minum"
		case (ph < 6.5 || ph > 8.6) && ph >= 0 && ph <= 14 :
			label = "Air tidak layak minum"
		default :
			label = "Nilai pH tidak valid. Nilai pH harus antara 0 dan 14."
	}
	fmt.Print(label)
}