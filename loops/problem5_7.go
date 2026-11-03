package loops

import (
	"fmt"
)

func MultiplicationTableofN() {

	var input int
	var prod int
	fmt.Print("Enter a number to generate: ")
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Invalid input!")
	}

	for i := 0; i <= 10; i++ {
		prod = input * i
		fmt.Printf("%d x %d = %d\n", input, i, prod)
	}

}
