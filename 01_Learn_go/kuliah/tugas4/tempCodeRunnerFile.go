package main
import "fmt"
func main() {
	var total, diskon int
	fmt.Print("Masukkan Total Belanja Awal: Rp ")
	fmt.Scan(&total)
	fmt.Print("Masukkan Diskon Dalam Bentuk Persen: ")
	fmt.Scan(&diskon)
	fmt.Print("Total Belanja Setelah Diskon: Rp ", total-(total*diskon/100))
}