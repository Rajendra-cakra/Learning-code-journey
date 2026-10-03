package main
import "fmt"
func main() {
	var siswa, grade, status string
	var nilai1, nilai2 float64
	fmt.Print("Masukkan nama anda: ")
	fmt.Scan(&siswa)
	fmt.Print("Masukkan nilai ujian 1 anda: ")
	fmt.Scan(&nilai1)
	fmt.Print("Masukkan nilai ujian 2 anda: ")
	fmt.Scan(&nilai2)
	ratarata := (nilai1 + nilai2) / 2
	if ratarata >= 85 {
		grade = "A"
		status = "LULUS"
	} else if ratarata >= 70 {
		grade = "B"
		status = "LULUS"
	} else if ratarata >= 55 {
		grade = "C"
		status = "LULUS"
	} else {
		grade = "D"
		status = "TIDAK LULLUS"
	}
	fmt.Println("Siswa", siswa, "dinyatakan", status, "dengan Grade", grade)
}
