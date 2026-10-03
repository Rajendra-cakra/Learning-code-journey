package main
import (
	"fmt"
	"math"
)
func main() {
	var fx float64
	fmt.Scan(&fx)

	x := 2/(fx-5) - 5
	fmt.Println(int(math.Round(x)))
}