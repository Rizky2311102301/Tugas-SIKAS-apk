package main

import "fmt"

func main() {
    input := 10
    if input%2 == 0 {
        for i := 0; i <= 9; i++ {
            if i > 0 {
                fmt.Print("+")
            }
            fmt.Print(i)
        }
        fmt.Println()
        fmt.Println("Inputan adalah genap")  // Fixed: Changed from "ganjil" to "genap"
    } else {
        fmt.Println("Inputan adalah ganjil") // Fixed: Changed from "genap" to "ganjil"
    }
}