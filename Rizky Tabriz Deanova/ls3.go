package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&input)
	switch {
	case input%10 == 0:
		hasil := input / 10
		fmt.Printf("Kategori: Bilangan Kelipatan 10\nHasil pembagian antara %d / 10 = %d\n", input, hasil)
	case input%5 == 0:
		hasil := input * input
		fmt.Printf("Kategori: Bilangan Kelipatan 5\nHasil kuadrat dari %d ^ 2 = %d\n", input, hasil)
	case input%2 == 0:
		nambah1 := input + 1
		hasil := input * nambah1
		fmt.Printf("Kategori: Bilangan Genap\nHasil perkalian dengan bilangan berikutnya %d * %d = %d\n", input, nambah1, hasil)
	default:
		nambah1 := input + 1
		hasil := input + nambah1
		fmt.Printf("Kategori: Bilangan Ganjil\nHasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", input, nambah1, hasil)
	}
}
