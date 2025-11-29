package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	switch {
	case n > 10 && n%10 == 0:
		m := 10
		j := n / m
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %v / %v = %v", n, m, j)
	case n > 5 && n%5 == 0:
		j := n * n
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %v ^2 = %v", n, j)
	case n%2 == 1:
		m := n + 1
		j := m + n
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %v + %v = %v", n, m, j)
	case n%2 == 0:
		m := n + 1
		j := m * n
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %v * %v = %v", n, m, j)
	default:
		fmt.Print("Masukan angka yang benar")
	}
}
