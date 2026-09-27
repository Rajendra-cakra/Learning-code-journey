package main
import "fmt"

func main() {
	var karakter rune
	var teks string

	karakter = 'A'
	teks = "A"

	fmt.Println("karakter sebagai angka :", karakter)
	fmt.Println("karakter sebagai huruf :", string(karakter))
	fmt.Println("string		:", teks)
}