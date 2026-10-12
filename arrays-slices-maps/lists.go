package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames = [4]string{"A book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)

	fmt.Println(prices[2])

	featuredPrices := prices[1:3]
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices) // slices
}
