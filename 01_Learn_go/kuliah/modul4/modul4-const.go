package main

import "fmt"

const pi float64 = 3.14

func main() {
	var r int = 10
	var luas float64
	pi = 1000
	luas = pi * float64(r*r)
	fmt.Println(luas, pi)
}