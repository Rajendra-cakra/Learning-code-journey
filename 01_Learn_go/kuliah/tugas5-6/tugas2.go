package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	prima := n > 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			prima = false
			break
		}
	}
	if prima {
		fmt.Println("Prima")
	} else {
		fmt.Println("Bukan Prima")
	}
}