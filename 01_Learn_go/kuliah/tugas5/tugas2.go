package main
import (
	"fmt"
	"math"
)
func main() {
	var n int
	var r, t float64
	fmt.Print("Masukkan jumlah perulangan: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Print("Masukkan jari-jari dan tinggi kerucut: ")
		fmt.Scan(&r, &t)
		volume := math.Pi * r * r * t / 3
		fmt.Println("Volume kerucut:", volume)
	}
}
