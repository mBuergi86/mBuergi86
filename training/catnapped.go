package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomElement(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

func main() {
	rand.Seed(time.Now().UnixNano())

	guest, object, name :=
		[3]string{"Famous people", "People you know", "Goofy names"},
		[3]string{"A Toy Chest", "A Crate", "A Box"},
		[3]string{"A Tale of Two Cities", "A Tale of"}

	fmt.Println("Guests:", guest)
	fmt.Println("Objects:", object)
	fmt.Println("Names:", name)

	culprit := getRandomElement(guest[:])
	secretObject := getRandomElement(object[:])

	fmt.Println(culprit, "hid the car by putting it in the", secretObject)
}
