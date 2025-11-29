package main

import "fmt"

func main() {
	var jam12, jam24 int
	var keterangan string

	fmt.Print("Masukan Jam : ")
	fmt.Scan(&jam24)

	switch {
	case jam24 == 0:
		jam12 = 12
		keterangan = "AM"
	case jam24 < 12:
		jam12 = jam24
		keterangan = "AM"
	case jam24 == 12:
		jam12 = 12
		keterangan = "PM"
	case jam24 > 12:
		jam12 = jam24 - 12
		keterangan = "PM"
	}
	fmt.Println(jam12, keterangan)
}
