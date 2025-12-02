package main

import "fmt"

func main(){
	var tanaman string
	fmt.Print("Masukkan nama tanaman: ")
	fmt.Scan(&tanaman)

	switch tanaman {
	case "nepenthes", "kantongsemar" :
		fmt.Println("termasuk tanaman karnivora")
		fmt.Println("asli indonesia")
	case "venus", "venus flytrap" :
		fmt.Println("termasuk tanaman karnivora")
		fmt.Println("bukan asli indonesia")
	default :
		fmt.Print("bukan tanaman karnivora")
	}
}