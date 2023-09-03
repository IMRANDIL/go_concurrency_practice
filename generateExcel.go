package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// func readJSONFile(filename string, fileResultChan chan<- Result) {

// 	// Open the JSON file for reading.
// 	file, err := os.Open(filename)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()

// 	// Create a variable to hold the decoded JSON data.
// 	var resultData Result

// 	// Create a JSON decoder and decode the file into the resultData variable.
// 	decoder := json.NewDecoder(file)
// 	if err := decoder.Decode(&resultData); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fileResultChan <- resultData
// 	close(fileResultChan)
// }

// func readJSONFile(filename string, fileResultChan chan<- Result) {
// 	// Open the JSON file for reading.
// 	file, err := os.Open(filename)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()

// 	// Create a variable to hold the decoded JSON data.
// 	var resultData Result

// 	// Create a JSON decoder and decode the file into the resultData variable.
// 	decoder := json.NewDecoder(file)
// 	if err := decoder.Decode(&resultData); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fileResultChan <- resultData
// }

// func readJSONFile(filename string, fileResultChan chan<- Result) {
// 	// Open the JSON file for reading.
// 	file, err := os.Open(filename)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()

// 	// Create a variable to hold the decoded JSON data.
// 	var resultData Result

// 	// Create a JSON decoder and decode the file into the resultData variable.
// 	decoder := json.NewDecoder(file)
// 	if err := decoder.Decode(&resultData); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fileResultChan <- resultData
// }

// func readJSONFile(filename string, result Result, excelDone chan<- bool) {
// 	// Open the JSON file for reading.
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()

// 	// Create a variable to hold the decoded JSON data.
// 	var resultData Result

// 	// Create a JSON decoder and decode the file into the resultData variable.
// 	decoder := json.NewDecoder(file)
// 	if err := decoder.Decode(&resultData); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	if result != resultData {
// 		fmt.Println("Mismatch between calculated result and JSON data")
// 		return
// 	}

// 	generateExcelAsync(result, excelDone)
// }

// func readJSONFile(filename string, result Result, excelDone chan<- bool) bool {
// 	// Open the JSON file for reading.
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println(err)
// 		return false
// 	}
// 	defer file.Close()

// 	// Create a variable to hold the decoded JSON data.
// 	var resultData Result

// 	// Create a JSON decoder and decode the file into the resultData variable.
// 	decoder := json.NewDecoder(file)
// 	if err := decoder.Decode(&resultData); err != nil {
// 		fmt.Println(err)
// 		return false
// 	}

// 	if result != resultData {
// 		fmt.Println("Mismatch between calculated result and JSON data")
// 		return false
// 	}

// 	if generateExcelAsync(result) {
// 		excelDone <- true
// 		return true
// 	}

// 	return false
// }

// func readJSONFile(filename string, result Result) {
// 	// Open the JSON file for reading.
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()

// 	// Create a variable to hold the decoded JSON data.
// 	var resultData Result

// 	// Create a JSON decoder and decode the file into the resultData variable.
// 	decoder := json.NewDecoder(file)
// 	if err := decoder.Decode(&resultData); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	if result != resultData {
// 		fmt.Println("Mismatch between calculated result and JSON data")
// 		return
// 	}
// }

func readJSONFile(filename string, resultChan chan<- Result) {
	// Open the JSON file for reading.
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a variable to hold the decoded JSON data.
	var resultData Result

	// Create a JSON decoder and decode the file into the resultData variable.
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&resultData); err != nil {
		fmt.Println(err)
		return
	}

	resultChan <- resultData
}

// func generateExcelAsync(resultChan <-chan Result, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	result := <-resultChan // Wait for the result from the channel.

// 	f := excelize.NewFile()

// 	// Create a new sheet.
// 	index := f.NewSheet("Sheet1")

// 	// Set the active sheet to the newly created sheet.
// 	f.SetActiveSheet(index)

// 	// Set the header row values.
// 	f.SetCellValue("Sheet1", "A1", "Number")
// 	f.SetCellValue("Sheet1", "B1", "Factorial")

// 	// Set the data row values.
// 	f.SetCellValue("Sheet1", "A2", result.Number)
// 	f.SetCellValue("Sheet1", "B2", result.Factorial)

