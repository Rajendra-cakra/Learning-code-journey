package main
import "fmt"
func main() {
	var j, alas, tinggi, n int
	var luas float64

	fmt.Scan(&n)

	for j = 1; j <= n; j++{

		fmt.Scan(&alas, &tinggi)

		luas = 0.5 * float64(alas * tinggi)
		fmt.Println("Luas Segitiga adalah: ", luas)
	}


}