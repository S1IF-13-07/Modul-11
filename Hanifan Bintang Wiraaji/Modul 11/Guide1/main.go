package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	switch {
	case n == 0:
		m := 12
		fmt.Print(m, " AM")
	case n < 12:
		fmt.Print(n, " AM")
	case n == 12:
		fmt.Print(n, " PM")
	case n > 12 && n <24 :
		m :=  n % 12
		fmt.Print(m, " PM")
	default:
		fmt.Print("Masukan angka yg bener")
	}
}
