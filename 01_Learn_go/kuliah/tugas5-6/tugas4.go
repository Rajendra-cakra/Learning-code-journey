package main
import "fmt"
func main() {
	var n, x, total int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x%2 == 0 {
			total += x
		}
	}
	fmt.Println(total)
}
