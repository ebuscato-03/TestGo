package loops

import (
	"fmt"
)

func SumOfN() {

	var input int
	fmt.Print("Enter a number: ")

	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Invalid input!")
	}

	accumulator := 0
	first := true
	for i := 0; i <= input; i++ {
		if !first {
			fmt.Print(" + ")
		}
		fmt.Print(i)
		first = false
		accumulator += i
	}
	fmt.Println(" = ", accumulator)

}
