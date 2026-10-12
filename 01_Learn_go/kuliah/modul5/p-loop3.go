package main

import "fmt"

func main() {
	var n1, n2, hasil int
	fmt.Scan(&n1, &n2)

	for j := 1; j <= n2; j++ {
		hasil += n1
		fmt.Println(n1, "+ ")
	}

	fmt.Println("\n=",hasil)

}