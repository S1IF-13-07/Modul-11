package main

import "fmt"

func main() {
	var nama_tanaman string
	fmt.Scan(&nama_tanaman)

	switch nama_tanaman {
	case "nepenthes", "drosera":
		fmt.Println("TERMASUK TANAMAN KARNIVORA.")
		fmt.Println("ASLI INDONESIA.")
	case "venus", "sarracenia":
		fmt.Println("TERMASUK TANAMAN KARNIVORA.")
		fmt.Println("TIDAK ASLI INDONESIA.")
	default:
		fmt.Println("TIDAK TERMASUK TANAMAN KARNIVORA.")
	}
}