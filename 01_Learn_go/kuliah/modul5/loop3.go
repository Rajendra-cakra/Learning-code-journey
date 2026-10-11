package main
import "fmt"
func main() {
	var n,j,x,jrk int
	fmt.Scan(&n)
	for j = 1; j <= n; j++{
		fmt.Scan(&x)
		jrk = x*5
		fmt.Println("jarak lari adalah: ", j)
		fmt.Println("jarak renang adalah: ", jrk)
	}
	fmt.Println("Selesai")
}