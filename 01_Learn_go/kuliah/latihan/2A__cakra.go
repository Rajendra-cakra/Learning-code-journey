// Rajendra Cakra Adji Nugraha
// 109082600052
// 2A
// IF-14-04

package main

import "fmt"

func main() {
	var n, r, t float64
	fmt.Print("masukkan nilai n, r, t: ")
	fmt.Scan(&n, &r, &t)

	luas := (2*3.14*r*(r+t))/(n) + 2*r*t
	volume := 3.14 * r * r * t / n

	fmt.Print(luas, " cm2 ")
	fmt.Println(volume, "cm3")
}