// 	// Save the Excel file.
// 	if err := f.SaveAs("resultAsync.xlsx"); err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// }

// func generateExcelAsync(resultChan <-chan Result, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	result := <-resultChan // Wait for the result from the channel.

// 	f := excelize.NewFile()

// 	// Create a new sheet.
// 	if err := f.NewSheet("Sheet1"); err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// Set the active sheet to the newly created sheet.
// 	f.SetActiveSheet(1) // Assuming you want to set the first sheet as active.

// 	// Set the header row values.
// 	f.SetCellValue("Sheet1", "A1", "Number")
// 	f.SetCellValue("Sheet1", "B1", "Factorial")

// 	// Set the data row values.
// 	f.SetCellValue("Sheet1", "A2", result.Number)
// 	f.SetCellValue("Sheet1", "B2", result.Factorial)

// 	// Save the Excel file.
// 	if err := f.SaveAs("resultAsync.xlsx"); err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// }

// func generateExcelAsync(result Result, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	f := excelize.NewFile()

// 	// Create a new sheet.
// 	index, err := f.NewSheet("Sheet1")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// Set the active sheet to the newly created sheet.
// 	f.SetActiveSheet(index)

// 	// Set the header row values.
// 	f.SetCellValue("Sheet1", "A1", "Number")
// 	f.SetCellValue("Sheet1", "B1", "Factorial")

// 	// Set the data row values.
// 	f.SetCellValue("Sheet1", "A2", result.Number)
// 	f.SetCellValue("Sheet1", "B2", result.Factorial)

// 	// Save the Excel file.
// 	if err := f.SaveAs("resultAsync.xlsx"); err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// }

// func generateExcelAsync(result Result, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	f := excelize.NewFile()

// 	// Create a new sheet.
// 	index, err := f.NewSheet("Sheet1")
// 	if err != nil {
// 		fmt.Println("Error creating new sheet:", err)
// 		return
// 	}

// 	// Set the active sheet to the newly created sheet.
// 	f.SetActiveSheet(index)

// 	// Set the header row values.
// 	f.SetCellValue("Sheet1", "A1", "Number")
// 	f.SetCellValue("Sheet1", "B1", "Factorial")

// 	// Set the data row values.
// 	f.SetCellValue("Sheet1", "A2", result.Number)
// 	f.SetCellValue("Sheet1", "B2", result.Factorial)

// 	// Save the Excel file.
// 	if err := f.SaveAs("resultAsync.xlsx"); err != nil {
// 		fmt.Println("Error saving Excel file:", err)
// 		return
// 	}
// }

// func generateExcelAsync(fileResultChan chan Result, wg *sync.WaitGroup) error {
// 	defer wg.Done()
// 	f := excelize.NewFile()
// 	defer func() {
// 		if err := f.Close(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	// Create a new sheet.
// 	index, err := f.NewSheet("Sheet1")
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	// Set value of a cell.
// 	f.SetCellValue("Sheet1", "A2", "Hello world.")
// 	f.SetCellValue("Sheet1", "B2", 100)
// 	// Set active sheet of the workbook.
// 	f.SetActiveSheet(index)
// 	// Save spreadsheet by the given path.
// 	if err := f.SaveAs("Book1.xlsx"); err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Done excel file creation")
// 	return nil
// }

// func generateExcelAsync(fileResultChan chan Result, wg *sync.WaitGroup) error {
// 	defer wg.Done()
// 	f := excelize.NewFile()
// 	defer func() {
// 		if err := f.Close(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	// Create a new sheet.
// 	index, err := f.NewSheet("Sheet1")
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	// Receive data from the channel.
// 	resultData := <-fileResultChan
// 	//fmt.Print(resultData)

// 	// Set the header row values.
// 	f.SetCellValue("Sheet1", "A1", "Number")
// 	f.SetCellValue("Sheet1", "B1", "Factorial")
// 	// Set value of a cell.
// 	f.SetCellValue("Sheet1", "A2", resultData.Number)
// 	f.SetCellValue("Sheet1", "B2", resultData.Factorial)
// 	// Set active sheet of the workbook.
// 	f.SetActiveSheet(index)
// 	// Save spreadsheet by the given path.
// 	if err := f.SaveAs("resultAsync.xlsx"); err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Excel file generated successfully.")
// 	return nil
// }

