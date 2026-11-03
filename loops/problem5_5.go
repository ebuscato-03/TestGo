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
