package main
	
import "fmt"

func main() {
	var d1,d2,d3,d4,m,k int
	fmt.Scan(&m)
	d1 = m / 1000
	d2 = (m % 1000) / 100
	d3 = (m % 100) / 10
	d4 = m % 10

	k = (d1+d4)*(d2-d3)+((d1-d4)*(d2+d3))
	fmt.Println(k)

}