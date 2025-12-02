package main
import "fmt"

func main() {
  var jenis string
  var jam float64
  fmt.Print("Masukkan jenis kendaraan dan durasi: ")
  fmt.Scan(&jenis, &jam)

  if jam < 1 {
    jam = 1
  }
  var tarifPerJam int
    if jenis == "motor" {
      tarifPerJam = 2000
    } else if jenis == "mobil" {
      tarifPerJam = 5000
    } else if jenis == "truk" {
      tarifPerJam = 8000
    } else {
      fmt.Println("Jenis kendaraan tidak dikenali.")
    return
  }
    totalBiaya := tarifPerJam * int(jam)
      fmt.Printf("Total biaya parkir: Rp %d\n", totalBiaya)
}
