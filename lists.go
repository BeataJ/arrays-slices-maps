package main

import "fmt"

func main() {
	var productNames = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpert"

	fmt.Println(productNames)

	fmt.Println(prices[2])

	featuredPrices := prices[1:]
	featuredPrices[0] = 199.99
	highlitedPrices := featuredPrices[:1]
	fmt.Println(featuredPrices)
	fmt.Println(highlitedPrices)
	fmt.Println(prices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))
}
