package loops

import (
	"fmt"
)

func SentinelControlledSum() {

	var input int
	counter := 0
	sum := 0
	for {

		fmt.Print("Enter a number (Press 0 to exit): ")

		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Invalid input!")
		}

		counter++
		sum += input

		if input == 0 {
			fmt.Println("Sum: ", sum)
			fmt.Println("Count: ", counter-1)
			break
		}
	}
}
