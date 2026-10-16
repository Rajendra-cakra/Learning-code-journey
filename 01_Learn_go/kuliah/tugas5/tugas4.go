package main
import "fmt"
func main() {
	var n int
	fmt.Print("Masukkan Bilangan faktorial: ")
	fmt.Scan(&n)
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	fmt.Println("Faktorial dari", n, "adalah", hasil)
}
