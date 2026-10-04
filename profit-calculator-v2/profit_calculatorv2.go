package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// revenue, expenses, taxRate = returnValues()
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n ", ebt)
	fmt.Printf("Profit: %.1f\n ", profit)
	fmt.Printf("Ratio: %.3f\n ", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Invalid number. Must be bigger than 0")
	}

	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {

	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio

}

// func returnValues() (revenue float64, expenses float64, taxRate float64) {

// 	fmt.Print("Revenue: ")
// 	fmt.Scan(&revenue)

// 	fmt.Print("Expenses: ")
// 	fmt.Scan(&expenses)

// 	fmt.Print("Tax rate: ")
// 	fmt.Scan(&taxRate)

// 	return revenue, expenses, taxRate

// }
