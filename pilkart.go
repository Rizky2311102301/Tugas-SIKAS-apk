package main
import "fmt"
func Datavote(suara *[21]int) (int, int) {
    var totalmasuk, totalsah int
    var angka int
    fmt.Println("Masukkan data suara:")
    fmt.Scan(&angka)
    for angka != 0 {
        totalmasuk++
        if angka >= 1 && angka <= 20 {
            suara[angka]++
            totalsah++
        }
        fmt.Scan(&angka)
    }
    return totalmasuk, totalsah
}
func cetakHasil(suara [21]int, totalmasuk, totalsah int) {
    fmt.Printf("Suara masuk: %d\n", totalmasuk)
    fmt.Printf("Suara sah: %d\n", totalsah)
    for i := 1; i <= 20; i++ {
        if suara[i] > 0 {
            fmt.Printf("%d : %d\n", i, suara[i])
        }
    }
}
func main() {
    var suara [21]int
    totalmasuk, totalsah := Datavote(&suara)
    cetakHasil(suara, totalmasuk, totalsah)
}