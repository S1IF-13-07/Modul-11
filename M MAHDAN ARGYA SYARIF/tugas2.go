package main
import "fmt"

func main() {
	var nama_kendaraan string
	var jam, tarif, motor, mobil, truk int
	fmt.Scan(&nama_kendaraan, &jam)

	switch nama_kendaraan {
	case "Motor":
		if jam >= 1 {
			tarif = 2000
			motor = jam * tarif
			fmt.Print("Rp.", motor)
		}
	case "Mobil":
		if jam >= 1 {
			tarif = 5000
			mobil = jam * tarif
			fmt.Print("Rp.", mobil)
		}
	case "Truk":
		if jam >= 1 {
			tarif = 8000
			truk = jam * tarif
			fmt.Print("Rp.", truk)
		}
	}
}