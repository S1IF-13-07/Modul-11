package main

import "fmt"

func main() {
	var num int
	fmt.Print("MASUKAN BILANGAN NYA MAS: ")
	fmt.Scan(&num)

	switch {
	case num == 5:
		hasil := num + (num + 1)
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", num, num+1, hasil)
	case num%10 == 0:
		hasil := num / 10
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", num, hasil)
	case num%5 == 0:
		hasil := num * num
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Printf("Hasil kuadrat dari %d ^2 = %d\n", num, hasil)
	case num%2 == 0:
		hasil := num * (num + 1)
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", num, num+1, hasil)
	default:
		hasil := num + (num + 1)
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", num, num+1, hasil)
	}
}
