package main
import "fmt"

func main() {
	var kendaraan string
	var jam, tarif, motor, mobil, truk int
	fmt.Print("Masukkan kendaraan dan jumlah total parkir(jam): ")
	fmt.Scan(&kendaraan, &jam)

	switch kendaraan {
	case "motor":
		if jam >= 1 {
			tarif = 2000
			motor = jam * tarif
			fmt.Print("Rp.", motor)
		}
	case "mobil":
		if jam >= 1 {
			tarif = 5000
			mobil = jam * tarif
			fmt.Print("Rp.", mobil)
		}
	case "truk":
		if jam >= 1 {
			tarif = 8000
			truk = jam * tarif
			fmt.Print("Rp.", truk)
		}
	}
}