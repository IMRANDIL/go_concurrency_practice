package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// func main() {
// 	// Define the number for which we want to calculate the factorial.
// 	num := 950

// 	// Create a channel to collect results from goroutines.
// 	resultChan := make(chan int, num)

// 	// Create a WaitGroup to wait for all goroutines to finish.
// 	var wg sync.WaitGroup

// 	for i := 1; i <= num; i++ {
// 		// Increment the WaitGroup counter for each goroutine.
// 		wg.Add(1)

// 		// Launch a goroutine to calculate factorial.
// 		go factorial(i, &wg, resultChan)
// 	}

// 	// Close the result channel once all goroutines are done.
// 	go func() {
// 		wg.Wait()
// 		close(resultChan)
// 	}()

// 	// Collect results from the channel and print them.
// 	var finalResult int
// 	for result := range resultChan {
// 		finalResult = result
// 	}
// 	fmt.Printf("Factorial: %d\n", finalResult)
// }

// func main() {
// 	// Define the number for which we want to calculate the factorial.
// 	num := 3

// 	// Create a channel to collect results from goroutines.
// 	resultChan := make(chan int, num)

// 	// Create a WaitGroup to wait for all goroutines to finish.
// 	var wg sync.WaitGroup

// 	for i := 1; i <= num; i++ {
// 		// Increment the WaitGroup counter for each goroutine.
// 		wg.Add(1)

// 		// Launch a goroutine to calculate factorial.
// 		go factorial(i, &wg, resultChan)
// 	}

// 	// Wait for all goroutines to finish.
// 	wg.Wait()

// 	// Close the result channel.
// 	close(resultChan)

// 	// Collect and print results from the channel.
// 	var finalResult int
// 	for result := range resultChan {
// 		finalResult = result
// 	}

// 	fmt.Printf("Factorial: %d\n", finalResult)
// }

// func main() {
// 	// Define the number for which we want to calculate the factorial.
// 	num := 3

// 	// Create a channel to collect results from goroutines.
// 	resultChan := make(chan int, num)

// 	// Create a WaitGroup to wait for all goroutines to finish.
// 	var wg sync.WaitGroup

// 	var finalResult int // Variable to accumulate the final result.

// 	for i := 1; i <= num; i++ {
// 		// Increment the WaitGroup counter for each goroutine.
// 		wg.Add(1)

// 		// Launch a goroutine to calculate factorial.
// 		go factorial(i, &wg, resultChan)
// 	}

// 	// Create a goroutine to collect results and accumulate the final result.
// 	go func() {
// 		for result := range resultChan {
// 			finalResult = result
// 		}
// 	}()

// 	// Wait for all goroutines to finish.
// 	wg.Wait()

// 	// Close the result channel.
// 	close(resultChan)

// 	// Print the final accumulated result.
// 	fmt.Printf("Final Factorial: %d\n", finalResult)
// }

// func main() {
// 	// Define the number for which we want to calculate the factorial.
// 	num := 3

// 	// Create a channel to collect results from goroutines.
// 	resultChan := make(chan int, num)

// 	// Create a WaitGroup to wait for all goroutines to finish.
// 	var wg sync.WaitGroup

// 	var finalResult = 1 // Initialize the final result to 1.

// 	for i := 1; i <= num; i++ {
// 		// Increment the WaitGroup counter for each goroutine.
// 		wg.Add(1)

// 		// Launch a goroutine to calculate factorial.
// 		go factorial(i, &wg, resultChan)
// 	}

// 	// Create a goroutine to collect results and compute the final result.
// 	go func() {

// 		for result := range resultChan {
// 			finalResult *= result
// 		}
// 	}()

// 	// Wait for all goroutines to finish.
// 	wg.Wait()

// 	// Close the result channel.
// 	close(resultChan)

// 	// Print the final accumulated result.
// 	fmt.Printf("Final Result: %d\n", finalResult)
// }

// func main() {
// 	// Define the number for which we want to calculate the factorial.
// 	num := 3

// 	// Create a channel to collect results from goroutines.
// 	resultChan := make(chan int, num)

// 	// Create a WaitGroup to wait for all goroutines to finish.
// 	var wg sync.WaitGroup

// 	for i := 1; i <= num; i++ {
// 		// Increment the WaitGroup counter for each goroutine.
// 		wg.Add(1)

