package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumberGen(min float32, max float32) float32 {
	return min + (max-min)*rand.Float32()
}

// Task implementation goes here

type Stock struct {
	name  string
	price float32
}

func (stock *Stock) updateStock() {
	change := randomNumberGen(-10000, 10000)
	stock.price += change
}

func displayMarket(market Stock) {
	fmt.Println(market)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// Function calls go here
	stockMarket := []Stock{{"GOOG", 2313.50}, {"APPL", 157.28}, {"FB", 203.77}, {"TWTR", 50.00}}

	stockMarket[0].updateStock()

	displayMarket(stockMarket[0])
	displayMarket(stockMarket[1])
	displayMarket(stockMarket[2])
	displayMarket(stockMarket[3])
}
