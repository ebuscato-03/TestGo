package loops

import (
	"fmt"
)

func NumberSequenceGenerator() {

	var firstNum int64
	var secondNum int64
	fmt.Println("Enter first number: ")
	_, err := fmt.Scan(&firstNum)

	if err != nil {
		fmt.Println("Invalid input!")
	}
}
