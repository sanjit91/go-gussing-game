package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// generate a random number between 1-100
	rand.Seed(time.Now().Unix())
	secretNumber := rand.Intn(100)

	// print the secret number
	fmt.Println(secretNumber)

	// store a attempt number
	attempt := 10
	for attempt > 0 {
		// read the input from stdin n
		fmt.Println("Enter your number: ")
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Err reading input")
			return
		}

		// if the input is less then the secret number --> print "Input is too low"
		if input < secretNumber {
			fmt.Printf("Input is too low (attempt left: %d), try again \n", attempt-1)
		} else if input > secretNumber {
			fmt.Printf("Input is too high (attempt left: %d), try again\n", attempt-1)
		} else {
			fmt.Printf("You guessed the number %d in %d attempts \n", secretNumber, attempt-1)
			return
		}

		// reduce the attemt counter
		attempt--
	}

	// Need to handle some edge cases

	// print the input
	// fmt.Println((input))

	// Check if the number is correct or wrong
	// if input == secretNumber {
	// 	fmt.Println("Correct!")
	// } else {
	// 	fmt.Println("Wrong!")
	// }
}
