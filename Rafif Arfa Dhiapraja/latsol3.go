package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&n)

	switch {
		case n%10 == 0:
			fmt.Println("Kategantaraori: Bilangan Kelipatan 10")
			fmt.Printf("Hasil pembagian  %d / 10= %d\n", n, n/10)
		case n%5 == 0 && n != 5:
			fmt.Println("Kategori: Bilangan Kelipatan 5")
			fmt.Printf("hasil kuadrat dari %d adalah %d\n", n, n*n)
		case n%2 == 1:
			fmt.Println("Kategori: bilangan Ganjil")
			fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", n, n+1, n+(n+1))
		case n%2 == 0:
			fmt.Println("Kategori: Bilangan Genap")
			fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", n, n+1, n*(n+1))
		
	}
}