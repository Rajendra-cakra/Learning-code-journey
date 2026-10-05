package main
	import "fmt"
func main() {
	var berat, tinggi, bmi float64
	fmt.Scan(&berat, &tinggi)
	bmi = berat / (tinggi * tinggi)
	fmt.Println(bmi)
}
