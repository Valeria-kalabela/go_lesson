package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Print("Введите год: ")
	fmt.Fscan(os.Stdin, &n)
	if (n%4 == 0 && n%100 != 0) || (n%400 == 0) {
		fmt.Print("Год високосный")
	} else {
		fmt.Print("Год не високосный")
	}
}
