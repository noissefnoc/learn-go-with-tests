package tdd_in_go

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n - 1) + fibonacci(n - 2)
}




