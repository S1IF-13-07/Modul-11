package main
import "fmt"
func main() {
	var kendaraan string
	var durasi int
	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)
	switch kendaraan {
	case "Motor":
		if durasi <= 2 {
			fmt.Print("Tarif Parkir: Rp 7000")
		} else if durasi > 2 {
			fmt.Print("Tarif Parkir: Rp 9000")
		}
	case "Mobil":
		if durasi <= 2 {
			fmt.Print("Tarif Parkir: Rp 15000")
		} else if durasi > 2 {
			fmt.Print("Tarif Parkir: Rp 20000")
		}
	case "Truk":
		if durasi <= 2 {
			fmt.Print("Tarif Parkir: Rp 25000")
		} else if durasi > 2 {
			fmt.Print("Tarif Parkir: Rp 35000")
		}
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
		fmt.Println("Tarif Parkir: Rp 0")
	}		
}