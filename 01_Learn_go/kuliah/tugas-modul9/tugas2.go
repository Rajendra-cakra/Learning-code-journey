package main
import "fmt"
func main() {
    var bilangan int
    var genNeg bool
    fmt.Scan(&bilangan)
    genNeg = bilangan%2 == 0 && bilangan < 0
    if genNeg {
        fmt.Println("genap negatif")
    } else {
        fmt.Println("bukan")
    }
}