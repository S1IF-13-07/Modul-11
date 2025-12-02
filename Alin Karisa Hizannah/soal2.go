package main

import "fmt"

func main() {
    var kendaraan string
    var durasi int
    var tarif int

    fmt.Scan(&kendaraan, &durasi)

    if durasi < 1 {
        durasi = 1
    }

    switch kendaraan {
    case "motor":
        tarif = 2000
    case "mobil":
        tarif = 5000
    case "truk":
        tarif = 8000
    }

    total := tarif * durasi
    fmt.Printf("Rp %d\n", total)
}
