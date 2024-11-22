package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	ans := []int{1}
	fmt.Print("Введите количество строк: ")
	fmt.Fscan(os.Stdin, &n)
	fmt.Println(ans[0])
	for i := 1; i < n; i++ {
		a := []int{}
		for j := 0; j < (i + 1); j++ {
			if ((j - 1) == -1) || (j == i) {
				a = append(a, ans[0])
			} else {
				a = append(a, ans[j-1]+ans[j])
			}
		}
		ans = a
		for j := 0; j < len(a); j++ {
			fmt.Print(a[j], " ")
		}
		fmt.Println()
	}
}
