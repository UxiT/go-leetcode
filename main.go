package main

import (
	"fmt"
	"leetcode/medium"
)

func main() {
	v1 := "1.0"
	v2 := "1.0.0.0"

	r := medium.CompareVersion(v1, v2)

	fmt.Print(r)
}
