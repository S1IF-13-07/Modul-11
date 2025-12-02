package main

import "fmt"

func main(){
	var angka int

	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&angka)

	switch {
	case angka%2 == 0 && angka%10 != 0 && angka%5 != 0:
		fmt.Println("kategor : bilangan genap")
		fmt.Println("hasil perkalian dengan bilangan berikutnya", angka, "*", angka+1 ,"=", angka*(angka+1))
	case angka%2 == 1 && angka%10 != 0 && angka%5 != 0:
		fmt.Println("kategori : bilangan ganjil")
		fmt.Println("hasil penjumlahan dengan bilangan berikutnya", angka ,"+", angka+1, "=",angka+(angka+1))
	case angka%10 == 0:
		fmt.Println("kategori : bilangan kelipatan 10")
		fmt.Println("hasil pembagian antara", angka, "/", 10, "=", angka/10)
	case angka%5 == 0:
		fmt.Println("kategori : bilangan kelipatan 5")
		fmt.Println("hasil kuadrat dari", angka, "^", 5, "=", angka*angka)
	default:
		fmt.Println("kategori : angka tidak valid")
	}
}