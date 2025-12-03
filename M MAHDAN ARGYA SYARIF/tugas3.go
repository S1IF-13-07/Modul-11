package main
import "fmt"

func main(){
	var n int
	fmt.Scan(&n)

	switch {
	case n % 10 == 0:
		fmt.Println("Kategori: bilangan kelipatan 10")
		bagi := n / 10
		fmt.Println("Hasil pembagian antara", n, "/", "10", "=", bagi)
	case n % 5 == 0 && n != 5:
		fmt.Print("Kategori: bilangan kelipatan 5")
		pangkat := n * n
		fmt.Println("Hasil kuadrat dari", n,"^2", "=", pangkat )
	case n % 2 == 1:
		fmt.Println("Kategori: bilangan ganjil")
		tambah := n + 1
		total := n + tambah
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya", n, "+", tambah, "=", total)
	case n % 2 == 0:
		fmt.Println("Kategori: bilangan genap")
		kali := n + 1
		total := n * kali
		fmt.Println("Hasil perkalian dengan bilangan berikutnya", n, "*", kali, "=", total)
	}
}