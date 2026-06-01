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

func main() {
	getName()
}
