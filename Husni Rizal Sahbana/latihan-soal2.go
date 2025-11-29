package main

import "fmt"

func main() {
	var nama_tanaman string
	fmt.Print("Masukan Nama Tanaman : ")
	fmt.Scan(&nama_tanaman)

	switch {
	case nama_tanaman == "nepenthes":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Asli Indonesia")
	case nama_tanaman == "venus":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Bukan Asli Indonesia")
	default:
		fmt.Println("Bukan Tanaman Karnivora")
	}
}
