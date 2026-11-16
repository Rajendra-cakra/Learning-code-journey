package main

import "fmt"

func main() {
	var nilai int
	fmt.Scan(&nilai)
	if nilai >= 75 {
		fmt.Println("A")
	} else if nilai >= 65 {
		fmt.Println("B")
	} else if nilai >= 50 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}
}

// if then cuman bisa 1 kondisi.
// misal if benar, maka kebawah ga dijalankan
// misal if salah, maka akaan lanjut ke else if, jika else if kedua benar, maka kebawah ga dijalankan
// misal if salah, maka akan lanjut ke else if, jika salah lagi maka akan lanjut ke else
