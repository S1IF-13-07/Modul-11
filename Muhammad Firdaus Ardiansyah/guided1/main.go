package main

import "fmt"

func main() {
	var jam12, jam24 int
	var label string

	if _, err := fmt.Scan(&jam24); err != nil {
		fmt.Println("failed to read input:", err)
		return
	}

	switch {
	case jam24 == 0:
		jam12 = 12
		label = "AM"
	case jam24 < 12:
		jam12 = jam24
		label = "AM"
	case jam24 == 12:
		jam12 = 12
		label = "PM"
	case jam24 > 12 && jam24 < 24:
		jam12 = jam24 - 12
		label = "PM"
	default:
		fmt.Println("invalid hour")
		return
	}

	fmt.Println(jam12, label)
}