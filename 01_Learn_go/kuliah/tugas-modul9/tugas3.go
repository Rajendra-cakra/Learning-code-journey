package main
import "fmt"
func main() {
    var x, y int
    fmt.Scan(&x, &y)
    fmt.Println(y%x == 0)
    fmt.Println(x%y == 0)
}