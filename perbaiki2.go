package main

import "fmt"

func main() {
    var arr [10]int = [10]int{1, 8, 11, 25, 26, 32, 52, 77, 90, 93}
    fmt.Println(arr)

    var kr, kn int = 0, len(arr)
    var found bool = false
    var x, idx int = 100, -1
    var i int = 1

    for kr <= kn && !found {
        var mid int = (kr + kn) / 2

        fmt.Println("Perulangan ke - ", i)
        fmt.Println("Mid = ", mid)
        fmt.Println("Arr = ", arr[mid])
        fmt.Println("Found = ", arr[mid] == x)

        if arr[mid] < x {
            kr = mid + 1
        } else if arr[mid] > x {
            kn = mid - 1
        } else {
            found = true
            idx = mid
        }
        i++
    }

    if found {
        fmt.Println( idx)
    
    }
}