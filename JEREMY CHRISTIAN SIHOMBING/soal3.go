package main
import "fmt"
func main() {

	var bilangan int
	fmt.Print("Masukkan sebuah bilangan bulat: ")
	fmt.Scan(&bilangan)

	switch {
	case bilangan % 10 == 0:
		fmt.Println("Kategori: Bilangan kelipatan 10")
		bagi := bilangan / 10
		fmt.Println("Hasil pembagian antara", bilangan, "/", "10", "=", bagi)
	case bilangan % 5 == 0 && bilangan != 5:
		fmt.Println("Kategori: Bilangan kelipatan 5")
		pangkat := bilangan * bilangan
		fmt.Println("Hasil kuadrat dari", bilangan,"^2", "=", pangkat )
	case bilangan % 2 == 1:
		fmt.Println("Kategori: Bilangan Ganjil")
		tambah := bilangan + 1
		total := bilangan + tambah
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya", bilangan, "+", tambah, "=", total)
	case bilangan % 2 == 0:
		fmt.Println("Kategori: Bilangan Genap")
		kali := bilangan + 1
		total := bilangan * kali
		fmt.Println("Hasil perkalian dengan bilangan berikutnya", bilangan, "*", kali, "=", total)
	}
}