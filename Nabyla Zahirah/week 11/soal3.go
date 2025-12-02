package main
import "fmt"
func main() {
	var bilangan int

	fmt.Print("Masukan satu bilangan bulat : ")
	fmt.Scan(&bilangan)

	switch{
	case bilangan % 10 == 0:
		hasil := bilangan / 10
		fmt.Println("Kategori: Bilangan kelipatan 10")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", bilangan, hasil)
	case bilangan % 5 == 0 && bilangan != 5:
		hasil := bilangan * bilangan
		fmt.Println("Kategori: Bilangan kelipatan 5")
		fmt.Printf("Hasil kuadrat dari %d ^2 = %d\n", bilangan, hasil)
	case bilangan % 2 ==0:
		a := bilangan + 1
		hasil := bilangan * a
		fmt.Println("Kategori: Bilangan genap")
		fmt.Printf("Hasil perkalian dengan bilangan beriktnya %d * %d = %d\n", bilangan, a, hasil)
	default:
		a := bilangan + 1 
		hasil := bilangan + a
		fmt.Println("Kategori: Bilangan ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", bilangan, a, hasil)
	
	}
}