package main

import (
	"math/big"
	"sync"
)

// func factorialAsync(n int) string {

// 	result := 1
// 	for i := 1; i <= n; i++ {
// 		result *= i
// 	}
// 	return fmt.Sprintf("%d", result)
// }

func factorialAsync(n int, resultChan chan<- *big.Int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := new(big.Int).SetInt64(1)
	for i := 1; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	resultChan <- result
}
