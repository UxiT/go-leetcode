package main

import "fmt"

func main() {
	k := 2
	n := 2
	res := getFactorial(n) / (getFactorial(k) * getFactorial(n-k))

	fmt.Printf("res: %d", res)
}

func getFactorial(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	return res
}
