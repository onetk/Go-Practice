package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fib(n int) int {
	var summ int

	if n > 0 {
		fibo := []int{}
		fibo = append(fibo, 0)
		fibo = append(fibo, 1)

		// make fibonacci array
		for i := 1; i < n; i++ {
			fibo = append(fibo, fibo[i]+fibo[i-1])
		}
		fmt.Println(fibo)

		// calculate fibonacci sum
		summ = 0
		for array_i := 0; array_i < len(fibo); array_i++ {
			summ = summ + fibo[array_i]
		}
	} else {
		summ = 0
	}

	return summ
}

func main() {

	for {
		fmt.Print(">>")
		// Scannerを使って一行読み
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		if input.Text() == "q" {
			break
		}

		texts, _ := strconv.Atoi(input.Text())

		fmt.Println(fib(texts))
	}
}
