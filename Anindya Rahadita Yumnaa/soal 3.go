package main

import "fmt"

func main() {
    var n int
	fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&n)

    if n%10 == 0 {
        fmt.Println("Kategori: Bilangan Kelipatan 10")
        hasil := n / 10
        fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", n, hasil)

    } else if n%5 == 0 {
        fmt.Println("Kategori: Bilangan Kelipatan 5")
        hasil := n * n
        fmt.Printf("Hasil kuadrat dari %d ^2 = %d\n", n, hasil)

    } else if n%2 == 0 {
        fmt.Println("Kategori: Bilangan Genap")
        next := n + 1
        hasil := n * next
        fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", n, next, hasil)

    } else {
        fmt.Println("Kategori: Bilangan Ganjil")
        next := n + 1
        hasil := n + next
        fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", n, next, hasil)
    }
}