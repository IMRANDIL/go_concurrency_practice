package main

import "fmt"

func factorialSync(n int) string {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return fmt.Sprintf("%d", result)
}
