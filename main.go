package main

import (
	"encoding/json"
	"fmt"
	"math/big"
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

// func main() {
// 	num := 398 // Replace with the desired number.

// 	// Create a WaitGroup to wait for all goroutines to finish.
// 	var wg sync.WaitGroup

// 	// Create a channel to collect results.
// 	resultChan := make(chan *big.Int)

// 	//create a channel to collect file read result

// 	fileResultChan := make(chan Result)
// 	// Calculate factorial concurrently.
// 	wg.Add(1)
// 	go factorialAsync(num, resultChan, &wg)

// 	// Store the result in a JSON file concurrently.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		factorialResult := <-resultChan // Wait for the result from the channel.
// 		result := Result{Number: num, Factorial: factorialResult.String()}
// 		file, err := os.Create("resultAsync.json")
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return
// 		}
// 		defer file.Close()
// 		encoder := json.NewEncoder(file)
// 		err = encoder.Encode(result)
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return
// 		}
// 	}()

// 	// Convert the JSON file to an Excel spreadsheet (not implemented here).
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		readJSONFile("resultAsync.json", fileResultChan, &wg)
// 	}()

// 	// Wait for all goroutines to finish.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		defer close(fileResultChan)
// 		generateExcelAsync(<-fileResultChan, &wg)
// 	}()

// 	wg.Wait()

// 	fmt.Println("Asynchronous execution completed.")
// }

// func main() {
// 	num := 398 // Replace with the desired number.

// 	// Create a WaitGroup to wait for all goroutines to finish.
// 	var wg sync.WaitGroup

// 	// Create a channel to collect results.
// 	resultChan := make(chan *big.Int)

// 	// Create a channel to collect file read result.
// 	fileResultChan := make(chan Result)

// 	// Calculate factorial concurrently.
// 	wg.Add(1)
// 	go factorialAsync(num, resultChan, &wg)

// 	// Store the result in a JSON file concurrently.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		factorialResult := <-resultChan // Wait for the result from the channel.
// 		result := Result{Number: num, Factorial: factorialResult.String()}
// 		file, err := os.Create("resultAsync.json")
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return
// 		}
// 		defer file.Close()
// 		encoder := json.NewEncoder(file)
// 		err = encoder.Encode(result)
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return
// 		}
// 	}()

// 	// Convert the JSON file to an Excel spreadsheet.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		readJSONFile("resultAsync.json", fileResultChan, &wg)
// 	}()

// 	// Wait for all goroutines to finish.
// 	wg.Wait()

// 	fmt.Println("Asynchronous execution completed.")
// }

// func main() {
// 	num := 398 // Replace with the desired number.

// 	// Create a WaitGroup to wait for all goroutines to finish.
// 	var wg sync.WaitGroup

// 	// Create a channel to collect results.
// 	resultChan := make(chan *big.Int)

// 	// Create a channel to collect file read result.
// 	fileResultChan := make(chan Result)

// 	// Calculate factorial concurrently.
// 	wg.Add(1)
// 	go factorialAsync(num, resultChan, &wg)

// 	// Store the result in a JSON file concurrently.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		factorialResult := <-resultChan // Wait for the result from the channel.
// 		result := Result{Number: num, Factorial: factorialResult.String()}
// 		file, err := os.Create("resultAsync.json")
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return
// 		}
// 		defer file.Close()
// 		encoder := json.NewEncoder(file)
// 		err = encoder.Encode(result)
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return
// 		}
// 	}()

// 	// Read the JSON file concurrently.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		readJSONFile("resultAsync.json", fileResultChan, &wg)
// 	}()

// 	// Convert the JSON file to an Excel spreadsheet.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		resultData := <-fileResultChan
// 		generateExcelAsync(resultData, &wg)
// 	}()

// 	// Wait for all goroutines to finish.
// 	wg.Wait()

// 	fmt.Println("Asynchronous execution completed.")
// }

// func main() {
// 	num := 398 // Replace with the desired number.

// 	// Create a WaitGroup to wait for all goroutines to finish.
// 	var wg sync.WaitGroup

// 	// Create a channel to collect results.
// 	resultChan := make(chan *big.Int)
// 	resultDataChan := make(chan Result)
// 	// Create a channel to collect file read result.
// 	fileResultChan := make(chan Result)

// 	// Calculate factorial concurrently.
// 	wg.Add(1)
// 	go factorialAsync(num, resultChan, &wg)

// 	// Store the result in a JSON file concurrently.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		factorialResult := <-resultChan // Wait for the result from the channel.
// 		result := Result{Number: num, Factorial: factorialResult.String()}
// 		file, err := os.Create("resultAsync.json")
// 		if err != nil {
// 			fmt.Println("Error creating JSON file:", err)
// 			return
// 		}
// 		defer file.Close()
// 		encoder := json.NewEncoder(file)
// 		if err = encoder.Encode(result); err != nil {
// 			fmt.Println("Error encoding JSON:", err)
// 			return
// 		}
// 		fmt.Println("JSON file saved successfully.")
// 	}()

// 	// Convert the JSON file to an Excel spreadsheet.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		readJSONFile("resultAsync.json", fileResultChan, &wg)

// 	}()
// 	//fmt.Println(<-fileResultChan)

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		resultData := <-fileResultChan
// 		resultDataChan <- resultData // Pass the data to the second goroutine
// 	}()
// 	// Generate the Excel file concurrently.
// 	wg.Add(1)
// 	go func() {
// 		fmt.Println("generateExcelAsync started")
// 		defer wg.Done()
// 		fmt.Println("generateExcelAsync started after wg.done")
// 		resultFactorial := <-resultDataChan
// 		fmt.Println("generateExcelAsync started after the extraction the result from channel")
// 		requiredResult := Result{Number: num, Factorial: resultFactorial.Factorial}
// 		err := generateExcelAsync(requiredResult, &wg)
// 		if err != nil {
// 			fmt.Println("Error generating Excel file:", err)
// 			return
// 		}
// 		fmt.Println("Excel file generated successfully.")
// 	}()

// 	// Wait for all goroutines to finish.
// 	wg.Wait()

// 	fmt.Println("Asynchronous execution completed.")
// }

// func main() {
// 	num := 398 // Replace with the desired number.

// 	// Create a WaitGroup to wait for all goroutines to finish.
// 	var wg sync.WaitGroup

// 	// Create channels to collect results.
// 	resultChan := make(chan *big.Int)
// 	//resultDataChan := make(chan Result)
// 	fileResultChan := make(chan Result)

// 	// Calculate factorial concurrently.
// 	wg.Add(1)
// 	go factorialAsync(num, resultChan, &wg)

// 	// Store the result in a JSON file concurrently.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		factorialResult := <-resultChan
// 		result := Result{Number: num, Factorial: factorialResult.String()}
// 		file, err := os.Create("resultAsync.json")
// 		if err != nil {
// 			fmt.Println("Error creating JSON file:", err)
// 			return
// 		}
// 		defer file.Close()
// 		encoder := json.NewEncoder(file)
// 		if err = encoder.Encode(result); err != nil {
// 			fmt.Println("Error encoding JSON:", err)
// 			return
// 		}
// 		fmt.Println("JSON file saved successfully.")

// 	}()

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		readJSONFile("resultAsync.json", fileResultChan, &wg)

// 	}()

// 	// Create a goroutine to receive and print values from fileResultChan.
// 	// wg.Add(1)
// 	// go func() {
// 	// 	defer wg.Done()
// 	// 	for result := range fileResultChan {
// 	// 		fmt.Println(result)
// 	// 		close(fileResultChan)
// 	// 	}
// 	// }()

// 	// Convert the JSON file to an Excel spreadsheet.
// 	// wg.Add(1)
// 	// go func() {
// 	// 	defer wg.Done()
// 	// 	resultData := <-fileResultChan
// 	// 	resultDataChan <- resultData // Pass the data to the second goroutine
// 	// 	close(resultDataChan)        // Close the channel when done
// 	// }()

// 	// Generate the Excel file concurrently.
// 	wg.Add(1)
// 	go func() {

// 		defer wg.Done()

// 		//resultFactorial := <-resultDataChan

// 		//requiredResult := Result{Number: num, Factorial: resultFactorial.Factorial}
// 		//fmt.Println(requiredResult)
// 		err := generateExcelAsync(fileResultChan, &wg)

// 		if err != nil {
// 			fmt.Print("Error generating Excel file:", err)
// 			return
// 		}
// 		fmt.Println("Excel file generated successfully.")
// 	}()

// 	// Wait for all goroutines to finish.
// 	wg.Wait()

// 	fmt.Println("Asynchronous execution completed.")
// }

// func main() {
// 	num := 8 // Replace with the desired number.

// 	// Create a WaitGroup to wait for all goroutines to finish.
// 	var wg sync.WaitGroup

// 	// Create channels to collect results.
// 	resultChan := make(chan *big.Int)
// 	fileResultChan := make(chan Result)
// 	excelDone := make(chan bool) // Channel to signal Excel generation completion

// 	// Calculate factorial concurrently.
// 	wg.Add(1)
// 	go factorialAsync(num, resultChan, &wg)

// 	// Store the result in a JSON file concurrently.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		factorialResult := <-resultChan
// 		result := Result{Number: num, Factorial: factorialResult.String()}
// 		file, err := os.Create("resultAsync.json")
// 		if err != nil {
// 			fmt.Println("Error creating JSON file:", err)
// 			return
// 		}
// 		defer file.Close()
// 		encoder := json.NewEncoder(file)
// 		if err = encoder.Encode(result); err != nil {
// 			fmt.Println("Error encoding JSON:", err)
// 			return
// 		}
// 		fmt.Println("JSON file saved successfully.")

// 	}()

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		readJSONFile("resultAsync.json", fileResultChan, &wg)

// 	}()

// 	// Generate the Excel file concurrently.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		err := generateExcelAsync(fileResultChan, excelDone)
// 		if err != nil {
// 			fmt.Print("Error generating Excel file:", err)
// 			return
// 		}
// 		fmt.Println("Excel file generated successfully.")

// 	}()

// 	// Wait for Excel generation to complete.
// 	<-excelDone

// 	// Close the fileResultChan channel after Excel generation is complete.
// 	//close(fileResultChan)

// 	// Wait for all goroutines to finish.
// 	wg.Wait()

// 	fmt.Println("Asynchronous execution completed.")
// }

// func main() {
// 	num := 398 // Replace with the desired number.

// 	// Create a WaitGroup to wait for all goroutines to finish.
// 	var wg sync.WaitGroup

// 	// Create channels to collect results.
// 	resultChan := make(chan *big.Int)

// 	// Calculate factorial concurrently.
// 	wg.Add(1)
// 	go factorialAsync(num, resultChan, &wg)

// 	// Store the result in a JSON file concurrently.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		factorialResult := <-resultChan
// 		result := Result{Number: num, Factorial: factorialResult.String()}
// 		file, err := os.Create("resultAsync.json")
// 		if err != nil {
// 			fmt.Println("Error creating JSON file:", err)
// 			return
// 		}
// 		defer file.Close()
// 		encoder := json.NewEncoder(file)
// 		if err = encoder.Encode(result); err != nil {
// 			fmt.Println("Error encoding JSON:", err)
// 			return
// 		}
// 		fmt.Println("JSON file saved successfully.")
// 	}()

// 	// Create a channel for reading JSON and populating fileResultChan.
// 	fileResultChan := make(chan Result)
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		readJSONFile("resultAsync.json", fileResultChan, &wg)
// 		close(fileResultChan)
// 	}()

// 	// Generate the Excel file concurrently.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		err := generateExcelAsync(fileResultChan, &wg)
// 		if err != nil {
// 			fmt.Print("Error generating Excel file:", err)
// 			return
// 		}
// 		fmt.Println("Excel file generated successfully.")
// 	}()

// 	// Wait for all goroutines to finish.
// 	wg.Wait()

// 	fmt.Println("Asynchronous execution completed.")
// }

// func main() {
// 	num := 398 // Replace with the desired number.

// 	// Create a WaitGroup to wait for all goroutines to finish.
// 	var wg sync.WaitGroup

// 	// Create channels to collect results.
// 	resultChan := make(chan *big.Int)

// 	// Calculate factorial concurrently.
// 	wg.Add(1)
// 	go factorialAsync(num, resultChan, &wg)

// 	// Store the result in a JSON file concurrently.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		factorialResult := <-resultChan
// 		result := Result{Number: num, Factorial: factorialResult.String()}
// 		file, err := os.Create("resultAsync.json")
// 		if err != nil {
// 			fmt.Println("Error creating JSON file:", err)
// 			return
// 		}
// 		defer file.Close()
// 		encoder := json.NewEncoder(file)
// 		if err = encoder.Encode(result); err != nil {
// 			fmt.Println("Error encoding JSON:", err)
// 			return
// 		}
// 		fmt.Println("JSON file saved successfully.")
// 	}()

// 	// Create a channel for reading JSON and populating fileResultChan.
// 	fileResultChan := make(chan Result)

// 	// Use a separate WaitGroup for the JSON reading goroutine.
// 	var jsonWg sync.WaitGroup
// 	jsonWg.Add(1)
// 	go func() {
// 		defer jsonWg.Done()
// 		readJSONFile("resultAsync.json", fileResultChan, &jsonWg)
// 		close(fileResultChan)
// 	}()

// 	// Wait for the JSON reading to complete.
// 	jsonWg.Wait()

// 	// Generate the Excel file concurrently, waiting for the JSON reading to complete.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		// Wait for the JSON reading to complete before generating the Excel file.
// 		jsonWg.Wait()
// 		err := generateExcelAsync(fileResultChan, &wg)
// 		if err != nil {
// 			fmt.Print("Error generating Excel file:", err)
// 			return
// 		}
// 		fmt.Println("Excel file generated successfully.")
// 	}()

// 	// Wait for all goroutines to finish.
// 	wg.Wait()

// 	fmt.Println("Asynchronous execution completed.")
// }

// func main() {
// 	num := 398 // Replace with the desired number.

// 	// Create a WaitGroup to wait for all goroutines to finish.
// 	var wg sync.WaitGroup

// 	// Create channels to collect results.
// 	resultChan := make(chan *big.Int)

// 	// Calculate factorial concurrently.
// 	wg.Add(1)
// 	go factorialAsync(num, resultChan, &wg)

// 	// Store the result in a JSON file concurrently.
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		factorialResult := <-resultChan
// 		result := Result{Number: num, Factorial: factorialResult.String()}
// 		file, err := os.Create("resultAsync.json")
// 		if err != nil {
// 			fmt.Println("Error creating JSON file:", err)
// 			return
// 		}
// 		defer file.Close()
// 		encoder := json.NewEncoder(file)
// 		if err = encoder.Encode(result); err != nil {
// 			fmt.Println("Error encoding JSON:", err)
// 			return
// 		}
// 		fmt.Println("JSON file saved successfully.")
// 	}()

// 	// Create a channel for reading JSON and populating fileResultChan.
// 	fileResultChan := make(chan Result)
// 	var excelWg sync.WaitGroup
// 	// Use a separate WaitGroup for the JSON reading goroutine.
// 	var jsonWg sync.WaitGroup
// 	jsonWg.Add(1)
// 	go func() {
// 		defer jsonWg.Done()
// 		excelWg.Wait()
// 		readJSONFile("resultAsync.json", fileResultChan, &jsonWg)
// 		close(fileResultChan)
// 	}()

// 	// Create a WaitGroup for the Excel generation goroutine.

// 	// Generate the Excel file concurrently.
// 	excelWg.Add(1)
// 	go func() {
// 		defer excelWg.Done()
// 		// Wait for the JSON reading to complete before generating the Excel file.
// 		jsonWg.Wait()
// 		err := generateExcelAsync(fileResultChan, &excelWg)
// 		if err != nil {
// 			fmt.Print("Error generating Excel file:", err)
// 			return
// 		}
// 		fmt.Println("Excel file generated successfully.")
// 	}()

// 	// Wait for all goroutines to finish.
// 	wg.Wait()

// 	// Wait for the Excel generation to complete.
// 	//excelWg.Wait()

// 	fmt.Println("Asynchronous execution completed.")
// }

func main() {
	num := 8 // Replace with the desired number.

	// Create a WaitGroup to wait for all goroutines to finish.
	var wg sync.WaitGroup

	// Create channels to collect results.
	resultChan := make(chan *big.Int)
	fileResultChan := make(chan Result)
	excelDone := make(chan bool) // Channel to signal Excel generation completion

	// Calculate factorial concurrently.
	wg.Add(1)
	go factorialAsync(num, resultChan)

	// Store the result in a JSON file concurrently.
	wg.Add(1)
	go func() {
		defer wg.Done()
		factorialResult := <-resultChan
		result := Result{Number: num, Factorial: factorialResult.String()}
		file, err := os.Create("resultAsync.json")
		if err != nil {
			fmt.Println("Error creating JSON file:", err)
			return
		}
		defer file.Close()
		encoder := json.NewEncoder(file)
		if err = encoder.Encode(result); err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
		fmt.Println("JSON file saved successfully.")

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		readJSONFile("resultAsync.json", fileResultChan, &wg)
		close(fileResultChan)
	}()

	// Generate the Excel file concurrently.
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := generateExcelAsync(fileResultChan, excelDone)
		if err != nil {
			fmt.Print("Error generating Excel file:", err)
			return
		}
		fmt.Println("Excel file generated successfully.")

	}()

	// Wait for Excel generation to complete.
	<-excelDone

	// Wait for all goroutines to finish.
	wg.Wait()

	fmt.Println("Asynchronous execution completed.")
}
