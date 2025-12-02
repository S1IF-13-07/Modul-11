package main
import "fmt"
func main() {

    var kendaraan string
    var jam, tarif, Motor, Mobil, Truk int
    fmt.Print("Masukkan jenis kendaraan dan durasi parkir dalam jam: ")
    fmt.Scan(&kendaraan, &jam)

    switch kendaraan {
    case "Motor":
		if jam >= 1 {
			tarif = 2000
			Motor = tarif * jam
			fmt.Print("Rp.", Motor)
		}
    case "Mobil":
		if jam >= 1 {
			tarif = 5000
			Mobil = tarif * jam
			fmt.Print("Rp.", Mobil)
		}
    case "Truk":
		if jam >= 1 {
			tarif = 8000
			Truk = tarif * jam
			fmt.Print("Rp.", Truk)
		}
    }
}