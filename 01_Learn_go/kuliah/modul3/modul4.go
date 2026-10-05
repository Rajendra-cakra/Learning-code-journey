package  main
import "fmt"

type bulat int
type desimal float64
type huruf string

func main() {
	var nama huruf = "cakra"
	var umur bulat = 20
	var tinggi desimal = 1.75

	fmt.Println("masukkan nama, umur, dan tinggi badan anda: ")
	fmt.Scan(&nama, &umur, &tinggi)
	fmt.Println("Nama saya", nama, "umur saya", umur, "tahun dan tinggi badan saya", tinggi, "meter")
}