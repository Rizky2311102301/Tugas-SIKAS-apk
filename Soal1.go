package main
import "fmt"
func main() {
    for i := 0; i <= 13; i++ {
        if i == 6 || i == 8 {
            continue
        }
        if i > 0 && i != 1 {
            fmt.Print("+")
        }
        fmt.Print(i)
    }
    fmt.Println()
}