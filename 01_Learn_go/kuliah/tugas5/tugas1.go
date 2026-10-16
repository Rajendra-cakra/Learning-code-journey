package main
import "fmt"
func main() {
	var n, total int
	fmt.Print("Masukkan Nilai N: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		total += i
	}
	fmt.Println("Hasil Penjumlahan Berulang dari 1 sampai", n, "adalah", total)
}
