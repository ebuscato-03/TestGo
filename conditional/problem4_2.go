package conditional

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func LeapYearChecker() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a year: ")
		year, readErr := reader.ReadString('\n')

		year = strings.TrimSpace(year)

		if readErr != nil && readErr != io.EOF {
			fmt.Println("Error reading: ", readErr)
			return
		}

		if year == "" {
			fmt.Println("Please enter a year!")
		} else {
			yearConverted, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("Please enter a valid year!")
			}else{
				if yearConverted < 1 || yearConverted > 9999 {
					fmt.Println("Please enter year within the range!")
				} else {
					//LEAP YEAR RULES
					/*	- divisible by 4 - leap year
						- divisible by 100 - not a leap year
						- divisible by 400 - leap year even though its divisible by 100
					*/

					if yearConverted % 400 == 0 || (yearConverted % 4 == 0 && yearConverted % 100 != 0){
						fmt.Println(yearConverted, " is a leap year!")
						return
					} else {
						fmt.Println(yearConverted, " is not a leap year!")
						return
					}
				}
			}

		}

	}
}
