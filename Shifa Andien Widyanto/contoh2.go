package main
import "fmt"
func main() {
	var tanaman string

	fmt.Print("Masukan satu nama tanaman : ")
	fmt.Scan(&tanaman)

	switch tanaman {
	case "nepenthes" :
		fmt.Println("Termasuk Tanaman Karnivora.")
		fmt.Println("Asli Indonesia.")
	case "venus" :
		fmt.Println("Termasuk Tanaman Karnivora.")
		fmt.Println("Bukan Asli Indonesia.")
	default:
		fmt.Println("Tidak termasuk tanaman karnivora.")
	}
}