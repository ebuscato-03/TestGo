package operations

import (
	"fmt"
	"math"
)

func ArithmeticCalculator() {

	var input1 float64
	var input2 float64

	var choice string

	fmt.Print("Enter first number: ")
	_, err1 := fmt.Scan(&input1)
	if err1 != nil {
		fmt.Println("Invalid input: ", err1)
		return
	}

	fmt.Print("Enter second number: ")
	_, err2 := fmt.Scan(&input2)
	if err2 != nil {
		fmt.Println("Invalid input: ", err2)
		return
	}

	fmt.Print("Enter choice: ")
	_, err3 := fmt.Scan(&choice)
	if err3 != nil {
		fmt.Println("Invalid input: ", err3)
		return
	}
	switch choice {
	case "+":
		sum := input1 + input2
		fmt.Println(sum)
	case "-":
		diff := input1 - input2
		fmt.Println(diff)
	case "*":
		prod := input1 * input2
		fmt.Println(prod)

	// Handles division by zero
	case "/":
		if input2 == 0 {
			fmt.Println("division by zero")
		} else {
			quo := input1 / input2
			fmt.Println(quo)
		}
	case "%":
		mod := math.Mod(input1, input2)
		// mod := input1 % input2
		fmt.Println(mod)

	default:
		fmt.Println("Invalid operator")
	}

}
