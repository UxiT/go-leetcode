package algos

func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}

	fibonacciMap := make(map[int]int, n)
	fibonacciMap[1] = 1
	fibonacciMap[2] = 2

	for i := 3; i <= n; i++ {
		fibonacciMap[i] = fibonacciMap[i-1] + fibonacciMap[i-2]
	}

	return fibonacciMap[n]
}
