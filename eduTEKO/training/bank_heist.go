package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	isHeistOn := true

	if eludedGuards := rand.Intn(100); eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		fmt.Println("Plan a better disguise next time?")
		isHeistOn = !isHeistOn
	}

	if openedVault := rand.Intn(100); openedVault >= 70 && isHeistOn {
		fmt.Println("Grab an GO!")
	} else if isHeistOn == true { //but openedVault is less than 70 (so we couldn’t open the vault)
		isHeistOn = false
		fmt.Println("The Vault can't be opened.")
	} else {
		fmt.Println("What's the combo to this lock again??")
	}

	if leftSafely := rand.Intn(5); isHeistOn == true {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Looks like you tripped an alarm... run?")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)

		fmt.Println("$", amtStolen, "not bad!")
	}
}