// 		// Launch a goroutine to calculate factorial.
// 		go func(i int) {
// 			defer wg.Done()
// 			factorial(i, &wg, resultChan)
// 		}(i)
// 	}

// 	// Wait for all goroutines to finish.
// 	wg.Wait()

// 	// Close the result channel once all goroutines are done.
// 	close(resultChan)

// 	finalResult := 1
// 	// Collect and multiply results from the channel.
// 	for result := range resultChan {
// 		finalResult *= result
// 	}

// 	// Print the final accumulated result.
// 	fmt.Printf("Final Result: %d\n", finalResult)
// }

// func main() {
// 	// Define the number for which we want to calculate the factorial.
// 	num := 6

// 	// Create a channel to collect results from goroutines.
// 	resultChan := make(chan *big.Int, num)

// 	// Create a WaitGroup to wait for all goroutines to finish.
// 	var wg sync.WaitGroup

// 	for i := 1; i <= num; i++ {
// 		// Increment the WaitGroup counter for each goroutine.
// 		wg.Add(1)

// 		// Launch a goroutine to calculate factorial.
// 		go factorial(i, &wg, resultChan)
// 	}

// 	// Close the result channel once all goroutines are done.
// 	go func() {
// 		wg.Wait()
// 		close(resultChan)
// 	}()

// 	// Collect results from the channel and print the last result.
// 	var finalResult *big.Int
// 	for result := range resultChan {
// 		finalResult = result
// 	}
// 	fmt.Printf("Factorial: %s\n", finalResult.String())
// }

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.The right way>>>>>>>>>>>

// func main() {
// 	// Define the number for which we want to calculate the factorial.
// 	num := 2890

// 	// Create a channel to collect the result.
// 	resultChan := make(chan *big.Int)

// 	// Launch a goroutine to calculate the factorial.
// 	go factorial(num, resultChan)

// 	// Collect the result from the channel.
// 	finalResult := <-resultChan

// 	// Print the result.
// 	fmt.Printf("Factorial of %d: %s\n", num, finalResult.String())
// }

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

type Result struct {
	Number    int    `json:"number"`
	Factorial string `json:"factorial"`
}

//synchronous one below

// func main() {
// 	num := 10 // Replace with the desired number.

// 	// Calculate factorial.
// 	factorialResult := factorialSync(num)

// 	// Store the result in a JSON file.
// 	result := Result{Number: num, Factorial: factorialResult}
// 	file, _ := os.Create("result.json")
// 	defer file.Close()
// 	encoder := json.NewEncoder(file)
// 	encoder.Encode(result)

// 	// Convert the JSON file to an Excel spreadsheet (not implemented here).
// 	fmt.Println("Synchronous execution completed.")
// }

//asynchroous way below

// func main() {
// 	num := 10 // Replace with the desired number.

// 	// Create a WaitGroup to wait for all goroutines to finish.
// 	var wg sync.WaitGroup

// 	// Calculate factorial concurrently.
// 	var factorialResult string
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		factorialResult = factorialAsync(num)
// 		fmt.Println(factorialResult)
// 	}()

// 	// Store the result in a JSON file concurrently.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		result := Result{Number: num, Factorial: factorialResult}
// 		file, _ := os.Create("resultAsync.json")
// 		defer file.Close()
// 		encoder := json.NewEncoder(file)
// 		encoder.Encode(result)
// 	}()

// 	// Convert the JSON file to an Excel spreadsheet (not implemented here).

// 	// Wait for all goroutines to finish.
// 	wg.Wait()

// 	fmt.Println("Asynchronous execution completed.")
// }

func main() {
	num := 10 // Replace with the desired number.

	// Create a WaitGroup to wait for all goroutines to finish.
	var wg sync.WaitGroup

	// Create a channel to collect results.
	resultChan := make(chan Result, 1)

	// Calculate factorial concurrently.
	wg.Add(1)
	go factorialAsync(num, resultChan, &wg)

	// Store the result in a JSON file concurrently.
	wg.Add(1)
	go func() {
		defer wg.Done()
		result := <-resultChan // Wait for the result from the channel.
		file, err := os.Create("resultAsync.json")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()
		encoder := json.NewEncoder(file)
		err = encoder.Encode(result)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}()

	// Convert the JSON file to an Excel spreadsheet (not implemented here).

	// Wait for all goroutines to finish.
	wg.Wait()

	fmt.Println("Asynchronous execution completed.")
}
