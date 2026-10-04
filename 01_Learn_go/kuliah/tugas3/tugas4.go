package main
import "fmt"
func main() {
	var celsius, reamur, fahrenheit, kelvin float64
	fmt.Print("Masukkan temperatur Celsius: ")
	fmt.Scan(&celsius)
	reamur = celsius * 4.0 / 5.0
	fahrenheit = celsius*9.0/5.0 + 32
	kelvin = (fahrenheit + 459.67) * 5.0 / 9.0
	fmt.Printf("Derajat Reamur: %.0f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)
}