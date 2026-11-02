package main
import "fmt"
func main() {
	var n int
	r,t, luas, volume float64

	fmt.Scan(&n,&r,&t)

	luas = (2*3.14*r*(r+t))/(n) + 2*r*t
	volume = 3.14*r*r*t/n

	fmt.Println(luas, "cm2")
	fmt.Println(volume, "cm3")
}