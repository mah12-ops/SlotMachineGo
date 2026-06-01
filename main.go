package main

import "fmt"

func getName() string {
	name := ""

	fmt.Println("welcoeme to mercy's game")
	fmt.Printf("Enter your name : ")
	_, err := fmt.Scanln(&name)

	if err != nil {
		return ""
	}

	println("welcome to the game", name)
	return name
}

func getBet(balance uint) uint {
	var bet uint

	for true {
		fmt.Printf("Enter your bet , or 0 to quit (balance = $%d ) : ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println(" Bet can not be greater that balance .")
		} else {
			break
		}
	}
	return bet
}

func main() {
	balance := uint(200)

	getName()
	bet := getBet(balance)
	fmt.Println(bet)
}
