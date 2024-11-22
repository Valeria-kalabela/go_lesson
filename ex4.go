package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var a string
	var ans string = ""
	var b string = ""
	fmt.Print("Введите строку: ")
	n, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	mas := strings.Fields(n)
	for i := 0; i < len(mas[0]); i++ {
		a = string((mas[0])[i])
		b = ""
		for j := 1; j < len(mas); j++ {
			if (len(mas[j]) - 1) >= i {
				b = b + string((mas[j])[i])
			}
		}
		if (strings.Repeat(a, len(mas)-1)) == b {
			ans = ans + a
		} else {
			break
		}
	}
	fmt.Print("Ответ: ", ans)
}
