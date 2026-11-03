package loops

import (
	"fmt"
)

func EvenNumCounter() {

	var input int
	fmt.Print("Enter a number: ")

	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Invalid input")
	}

	accumulator := 0
	first := true
	for i := 0; i <= input; i++ {
		if i%2 == 0 {
			accumulator++

			if !first {
				fmt.Print(",")
			}

			fmt.Print(i)
			first = false
		}
	}
	fmt.Println("-number of even numbers: ", accumulator)

}

func EvenNumCounterVer2() {
	var input int64

	evenCounter := 0
	for {

		fmt.Print("Enter a number (Press 0 to count number of even numbers): ")

		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Invalid input")
		}

		if input%2 == 0 {
			evenCounter++
		}
		if input == 0 {
			fmt.Println("Total even numbers: ", evenCounter)
			break
		}
	}

}
