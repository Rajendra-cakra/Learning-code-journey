package main
import "fmt"
func main() {
	var rupiah, dolar int

	fmt.Print("Masukkan jumlah rupiah: Rp.")
	fmt.Scan(&rupiah)
	dolar = rupiah / 15000
	fmt.Println("Jumlah rupiah dalam dolar adalah: $", dolar)
}