package main

import "fmt"

func main() {
    var input int
    fmt.Scan(&input)
    fmt.Print("Kategori: ")

    switch {
    case input%10 == 0:
        fmt.Println("Bilangan Kelipatan 10")
        hasil := input / 10
        fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", input, hasil)

    case input%5 == 0 && input != 5:
 		fmt.Println("Bilangan Kelipatan 5")
        hasil := input * input
        fmt.Printf("Hasil kuadrat dari %d^2 = %d\n", input, hasil)

    case input%2 == 0:
        fmt.Println("Bilangan Genap")
        hasil := input * (input + 1)
        fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n",
            input, input+1, hasil)

	case input%2 == 1:
        fmt.Println("Bilangan Ganjil")
        hasil := input + (input + 1)
        fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n",
            input, input+1, hasil)
    }
}
