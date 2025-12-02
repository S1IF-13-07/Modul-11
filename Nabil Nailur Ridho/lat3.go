package main

import "fmt"

func main() {
	var num int
	fmt.Print("Masukkan Bilangan Bulat: ")
	fmt.Scan(&num)

	switch {
	case num%2 == 0 && num%10 != 0 && num%5 != 0:
		fmt.Println("Kategori : Bilangan Genap")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", num, num+1, num*(num+1))

	case num%2 == 1 && num%10 != 0 && (num%5 != 0 || num == 5):
		fmt.Println("Kategori : Bilangan Ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", num, num+1, num+(num+1))

	case num%10 == 0:
		fmt.Println("Kategori : Bilangan Kelipatan 10")
		fmt.Printf("Hasil Pembagian antara %d / %d = %d\n", num, 10, num/10)

	case num%5 == 0:
		fmt.Println("Kategori : Bilangan Kelipatan 5")
		fmt.Printf("Hasil Kuadrat dari %d ^ 2 = %d\n", num, num*num)

	default:
		fmt.Println("Invalid Input")
	}
}