package main
import "fmt"

func main(){
	var bil int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bil)

	switch {
	case bil % 10 == 0:
		fmt.Println("Kategori: Bilangan kelipatan 10")
		bagi := bil / 10
		fmt.Println("Hasil pembagian antara", bil, "/", "10", "=", bagi)
	case bil % 5 == 0 && bil != 5:
		fmt.Print("Kategori: Bilangan kelipatan 5")
		pangkat := bil * bil
		fmt.Println("Hasil kuadrat dari", bil,"^2", "=", pangkat )
	case bil % 2 == 1:
		fmt.Println("Kategori: Bilangan ganjil")
		tambah := bil + 1
		total := bil + tambah
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya", bil, "+", tambah, "=", total)
	case bil % 2 == 0:
		fmt.Println("Kategori: Bilangan genap")
		kali := bil + 1
		total := bil * kali
		fmt.Println("Hasil perkalian dengan bilangan berikutnya", bil, "*", kali, "=", total)
	}
}