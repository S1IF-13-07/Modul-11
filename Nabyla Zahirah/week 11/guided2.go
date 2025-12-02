package main
import "fmt"
func main (){
	var nama_tanaman string 
	fmt.Print("Masukkan nama tanaman: ")
	fmt.Scan(&nama_tanaman)
	switch nama_tanaman {
		case "nepenthes", "drosera":
			fmt.Println("Termasuk Tanaman Karnivora.")
			fmt.Println("Asli Indonesia.")
		case "venus", "sarracenia":
			fmt.Println("Termasuk Tanaman Karnivora.")
			fmt.Println("Tidak Asli Indonesia.")
		default:
			fmt.Println("Tidak termasuk tanaman karnivora.")
	
	}
	
}