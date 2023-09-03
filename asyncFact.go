package main

import (
	"fmt"
	"sync"
)

// func factorialAsync(n int) string {

// 	result := 1
// 	for i := 1; i <= n; i++ {
// 		result *= i
// 	}
// 	return fmt.Sprintf("%d", result)
// }

func factorialAsync(n int, resultChan chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	resultChan <- Result{Number: n, Factorial: fmt.Sprintf("%d", result)}
}
