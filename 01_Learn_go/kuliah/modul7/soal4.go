package main

import "fmt"

func main() {
	var totalqirat int
	fmt.Scan(&totalqirat)

	dinar := totalqirat / 600
	totalqirat %= 600
	dirham := totalqirat / 60
	totalqirat %= 60
	fals := totalqirat  /6
	qirat := totalqirat % 6

	fmt.Print(dinar, dirham, fals, qirat)

}
