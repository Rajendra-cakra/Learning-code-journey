package main

import "fmt"

func main() {
	var (
		nama, nim, kelas string
	)
	fmt.Print("Masukan Nama Anda: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukan Nim Anda: ")
	fmt.Scanln(&nim)
	fmt.Print("Masukan Kelas Anda: ")
	fmt.Scanln(&kelas)

	fmt.Println("Halo, nama saya " + nama + ", dari kelas " + kelas + ", dan nim saya adalah " + nim)
}
