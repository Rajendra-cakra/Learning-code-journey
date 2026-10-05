package  main
import "fmt"

func main() {
	var nama string = "cakra"
	var umur int = 20
	var tinggi float64 = 1.65

	fmt.Println("masukkan nama, umur, dan tinggi badan anda: ")
	fmt.Scan(&nama, &umur, &tinggi)
	fmt.Println("Nama saya", nama, "umur saya", umur, "tahun dan tinggi badan saya", tinggi, "meter")
}