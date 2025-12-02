package main

import (
    "fmt"
    "strings"
)

func main() {
    var kendaraan string
    var durasi, tarif, total int

    fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
    fmt.Scan(&kendaraan)
    kendaraan = strings.ToLower(kendaraan)

    fmt.Print("Masukkan durasi parkir (jam): ")
    fmt.Scan(&durasi)

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
    default:
        fmt.Println("Jenis kendaraan tidak valid")
        return
    }

    total = tarif * durasi

    fmt.Println("Total biaya parkir: Rp", total)
}
