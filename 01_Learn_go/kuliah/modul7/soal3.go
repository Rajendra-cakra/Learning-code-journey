package main

import "fmt"

func main() {
	var dinar, qirat, dirham, fals int
	fmt.Scan(&qirat)

	fals = qirat / 6
	dirham = qirat/6 / 10
	dinar = qirat/6/10 / 10

	

	
	


	fmt.Println(fals)
	fmt.Println(dirham)
	fmt.Println(dinar)

}
