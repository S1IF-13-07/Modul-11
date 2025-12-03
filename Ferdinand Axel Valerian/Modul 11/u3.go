package main
import "fmt"

func main() {
    var jenis string
    var jam, tarif int
    
    fmt.Print("Jenis kendaraan (Motor/Mobil/Truk): ")
    fmt.Scan(&jenis)
    
    fmt.Print("Durasi parkir (jam): ")
    fmt.Scan(&jam)
    
    if jam < 1 {
        fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
        tarif = 0
    } else {
        switch jenis {
        case "Motor":
            if jam <= 2 {
                tarif = 7000
            } else {
                tarif = 9000
            }
        case "Mobil":
            if jam <= 2 {
                tarif = 15000
            } else {
                tarif = 20000
            }
        case "Truk":
            if jam <= 2 {
                tarif = 25000
            } else {
                tarif = 35000
            }
        default:
            fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
            tarif = 0
        }
    }
    
    fmt.Printf("Tarif Parkir: Rp %d\n", tarif)
}