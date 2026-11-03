package conditional

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func CommandLineMenu() {
	reader := bufio.NewReader(os.Stdin)

	const pinCode = 1234

	for {
		// scans the 4 digit code
		fmt.Print("Enter 4-digit pin code: ")
		inputCode, readErr := reader.ReadString('\n')

		//trims the white spaces
		inputCode = strings.TrimSpace(inputCode)

		//checks if the input is empty
		if inputCode == "" {
			fmt.Println("Inputted code is empty!")
		}

		//checks if the inputted is a 4 digit number
		if len(inputCode) != 4 {
			fmt.Println("PIN must be exactly 4 characters")
		}

		// checks if there is a read error problem
		if readErr == nil || readErr == io.EOF {

			//converts the string to integer
			convertedCode, err := strconv.Atoi(inputCode)

			//checks if there is error when converting
			if err == nil {
				if convertedCode == pinCode {
					var balance float64
					balance = 2000
				menu:
					for {

						fmt.Print("Choose transaction: ")
						fmt.Println("")
						fmt.Println("[1] Balance Inquiry")
						fmt.Println("[2] Deposit")
						fmt.Println("[3] Withdraw")
						fmt.Println("[4] Exit")
						fmt.Print("Enter your choice: ")

						input, err1 := reader.ReadString('\n')
						if err1 != nil {
							fmt.Println("Error reading transaction:", err1)
							continue
						}

						input = strings.TrimSpace(input)

						transaction, err1 := strconv.Atoi(input)

						if err1 != nil {
							fmt.Println("Invalid transaction!")
							continue
						}

						if transaction != 1 && transaction != 2 && transaction != 3 && transaction != 4 {
							fmt.Println("Invalid transaction! Please choose from 1-4.")
							continue
						}

						switch transaction {
						case 1:
							fmt.Println("Balance: ", balance)
						case 2:
							fmt.Print("Enter value you want to deposit: ")
							inputDeposit, err3 := reader.ReadString('\n')
							if err3 != nil {
								fmt.Println("Error reading transaction:", err3)
								continue
							}

							inputDeposit = strings.TrimSpace(inputDeposit)

							deposit, err3 := strconv.ParseFloat(inputDeposit, 64)

							if err3 != nil {
								fmt.Println("Invalid input deposit!")
								continue
							}

							if deposit <= 0 {
								fmt.Println("Please enter greater than zero!")
								continue
							}

							balance += deposit
							fmt.Println("Your new balance: ", balance)
						case 3:
							fmt.Print("Enter value to withdraw: ")
							inputWithdraw, err4 := reader.ReadString('\n')
							if err4 != nil {
								fmt.Println("Error reading transaction:", err4)
								continue
							}

							inputWithdraw = strings.TrimSpace(inputWithdraw)

							withdraw, err3 := strconv.ParseFloat(inputWithdraw, 64)

							if err3 != nil {
								fmt.Println("Invalid input deposit!")
								continue
							}

							if withdraw > balance || balance == 0 {
								fmt.Println("Insufficient balance! Try another value.")
								continue
							}

							if withdraw < 0 {
								fmt.Println("Please enter greater than zero!")
								continue
							}

							balance -= withdraw
							fmt.Println("You have withdrawn ", withdraw)
							fmt.Println("Your new balance is ", balance)
						case 4:
							fmt.Println("Goodbye! Thank you for using our service.")
							break menu
						}

					}

				} else {
					fmt.Println("Invalid pincode! Try again.")
				}

			} else {
				fmt.Println("invalid input")
			}
		} else {
			fmt.Println("Error input: ", readErr)
		}

	}

}
