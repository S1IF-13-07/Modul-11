package main
import "fmt"
func main(){
	var tanaman string
	fmt.Scan(&tanaman)

	switch tanaman {
	case "nepenthese":
		fmt.Print("Termasuk Tanaman Karnivora Asli Indonesia")
		case "venus":
		fmt.Print("Termasuk Tanaman Karnivora Bukan Asli Indonesia")
		default:
			fmt.Print("Tidak termasuk Tanaman Karnivora")
	}

}