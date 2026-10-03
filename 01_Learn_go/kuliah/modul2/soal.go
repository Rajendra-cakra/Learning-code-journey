package main

import "fmt"

func main() {
	var bilangan int
	var status bool
	fmt.Scan(&bilangan)
	status = bilangan >= 0
	fmt.Println(status)
}
