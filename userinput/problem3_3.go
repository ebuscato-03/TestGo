package userinput

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func RobustInputValidator() {
	//RANGE 1- 100

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter an integer: ")
		input, readErr := reader.ReadString('\n')
		if readErr != nil && readErr != io.EOF {
			fmt.Println("Error reading input:", readErr)
			return
		}
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Input cannot be empty")
		} else {
			inputConverted, err := strconv.Atoi(input)

			if err == nil {
				if inputConverted >= 1 && inputConverted <= 100 {
					fmt.Println("Input is valid!")
					break
				} else {
					fmt.Println("Input is out of range!")
				}
			} else {
				fmt.Println("Invalid input: please enter a whole number.")
			}
		}
	}
}

//NOT NEEDED
// func isDecimal(input string) bool {
// 	value, err := strconv.ParseFloat(input, 64)

// 	if err != nil {
// 		return false
// 	}

// 	return value != float64(int64(value))
// }

// func isWithinIntegerRange(input string) bool {
// 	_, err := strconv.ParseInt(input, 10, 64)

// 	return err == nil
// }
