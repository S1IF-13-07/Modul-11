package main
import "fmt"
func main() {
	var jenis string
	var jam int
	fmt.Print("Masukan jenis kendaraan (motor/mobil/turk)")
	fmt.Scan(&jenis)
	fmt.Print("Masukan durasi parkir (jam)")
	fmt.Scan(&jam)
	if jam < 1 {
		jam = 1
	}
	var tarifPerjam int
	switch jenis {
	case "motor":
		tarifPerjam = 2000
	case "mobil":
		tarifPerjam = 5000
	case "truk":
		tarifPerjam = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid! (gunakan: motor, mobil, atau truk)")
		return

	}
	total := tarifPerjam * jam
	fmt.Printf("Total biaya parkir adalah Rp %d\n", total)

	
}
