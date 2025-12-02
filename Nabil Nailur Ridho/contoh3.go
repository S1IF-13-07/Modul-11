package main
import "fmt"
func main(){
	var kendaraan string
	var durasi, tarif int
	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	switch kendaraan {
	case "motor":
		if durasi <= 2{
		tarif = 7000
		fmt.Print("Tarif parkir:Rp.", tarif)
		} else {
			tarif = 9000
		fmt.Print("Tarif parkir:Rp.", tarif)
		}
		case "mobil":
		if durasi <= 2{
		tarif = 15000
		fmt.Print("Tarif parkir:Rp.", tarif)
		} else {
			tarif = 20000
		fmt.Print("Tarif parkir:Rp.", tarif)
		}
			case "truk":
			if durasi <= 2{
		tarif = 25000
		fmt.Print("Tarif parkir:Rp.", tarif)
		} else {
			tarif = 35000
		fmt.Print("Tarif parkir:Rp.", tarif)
		}
	default:
		tarif = 0
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
		fmt.Println("Tarif parkir:Rp. ", tarif)
	}
}