// func generateExcelAsync(fileResultChan chan Result, wg *sync.WaitGroup) error {
// 	defer wg.Done()
// 	f := excelize.NewFile()
// 	defer func() {
// 		if err := f.Close(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	// Create a new sheet.
// 	index, err := f.NewSheet("Sheet1")
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	// Set the header row values.
// 	f.SetCellValue("Sheet1", "A1", "Number")
// 	f.SetCellValue("Sheet1", "B1", "Factorial")

// 	// Use a select statement to receive data from the channel.
// 	select {
// 	case resultData := <-fileResultChan:
// 		// Data received, set values in cells.
// 		f.SetCellValue("Sheet1", "A2", resultData.Number)
// 		f.SetCellValue("Sheet1", "B2", resultData.Factorial)
// 	case <-time.After(5 * time.Second):
// 		// Timeout (5 seconds) if data is not available.
// 		fmt.Println("Timed out waiting for data.")
// 	}

// 	// Set active sheet of the workbook.
// 	f.SetActiveSheet(index)
// 	// Save spreadsheet by the given path.
// 	if err := f.SaveAs("resultAsync.xlsx"); err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Excel file generated successfully.")
// 	return nil
// }

// func generateExcelAsync(wg *sync.WaitGroup) error {
// 	defer wg.Done()
// 	f := excelize.NewFile()
// 	defer func() {
// 		if err := f.Close(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	// Create a new sheet.
// 	index, err := f.NewSheet("Sheet1")
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	// Set the header row values.
// 	f.SetCellValue("Sheet1", "A1", "Number")
// 	f.SetCellValue("Sheet1", "B1", "Factorial")

// 	// Use a select statement to receive data from the channel with a timeout.

// 	// resultData := <-fileResultChan
// 	// Data received, set values in cells.
// 	f.SetCellValue("Sheet1", "A2", 5)
// 	f.SetCellValue("Sheet1", "B2", 120)
// 	//fmt.Println(resultData)

// 	// Set active sheet of the workbook.
// 	f.SetActiveSheet(index)

// 	// Save the spreadsheet only if data was successfully received.
// 	if err := f.SaveAs("resultAsync.xlsx"); err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	fmt.Println("Excel file generated successfully.")
// 	return nil
// }

// func generateExcelAsync(fileResultChan chan Result, wg *sync.WaitGroup) error {
// 	defer wg.Done()
// 	f := excelize.NewFile()
// 	defer func() {
// 		if err := f.Close(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	// Create a new sheet.
// 	index, err := f.NewSheet("Sheet1")
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	// Set the header row values.
// 	f.SetCellValue("Sheet1", "A1", "Number")
// 	f.SetCellValue("Sheet1", "B1", "Factorial")

// 	// Use a select statement to receive data from the channel with a timeout.
// 	select {
// 	case resultData := <-fileResultChan:
// 		// Data received, set values in cells.
// 		f.SetCellValue("Sheet1", "A2", resultData.Number)
// 		f.SetCellValue("Sheet1", "B2", resultData.Factorial)
// 	case <-time.After(10 * time.Second):
// 		// Timeout (5 seconds) if data is not available.
// 		fmt.Println("Timed out waiting for data.")
// 	}

// 	// Set active sheet of the workbook.
// 	f.SetActiveSheet(index)

// 	// Save the spreadsheet only if data was successfully received.
// 	if err := f.SaveAs("resultAsync.xlsx"); err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	fmt.Println("Excel file generated successfully.")
// 	return nil
// }

// func generateExcelAsync(fileResultChan chan Result, excelDone chan bool) error {
// 	///defer wg.Done()
// 	// Create a new Excel file.
// 	file := xlsx.NewFile()
// 	defer func() {
// 		if err := file.Save("resultAsync.xlsx"); err != nil {
// 			fmt.Println("Error saving Excel file:", err)
// 		}
// 	}()

// 	// Add a new sheet.
// 	sheet, err := file.AddSheet("Sheet1")
// 	if err != nil {
// 		fmt.Println("Error creating sheet:", err)
// 		return err
// 	}

// 	// Set the header row values.
// 	headerRow := sheet.AddRow()
// 	headerRow.AddCell().SetValue("Number")
// 	headerRow.AddCell().SetValue("Factorial")

