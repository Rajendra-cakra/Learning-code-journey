package main

import "fmt"

func main() {
	var alas, tinggi, luas float64
	fmt.Print("Masukkan alas dan tinggi: ")
	fmt.Scan(&alas, &tinggi)
	
	luas = 0.5 * alas * tinggi
	fmt.Println("Luas segitiga adalah : ", luas)

}