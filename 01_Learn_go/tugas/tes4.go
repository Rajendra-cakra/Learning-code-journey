package main

import "fmt"

func main() {
	fmt.Println("INI ADALAH PROGRAM KONVERSI SUHU DARI FAHRENHEIT KE CELCIUS")
	var fahrenheit int
	fmt.Print("Masukan suhu dalam fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("Suhu dalam satuan Celcius adalah: %d\n", celsius)
}
