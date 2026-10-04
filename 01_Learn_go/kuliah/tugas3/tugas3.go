package main
import "fmt"
func main() {
	var tahun int
	var kabisat bool
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)
	kabisat = (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0)
	fmt.Println("Apakah tahun", tahun, "adalah tahun kabisat?", kabisat)
}