package main

import "fmt"

var year int

func main() {
	fmt.Scan(&year)
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println("Високосный")
	} else {
		fmt.Println("Не высокосный")
	}
}
