package conditional

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func GradeClassification() {

	reader := bufio.NewReader(os.Stdin)

	for {
		//Asks the user to input their grade
		fmt.Print("Enter your grade: ")
		input, readErr := reader.ReadString('\n')

		//trims whitespaces to the input
		input = strings.TrimSpace(input)

		// checks if there is error in their input and end of line input
		if readErr != nil && readErr != io.EOF {
			fmt.Println("Error reading input: ", readErr)
			return
		}

		//shows the user if their input is empty
		if input == "" {

			fmt.Println("Input cannot be empty. ")

		} else {
			//converts the string to a float64
			inputConverted, conversionErr := strconv.ParseFloat(input, 64)

			//checks if there is error in the conversion
			if conversionErr != nil {
				fmt.Println("Please enter a valid grade!")

			} else {

				if inputConverted < 0 || inputConverted > 100 {
					fmt.Println("Please enter a valid grade!")
				} else {
					//initializing the gradeclass var as string
					var gradeClass string

					//creating if else if for the boundary
					if inputConverted >= 90 {
						gradeClass = "A"
					} else if inputConverted >= 80 {
						gradeClass = "B"
					} else if inputConverted >= 70 {
						gradeClass = "C"
					} else if inputConverted >= 60 {
						gradeClass = "D"
					} else if inputConverted >= 0 {
						gradeClass = "F"
					} else {
						gradeClass = "Invalid range"
					}

					//Prints the result
					fmt.Println("Grade: ", inputConverted)
					fmt.Println("Grade Classification: ", gradeClass)
					return
				}
			}
		}

	}

}
