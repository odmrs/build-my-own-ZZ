package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Guess Game
	//

	// Bufio -> package read things from users, and create a scanner to be analyzing the stdin, input from user in console
	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 100
	tries := 1

	fmt.Printf("Please think of a number between %d and %d \n", low, high)
	fmt.Println("Press ENTER when ready")

	scanner.Scan()

	for {
		guess := (low + high) / 2
		fmt.Printf("------- TRY %d -------\n", tries)
		fmt.Println("I guess the number is ", guess)
		fmt.Println("Is that:")
		fmt.Println("(a) too hight?")
		fmt.Println("(b) too low?")
		fmt.Println("(c) too correct?")

		scanner.Scan()
		response := scanner.Text() // Get all the things from user input as string

		if tries == 7 {
			fmt.Println("You lie, I quit!")
			break
		}

		switch response {
		case "a":
			high = guess - 1
		case "b":
			low = guess - 1
		case "c":
			fmt.Println("I won!")
			break
		default:
			fmt.Println("I don't know what to do here... Try again")
		}

		tries++
	}
}

// 6
