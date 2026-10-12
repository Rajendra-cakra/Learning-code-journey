package main
import "fmt"
func main() {
	var a,b,j int

	fmt.Print("Masukkan Nilai A dan B: ")
	fmt.Scan(&a, &b)

	for j = a; j <= b; j+=1{
		fmt.Println(j, "")
	}
}