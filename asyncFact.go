package main

import (
	"math/big"
)

// func factorialAsync(n int) string {

// 	result := 1
// 	for i := 1; i <= n; i++ {
// 		result *= i
// 	}
// 	return fmt.Sprintf("%d", result)
// }

func factorialAsync(n int, resultChan chan<- *big.Int) {

	result := new(big.Int).SetInt64(1)
	for i := 1; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	resultChan <- result
}
