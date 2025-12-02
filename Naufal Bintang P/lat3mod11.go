package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	switch {
	case bilangan%10 == 0:
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d", bilangan, bilangan/10)
	case bilangan%2 == 0:
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d", bilangan, bilangan+1, bilangan*(bilangan+1))
	case bilangan%2 != 0 && (bilangan%5 != 0 || bilangan < 10):
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d", bilangan, bilangan+1, bilangan+(bilangan+1))
	case bilangan%5 == 0:
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Printf("Hasil kuadrat dari %d ^ 2 = %d", bilangan, bilangan*bilangan)
	}
}