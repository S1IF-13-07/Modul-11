package main
import "fmt"
func main (){
	var jenis string
	var durasi, tarif int 

	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&jenis)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	switch jenis {
	case "motor":
		switch {
		case durasi >= 1 && durasi <= 2:
			tarif = 7000
		case durasi > 2: 
			tarif = 9000
		}
	case "mobil":
		switch {
		case durasi >= 1 && durasi <= 2:
			tarif = 15000
		case durasi > 2: 
			tarif = 20000
		}
	case "truk":
		switch {
		case durasi >= 1 && durasi <= 2:
			tarif = 25000
		case durasi > 2: 
			tarif = 35000
		}
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
	}
	fmt.Printf("Trif parkir: Rp%d\n", tarif)
}