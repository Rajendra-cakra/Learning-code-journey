package main

import "fmt"

func main() {
	var sisi int
	var volume int

	fmt.Print("Masukkan sisi kubus : ")
	fmt.Scan(&sisi)

	volume = sisi * sisi * sisi
	fmt.Println("Volume kubus adalah : ", volume)
}
