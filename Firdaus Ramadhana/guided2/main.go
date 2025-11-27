package main 

import "fmt"

func main() {
	var namaTumbuhan string
	fmt.Print("Masukkan nama tumbuhan: ")
	fmt.Scan(&namaTumbuhan)
	switch namaTumbuhan {
	case "nepenthes":
		fmt.Println("Termasuk tumbuhan karnivora.")
		fmt.Println("Asli Indonesia.")
	case "venus":
		fmt.Println("Termasuk tumbuhan karnivora.")
		fmt.Println("Bukan Asli Indonesia.")
	default:
		fmt.Println("Tidak termasuk tanaman Karnivora")
	}
}