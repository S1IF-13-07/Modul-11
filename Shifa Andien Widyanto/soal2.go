package main
import "fmt"
func main(){
	var kendaraan string
	var durasi, tarif int

	fmt.Print("Masukan jenis kendaraan (Motor/Mobil/Truk) : ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukan durasi parkir : ")
	fmt.Scan(&durasi)

	switch {
	case kendaraan == "Motor" && durasi <= 1:
		tarif = 2000
	case kendaraan == "Motor" && durasi >= 1:
		tarif = 2000 * durasi
	case kendaraan == "Mobil" && durasi <= 1:
		tarif = 5000
	case kendaraan == "Mobil" && durasi >= 1:
		tarif = 5000 * durasi
	case kendaraan == "Truk" && durasi <= 1:
		tarif = 8000
	case kendaraan == "Truk" && durasi >= 1:
		tarif = 8000 * durasi
	}
	fmt.Printf("Tarif Parkir : Rp %d\n", tarif)
}
