package main

import "fmt"

func main(){
	var jam int
	fmt.Print(" masukkan jam (0-23) : ")
	fmt.Scan(&jam)

	switch {
	case jam == 0:
		fmt.Print("12 AM")
	case jam < 12:
		fmt.Printf("%d AM", jam)
	case jam == 12:
		fmt.Print("12 PM")
	default:
		fmt.Printf("%d PM", jam-12)
	}
}