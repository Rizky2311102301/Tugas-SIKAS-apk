package main

import "fmt"

func main() {
    var arr [10]int = [10]int{1, 8, 11, 25, 26, 32, 52, 77, 90, 93}
    var kr int = 0
    var kn int = len(arr) - 1 
    var found bool = false
    var x int = 100
    var idx int = -1
    var i int = 0 

    for kr <= kn && !found {
        var mid int = (kr + kn) / 2
        fmt.Println("Perulangan ke -", i)
        fmt.Println("kr =", kr, "kn =", kn)  
        fmt.Println("Mid =", mid)
        fmt.Println("Arr =", arr[mid])
        fmt.Println("Found =", arr[mid] == x)
        
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
    
    if idx != -1 {
        fmt.Println("Elemen ditemukan di indeks:", idx)
    } else {
        fmt.Println("Elemen tidak ditemukan")
    }
}

    
