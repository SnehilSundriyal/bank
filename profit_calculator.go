package main

//
//import (
//	"errors"
//	"fmt"
//	"os"
//)
//
//const resultsFile = "results.txt"
//
//func main() {
//	revenue, err := getUserInput("Enter Revenue: ")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	expenses, err := getUserInput("Enter Expenses ")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	taxRate, err := getUserInput("Enter Tax Rate: ")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
//
//	fmt.Printf("EBT is %.2f\nProfit is %.2f\nRatio is %.3f\n", ebt, profit, ratio)
//
//	storeResults(ebt, profit, ratio)
//}
//
//func getUserInput(message string) (float64, error) {
//	var someVar float64
//	fmt.Print(message)
//	fmt.Scan(&someVar)
//	for someVar <= 0 {
//		return 0, errors.New("Value must be positive")
//	}
//	return someVar, nil
//}
//
//func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
//	ebt := revenue - expenses
//	profit := ebt * (1 - taxRate/100)
//	ratio := ebt / profit
//	return ebt, profit, ratio
//}
//
//func storeResults(ebt, profit, ratio float64) {
//	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.3f\n", ebt, profit, ratio)
//	os.WriteFile(resultsFile, []byte(results), 0644)
//}
