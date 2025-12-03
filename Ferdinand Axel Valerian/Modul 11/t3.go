package main

import "fmt"

func main() {
	var a int
	fmt.Print("Masukkan sebuah bilangan (ganjil, genap, kelipatan 5, kelipatan 10) : ")
	fmt.Scanln(&a)

	switch {
	case a%2 != 0:
		fmt.Println("Kategori: Bilangan Ganjil")
		hasil := a + (a + 1)
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", a, a+1, hasil)

	case a%2 == 0:
		fmt.Println("Kategori: Bilangan Genap")
		hasil := a * (a + 1)
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", a, a+1, hasil)

	case a%5 == 0:
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		hasil := a ^ 2
		fmt.Printf("Hasil kuadrat dari %d ^2 = %d\n", a, hasil)

	case a%10 == 0:
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		hasil := a / 10
		fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", a, hasil)

	default:
		fmt.Println("Bilangan tidak termasuk kategori yang ditentukan.")
	}
}
