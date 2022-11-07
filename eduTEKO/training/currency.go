package main

import (
	"fmt"
)

func main() {
	currencies := map[string]float32{
		"JPY": 130.2,
		"EUR": 0.95,
	}

	var dollarAmount float32
	var currency string

	fmt.Println("Please enter Amount")
	fmt.Scan(&dollarAmount)

	if dollarAmount == 0 {
		fmt.Println("Invalid stock price")
	} else {
		fmt.Println("Please enter Curreny")
		fmt.Scan(&currency)
	}

	rate, isValid := currencies[currency]

	if isValid {
		fmt.Println("We got", rate*dollarAmount, "from the map")
	} else {
		fmt.Println(isValid, "was not found in the map")
	}
}
