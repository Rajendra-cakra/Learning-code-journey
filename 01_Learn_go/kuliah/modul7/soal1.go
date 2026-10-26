package main

import "fmt"

func main() {
	var x, y, hasil int
	fmt.Scan(&x, &y)

	for i := x; i <= y; i++ {
		hasil += i
	}
	fmt.Println(hasil)

}
