package main

import "fmt"

func main() {
	var jam24 int
	fmt.Print("Masukkan jam dalam format 24 jam: ")
	fmt.Scan(&jam24)

	var period string
	var jam12 int

	if jam24 == 0 {
		jam12 = 12
		period = "AM"
	} else if jam24 == 12 {
		jam12 = 12
		period = "PM"
	} else if jam24 > 12 {
		jam12 = jam24 - 12
		period = "PM"
	} else {
		jam12 = jam24
		period = "AM"
	}
	fmt.Printf("%d %s\n", jam12, period)
}