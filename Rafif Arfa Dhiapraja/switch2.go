package main

import "fmt"

func main() {
	fmt.Print("Masukkan tanaman : ")
	var tanaman string
	fmt.Scan(&tanaman)

	switch tanaman {
	case "nepenthes" :
		fmt.Println("Termasuk tanaman karnivora")
		fmt.Println("Asli Indonesia")
	case "venus" :
		fmt.Println("Termasuk tanaman karnivora")
		fmt.Println("bukan asli indonesia")
	default : 
		fmt.Println("Bukan tanaman karnivora")
	}
}