// 	// Receive data from the channel.
// 	resultData := <-fileResultChan

// 	// Data received, add it to the sheet.
// 	dataRow := sheet.AddRow()
// 	dataRow.AddCell().SetInt(resultData.Number)
// 	dataRow.AddCell().SetString(resultData.Factorial)

// 	fmt.Println("Excel file generated successfully.")
// 	excelDone <- true
// 	return nil
// }

// func generateExcelAsync(fileResultChan chan Result, excelDone chan bool) error {
// 	f := excelize.NewFile()
// 	// Create a new sheet.
// 	index := f.NewSheet("Sheet1")

// 	// Set the header row values.
// 	f.SetCellValue("Sheet1", "A1", "Number")
// 	f.SetCellValue("Sheet1", "B1", "Factorial")

// 	// Receive data from the channel.
// 	//resultData := <-fileResultChan

// 	for result := range fileResultChan {
// 		f.SetCellValue("Sheet1", "A2", result.Number)
// 		f.SetCellValue("Sheet1", "B2", result.Factorial)
// 	}

// 	// Data received, set values in cells.

// 	// Set active sheet of the workbook.
// 	f.SetActiveSheet(index)

// 	// Save the spreadsheet.
// 	if err := f.SaveAs("resultAsync.xlsx"); err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	fmt.Println("Excel file generated successfully.")

// 	// Signal Excel generation completion.
// 	excelDone <- true

// 	return nil
// }

// func generateExcelAsync(fileResultChan chan Result, excelDone chan bool) error {
// 	f := excelize.NewFile()
// 	// Create a new sheet.
// 	index := f.NewSheet("Sheet1")

// 	// Set the header row values.
// 	f.SetCellValue("Sheet1", "A1", "Number")
// 	f.SetCellValue("Sheet1", "B1", "Factorial")

// 	for result := range fileResultChan {
// 		f.SetCellValue("Sheet1", "A2", result.Number)
// 		f.SetCellValue("Sheet1", "B2", result.Factorial)
// 	}

// 	// Set active sheet of the workbook.
// 	f.SetActiveSheet(index)

// 	// Save the spreadsheet.
// 	if err := f.SaveAs("resultAsync.xlsx"); err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	fmt.Println("Excel file generated successfully.")

// 	// Signal Excel generation completion.
// 	excelDone <- true

// 	// Close the excelDone channel.
// 	close(excelDone)

// 	return nil
// }

// func generateExcelAsync(fileResultChan chan Result, excelDone chan bool) error {
// 	f := excelize.NewFile()
// 	// Create a new sheet.
// 	index := f.NewSheet("Sheet1")

// 	// Set the header row values.
// 	f.SetCellValue("Sheet1", "A1", "Number")
// 	f.SetCellValue("Sheet1", "B1", "Factorial")

// 	rowIndex := 2 // Start from the second row
// 	result := <-fileResultChan

// 	cellA := fmt.Sprintf("A%d", rowIndex)
// 	cellB := fmt.Sprintf("B%d", rowIndex)
// 	f.SetCellValue("Sheet1", cellA, result.Number)
// 	f.SetCellValue("Sheet1", cellB, result.Factorial)

// 	// Set active sheet of the workbook.
// 	f.SetActiveSheet(index)

// 	// Save the spreadsheet.
// 	if err := f.SaveAs("resultAsync.xlsx"); err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	//fmt.Println("Excel file generated successfully.")

// 	// Signal Excel generation completion.
// 	excelDone <- true

// 	// Close the excelDone channel.
// 	close(excelDone)

// 	return nil
// }

// func generateExcelAsync(fileResultChan chan Result, excelDone chan bool) error {

// 	f := excelize.NewFile()
// 	// Create a new sheet.
// 	index := f.NewSheet("Sheet1")

// 	// Set the header row values.
// 	f.SetCellValue("Sheet1", "A1", "Number")
// 	f.SetCellValue("Sheet1", "B1", "Factorial")

// 	rowIndex := 2 // Start from the second row
// 	result := <-fileResultChan
// 	fmt.Println(result)
// 	cellA := fmt.Sprintf("A%d", rowIndex)
// 	cellB := fmt.Sprintf("B%d", rowIndex)
// 	f.SetCellValue("Sheet1", cellA, result.Number)
// 	f.SetCellValue("Sheet1", cellB, result.Factorial)

