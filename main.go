package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {
	nums := []int{4, 1, 2, 1, 2}

	r := easy.SingleNumber(nums)

	fmt.Print(r)
}
