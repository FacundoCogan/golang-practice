package main

import (
	"fmt"
)

func main() {

	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// revenue, expenses, taxRate = returnValues()
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n ", ebt)
	fmt.Printf("Profit: %.1f\n ", profit)
	fmt.Printf("Ratio: %.3f\n ", ratio)

}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
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
