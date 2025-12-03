package main
import "fmt"

func main() {
    var jenis string
    var jam int
    
    fmt.Print("Masukkan jenis kendaraan dan durasi (contoh: motor 3): ")
    fmt.Scan(&jenis, &jam)
    
    if jam < 1 {
        jam = 1
    }
    
    var tarifPerJam int
    
    switch jenis {
    case "motor":
        tarifPerJam = 2000
    case "mobil":
        tarifPerJam = 5000
    case "truk":
        tarifPerJam = 8000
    default:
        fmt.Println("Jenis kendaraan tidak valid")
        return
    }
    
    total := tarifPerJam * jam
    fmt.Printf("Rp %d\n", total)
}