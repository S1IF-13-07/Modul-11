package main
import "fmt"

func main() {
	var kendaraan string
	var jam, parkir int
	fmt.Print("Masukkan jenis kendaraan(Motor/Mobil/Truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&jam)

	switch kendaraan {
	case "Motor":
		parkir = 7000
		if jam == 1 || jam == 2{
			fmt.Print("Tarif Parkir: Rp ", parkir)
		} else if jam > 2 {
			motor := parkir + 2000
			fmt.Print("Tarif Parkir: Rp ", motor)
		}
	case "Mobil":
		parkir = 15000
		if jam == 1 || jam == 2{
			fmt.Print(parkir)
		} else if jam > 2{
			mobil := parkir + 5000
			fmt.Print("Tarif Parkir: Rp ", mobil)
		}
	case "Truk":
		parkir = 25000
		if jam == 1 || jam == 2{
			fmt.Print("Tarif Parkir: Rp ", parkir)
		} else if jam > 2 {
			truk := parkir + 10000
			fmt.Print("Tarif Parkir: Rp ", truk)
		}
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
		fmt.Println("Tarif Parkir: Rp 0")
	}
}