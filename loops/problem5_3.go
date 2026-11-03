package loops

import (
	"fmt"
)

func SumOfNumbers() {
	var input int64
	accumulator := 0

	for {
		fmt.Print("Enter numbers to add (Enter 0 to stop): ")
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Invalid input!")
		}
		accumulator += int(input)

		if input == 0 {
			fmt.Println("Sum: ", accumulator)
			return
		}
	}

}
