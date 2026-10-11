package main
import "fmt"
func main() {
	var n,j,x,jrk int
	fmt.Print("Masukkan Nilai N: ")
	fmt.Scan(&n)
	for j = 1; j <= n; j++{
		fmt.Print("Masukkan Nilai Berapa Meter: ")
		fmt.Scan(&x)
		jrk = x*5
		fmt.Println("jarak renang adalah: ", jrk)
	}
	fmt.Println("Selesai")
}