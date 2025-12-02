package main
import "fmt"
func main() {
	var pH float64
	fmt.Print("Masukan nilai pH")
	fmt.Scan(&pH)

	if pH < 0 || pH > 14 {
		fmt.Println("Input tidak valid. rentang pH 0 - 14")
	} else if pH >= 6.5 && pH <= 8.6{
		fmt.Println("Air layak minum")
	} else {
		fmt.Println("Air tidak layak minum")
	}
	}