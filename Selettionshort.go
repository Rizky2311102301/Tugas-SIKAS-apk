package main
import "fmt"
func main() {
    var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    for i := 0; i < len(arr); i++ {
        min_idx := i
        for j := i + 1; j < len(arr); j++ {
            if arr[min_idx] > arr[j] {
                min_idx = j
            }
        }
        temp := arr[min_idx]
        arr[min_idx] = arr[i]
        arr[i] = temp
    }
    fmt.Println(arr)
}
