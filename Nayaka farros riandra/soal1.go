package main

import "fmt"

func main(){
	var pH float64
	fmt.Print("Masukkan kadar pH: ")
	fmt.Scanln(&pH)

	switch{
		case pH >= 6.5 && pH <= 8.6:
			fmt.Println("Air layak minum")
		case pH <= 14 && pH >= 0 && (pH < 6.5 || pH > 8.6):
			fmt.Print("Air tidak layak minum")
		case pH > 14:
			fmt.Print("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")
	}
}