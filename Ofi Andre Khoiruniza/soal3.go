package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	switch {
	case n%10 == 0:
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Println("Hasil pembagian antara", n, "/ 10 =", n/10)
	case n%5 == 0:
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Println("Hasil kuadrat dari", n, "^2 =", n*n)
	case n%2 == 0:
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Println("Hasil perkalian dengan bilangan berikutnya", n, "*", n+1, "=", n*(n+1))
	default:
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya", n, "+", n+1, "=", n+(n+1))
	}
}
