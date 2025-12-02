package main
import "fmt"
func main(){
	var jam int

	fmt.Print("Masukkin jam: ")
	fmt.Scan(&jam)

	switch {
	case jam == 0:
		fmt.Print("12 AM")
	case jam < 12:
		fmt.Println(jam, "AM")
	case jam >= 12:
		jam -= 12
		fmt.Print(jam, "PM")
	default:
		fmt.Print("Invalid")
	}
}