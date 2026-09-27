package main
import "fmt"
func main() {
	var r float64
	const pi = 3.14
	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scan(&r)
	luas := pi * r * r
	fmt.Printf("Luas Lingkaran dengan jari jari %.1f adalah: %.1f\n", r, luas)
}
