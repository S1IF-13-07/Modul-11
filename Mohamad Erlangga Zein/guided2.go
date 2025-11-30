package main 
import "fmt"

func main() {
	var tanaman string
	fmt.Print("masukkan nama tanaman: ")
	fmt.Scan(&tanaman)

	switch tanaman {
	case "nephentes":
		fmt.Print("Termasuk tanaman karnivora asli indonesia")
	case "venus":
		fmt.Print("Termasuk tanaman karnovora bukan asli Indonesia")
	case "karedok":
		fmt.Print("Tidak termasuk tanaman karnivora")
	default:
		fmt.Print("ga ada di opsi")
	}
}