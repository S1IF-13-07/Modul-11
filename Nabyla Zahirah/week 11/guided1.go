package main
import "fmt"
func main (){
	var jam int
	fmt.Print("Masukkan jam : ")
	fmt.Scan(&jam)

	switch jam {
	case 0:
		fmt.Println("12 AM")
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		fmt.Println(jam, "AM")
	case 12: 
		fmt.Println("12 PM")
	case 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23:
		fmt.Println(jam-12, "PM")
	default:
		fmt.Println("input tidak valid")
	}
}