package main
import "fmt"
func main() {
	var bmi, tinggi, berat float64
	fmt.Print("Masukkan BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan Tinggi Badan Dalam Meter: ")
	fmt.Scan(&tinggi)
	berat = bmi * tinggi * tinggi
	fmt.Printf("Berat Badan: %.0f kg", berat)
}