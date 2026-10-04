package main
import ("fmt")
func main() {
	var r int
	var volume, luas float64
	const phi float64 = 3.1415926535
	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scan(&r)
	volume = (1.333333333) * phi * float64(r*r*r)
	luas = 4 * phi * float64(r*r)
	fmt.Printf("Volume bola: %.4f\n", volume)
	fmt.Printf("Luas permukaan bola: %.4f\n", luas)
}