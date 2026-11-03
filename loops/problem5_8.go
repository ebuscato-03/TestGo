package loops

import (
	"fmt"
)

func GuessingGame() {

	var input int
	const num = 11
	counter := 0
	for {

		fmt.Print("Guess the number: ")

		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Invalid input!")
			return
		}
		counter++
		if input == num {
			fmt.Println("You have guessed the number! ")
			fmt.Println("Number of guesses: ", counter)
			break
		}

		if input > num {
			fmt.Println("Your guess is too high!")
			continue
		}

		if input < num {
			fmt.Println("Your guess is too low!")
			continue
		}
	}
}
