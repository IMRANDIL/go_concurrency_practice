package main

import (
	"math/big"
)

// factorial calculates the factorial of a given number.
// func factorial(n int, wg *sync.WaitGroup, resultChan chan<- int) {
// 	defer wg.Done()

// 	result := 1
// 	for i := 1; i <= n; i++ {
// 		result *= i
// 	}

// 	// Send the result to the channel
// 	resultChan <- result
// }

// factorial calculates the factorial of a given number.
// func factorial(n int, wg *sync.WaitGroup, resultChan chan<- *big.Int) {
// 	defer wg.Done()

// 	result := new(big.Int).SetInt64(1)
// 	for i := 1; i <= n; i++ {
// 		result.Mul(result, big.NewInt(int64(i)))
// 	}

// 	// Send the result to the channel
// 	resultChan <- result
// }

func factorial(n int, resultChan chan<- *big.Int) {
	result := new(big.Int).SetInt64(1)
	for i := 1; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	resultChan <- result
}
