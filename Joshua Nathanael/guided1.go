package main

import "fmt"

func main() {
	var n int
	var kase int
	fmt.Print("Masukan angkanya dong: ")
	fmt.Scan(&n)
	fmt.Print("Hasilnya adalah: ")
	if n > 12 && n < 23 {
		n = n % 12
		kase = 1
	} else if n == 0 || n == 24 {
		n = 12
		kase = 2

	} else if n <= 13 {
		kase = 2
	}
	switch kase {

	case 1 :
		fmt.Print(n, " PM")
	case 2 :
		fmt.Print(n, " AM")
	default :
		fmt.Print("KELEBIHAN MAS")
	}
}