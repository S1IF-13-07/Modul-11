package main

import "fmt"

func main() {
	var nama string
	fmt.Scan(&nama)

	switch nama {
	case "nepenthes", "drosera":
 		fmt.Println("Termasuk Tanaman Karnivora.")
 		fmt.Println("Asli Indonesia.")
 	case "venus", "sarracenia":
 		fmt.Println("Termasuk Tanaman Karnivora.")
 		fmt.Println("Tidak Asli Indonesia.")
	default :
		fmt.Print("Tidak Termasuk Tanaman Karnivora")
	}
}