package main
import "fmt"
type arrStr = [1234]string
func SeqSearch_1(T arrStr, n int, X string) bool {
	var found bool = false
	var j int = 0
	for j < n && !found {
		found = T[j] == X
		j = j + 1
	}
	return found
}
func SeqSearch_2(T arrStr, n int, X string) int {
	var found int = -1
	var j int = 0
	for j < n && found == -1 {
		if T[j] == X {
			found = j
		}
		j = j + 1
	}
	return found
}
func SeqSearch_3(T arrStr, n int, X string) int {
    var i int = 0
    for i < n && T[i] != X {
        i = i + 1
    }
    if i < n {
        return i
    } else {
        return -1
    }
}
func main() {
	var data arrStr 
	data[0] = "php"
	data[1] = "c++"
	data[2] = "golang"
	data[3] = "javascript"
	data[4] = "java"
	data[5] = "python"
	n := 6
	fmt.Println("Daftar Data:", data[:n])
	fmt.Println("================================")
	var target string
	fmt.Println("Masukkan bahasa pemrograman yang ingin dicari: ")
	fmt.Scanln(&target)
	fmt.Println("\n--------------------------------")
	fmt.Println("Hasil pencarian untuk data:", target)
	fmt.Println("SeqSearch_1 (Ketemu?):", SeqSearch_1(data, n, target))
	fmt.Println("SeqSearch_2 (Indeks):", SeqSearch_2(data, n, target))
	fmt.Println("SeqSearch_ 3 (Indeks):", SeqSearch_3(data, n, target))
}