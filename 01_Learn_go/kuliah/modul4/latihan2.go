package main
import "fmt"
func main() {
	var bil,d1,d2,d3 int
	fmt.Scan(&bil)

	d1 = bil / 100
	d2 = bil % 100 / 10
	d3 = bil % 100 % 10
	membesar := d1 <= d2 && d2 <= d3

	fmt.Println(membesar)
}