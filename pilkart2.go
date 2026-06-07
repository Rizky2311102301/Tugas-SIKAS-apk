package main
import "fmt"
func Data(suara *[21]int) (int, int) {
    var totalmasuk, totalsah int
    var angka int
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
func Cariketuadanwakil(suara [21]int) (int, int) {
    var ketua, wakil int
    var maxsuara, maxsuara2 int
    for i := 1; i <= 20; i++ {
        if suara[i] > maxsuara {
            maxsuara = suara[i]
            ketua = i
        }
    }
    for i := 1; i <= 20; i++ {
        if i != ketua {
            if suara[i] > maxsuara2 {
                maxsuara2 = suara[i]
                wakil = i
            } else if suara[i] == maxsuara2 && wakil > i {
                wakil = i
            }
        }
    }
    
    return ketua, wakil
}
func main() {
    var suara [21]int
    totalMasuk, totalSah := Data(&suara)
    ketua, wakil := Cariketuadanwakil(suara)
    fmt.Printf("Suara masuk: %d\n", totalMasuk)
    fmt.Printf("Suara sah: %d\n", totalSah)
    fmt.Printf("Ketua RT: %d\n", ketua)
    fmt.Printf("Wakil ketua: %d\n", wakil)
}