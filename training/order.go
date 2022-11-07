package main

import (
	"fmt"
)

func askOrder() string {
	var input string
	fmt.Print("What would you like to eat: ")
	// Get the input from the user
	fmt.Scan(&input)
	return input
}

// WRITE CONTAINS FUNCTION BELOW
func contains(menu []string, order string) bool {
	for _, item := range menu {
		if order == item {
			return true
		}
	}
	return false
}

func main() {
	fastfoodMenu := []string{"Burgers", "Nuggets", "Fries"}
	fmt.Println("The fast food menu has these items:", fastfoodMenu)

	var total int
	var order string
	// WRITE INDEFINITE LOOP ASKING FOR ORDERS BELOW
	for order != "quit" {
		order = askOrder()
		if contains(fastfoodMenu, order) {
			total += 4
		} else {
			fmt.Println("This item is not on the menu.")
		}
	}

	// WRITE DEFINITE LOOP COUNTING $2 BILLS BELOW
	for amount := total; amount > 0; amount -= 2 {
		fmt.Println("The amount owed is:", amount)
		fmt.Println("Handing over a two dollar bill.")
	}

	fmt.Println("The total for the order is", total)
}
