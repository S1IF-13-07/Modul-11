package main

import "fmt"

func main() {
    var jam int
    fmt.Print("Masukkan jam (0-23): ")
    fmt.Scan(&jam)

    if jam < 0 || jam > 23 {
        fmt.Println("Input tidak valid!")
        return
    }

    jam12 := jam
    period := "AM"

    if jam == 0 {
        jam12 = 12
    } else if jam == 12 {
        period = "PM"
    } else if jam > 12 {
        jam12 = jam - 12
        period = "PM"
    }

    fmt.Printf("%d %s\n", jam12, period)
}