package main

import "fmt"

func main() {
	var jam24 int
	fmt.Print("Masukkan jam 0-23: ")
	fmt.Scan(&jam24)

	var jam12 int
	var periode string

	switch {
	case jam24 == 0:
		jam12 = 12
		periode = "AM"
	case jam24 >= 1 && jam24 <= 11:
		jam12 = jam24
		periode = "AM"
	case jam24 == 12:
		jam12 = 12
		periode = "PM"
	case jam24 >= 13 && jam24 <= 23:
		jam12 = jam24 - 12
		periode = "PM"
	default:
		fmt.Println("Input tidak valid. Masukkan jam antara 0-23.")
		return

	}

	fmt.Print(jam12," ", periode)
}