// 	// Set active sheet of the workbook.
// 	f.SetActiveSheet(index)

// 	// Save the spreadsheet.
// 	if err := f.SaveAs("resultAsync.xlsx"); err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	//fmt.Println("Excel file generated successfully.")

// 	// Signal Excel generation completion.
// 	excelDone <- true

// 	// Close the excelDone channel.
// 	close(excelDone)

// 	return nil
// }

// func generateExcelAsync(result Result, excelDone chan bool) error {
// 	f := excelize.NewFile()
// 	// Create a new sheet.
// 	index := f.NewSheet("Sheet1")

// 	// Set the header row values.
// 	f.SetCellValue("Sheet1", "A1", "Number")
// 	f.SetCellValue("Sheet1", "B1", "Factorial")

// 	rowIndex := 2 // Start from the second row

// 	//fmt.Println(result)
// 	cellA := fmt.Sprintf("A%d", rowIndex)
// 	cellB := fmt.Sprintf("B%d", rowIndex)
// 	f.SetCellValue("Sheet1", cellA, result.Number)
// 	f.SetCellValue("Sheet1", cellB, result.Factorial)

// 	// Set active sheet of the workbook.
// 	f.SetActiveSheet(index)

// 	// Save the spreadsheet.
// 	if err := f.SaveAs("resultAsync.xlsx"); err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	//fmt.Println("Excel file generated successfully.")

// 	// Signal Excel generation completion.
// 	excelDone <- true

// 	return nil
// }

// func generateExcelAsync(result Result, excelDone chan<- bool) {
// 	f := excelize.NewFile()
// 	// Create a new sheet.
// 	index := f.NewSheet("Sheet1")

// 	// Set the header row values.
// 	f.SetCellValue("Sheet1", "A1", "Number")
// 	f.SetCellValue("Sheet1", "B1", "Factorial")

// 	rowIndex := 2 // Start from the second row

// 	//fmt.Println(result)
// 	cellA := fmt.Sprintf("A%d", rowIndex)
// 	cellB := fmt.Sprintf("B%d", rowIndex)
// 	f.SetCellValue("Sheet1", cellA, result.Number)
// 	f.SetCellValue("Sheet1", cellB, result.Factorial)

// 	// Set active sheet of the workbook.
// 	f.SetActiveSheet(index)

// 	// Save the spreadsheet.
// 	if err := f.SaveAs("resultAsync.xlsx"); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	//fmt.Println("Excel file generated successfully.")

// 	// Signal Excel generation completion.
// 	excelDone <- true
// }

// func generateExcelAsync(result Result) bool {
// 	f := excelize.NewFile()
// 	// Create a new sheet.
// 	index := f.NewSheet("Sheet1")

// 	// Set the header row values.
// 	f.SetCellValue("Sheet1", "A1", "Number")
// 	f.SetCellValue("Sheet1", "B1", "Factorial")

// 	rowIndex := 2 // Start from the second row

// 	cellA := fmt.Sprintf("A%d", rowIndex)
// 	cellB := fmt.Sprintf("B%d", rowIndex)
// 	f.SetCellValue("Sheet1", cellA, result.Number)
// 	f.SetCellValue("Sheet1", cellB, result.Factorial)

// 	// Set active sheet of the workbook.
// 	f.SetActiveSheet(index)

// 	// Save the spreadsheet.
// 	if err := f.SaveAs("resultAsync.xlsx"); err != nil {
// 		fmt.Println(err)
// 		return false
// 	}

// 	return true
// }

func generateExcelAsync(result Result) bool {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")

	// Set the header row values.
	f.SetCellValue("Sheet1", "A1", "Number")
	f.SetCellValue("Sheet1", "B1", "Factorial")

	rowIndex := 2 // Start from the second row

	cellA := fmt.Sprintf("A%d", rowIndex)
	cellB := fmt.Sprintf("B%d", rowIndex)
	f.SetCellValue("Sheet1", cellA, result.Number)
	f.SetCellValue("Sheet1", cellB, result.Factorial)

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	// Save the spreadsheet.
	if err := f.SaveAs("resultAsync.xlsx"); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
