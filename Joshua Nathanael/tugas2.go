package main
import "fmt"
func main() {
	var kendaraan string
	var total, waktu int

	fmt.Print("MASUKAN JENIS KENDARAAN: ")
	fmt.Scan(&kendaraan)
	fmt.Print("DURASI PARKIR: ")
	fmt.Scan(&waktu)
	switch kendaraan {
	case "motor", "Motor":
		total = waktu * 2000
	case "mobil", "Mobil":
		total = waktu * 5000
	case "truk", "Truk":
		total = waktu * 8000
	default:
		fmt.Print("Invalid")
	}
	fmt.Print("Tarif Parkir ", waktu, " jam : Rp. ", total)
}