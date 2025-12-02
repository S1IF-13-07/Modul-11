package main

import "fmt"

func main(){
	var tanaman string
	fmt.Scanf("%s",&tanaman)

	switch tanaman {
	case "nepenthes":
		fmt.Print("Termasuk Tanaman Karnivora \nAsli Indonesia")
	case "venus":
		fmt.Print("Termasuk Tanaman Karnivora \nBukan Asli Indonesia")
	default:
		fmt.Print("Tidak termasuk tanaman")
	}
}