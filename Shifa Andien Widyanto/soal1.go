package main
import "fmt"
func main() {
	var kadarPh float64

	fmt.Print("Masukan kadar pH : ")
	fmt.Scan(&kadarPh)

	switch {
	case kadarPh < 0 || kadarPh > 14:
		fmt.Println("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14")
	case kadarPh >= 6.5 && kadarPh <= 8.6:
		fmt.Println("Air layak minum")
	case kadarPh < 6.5 || kadarPh > 8.6:
		fmt.Println("Air tidak layak minum")
	}
}