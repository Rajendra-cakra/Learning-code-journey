package main
import "fmt"
func main() {
	var a, b int
	fmt.Print("Masukkan Angka dan Jumlah Pangkat: ")
	fmt.Scan(&a, &b)
	hasil := 1
	for i := 1; i <= b; i++ {
		hasil *= a
	}
	fmt.Println("Hasil", a, "pangkat", b, "adalah", hasil)
}
