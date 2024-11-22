package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Введите строку: ")
	var c float64
	n, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	mas := strings.Fields(n)
	for i := 1; i < len(mas)-1; i++ {
		n1, err1 := strconv.Atoi(mas[i-1])
		n2, err2 := strconv.Atoi(mas[i+1])
		if (err1 != nil) || (err2 != nil) {
			fmt.Println(err1)
		}
		switch mas[i] {
		case "+":
			c = float64(n1) + float64(n2)
			mas[i+1] = strconv.FormatFloat(c, 'g', -1, 64)
		case "-":
			c = float64(n1) - float64(n2)
			mas[i+1] = strconv.FormatFloat(c, 'g', -1, 64)
		case "/":
			c = float64(n1) / float64(n2)
			mas[i+1] = strconv.FormatFloat(c, 'g', -1, 64)
		case "*":
			c = float64(n1) * float64(n2)
			mas[i+1] = strconv.FormatFloat(c, 'g', -1, 64)
		case "^":
			c = math.Pow(float64(n1), float64(n2))
			mas[i+1] = strconv.FormatFloat(c, 'g', -1, 64)
		default:
			fmt.Print()
		}
	}
	fmt.Print(mas)
}
