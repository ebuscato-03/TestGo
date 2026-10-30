package userinput

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InteractiveProfile() {

	var age int64

	//NAME
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your full name: ")
	name, err := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if err != nil {
		fmt.Println("Input ended or an error occurred:", err)
		return
	}

	//AGE
	fmt.Print("Enter your age: ")
	_, err = fmt.Scan(&age)
	if err != nil {
		fmt.Println("Invalid age entered")
		return
	}

	if age < 1 || age > 120 {
		fmt.Println("Age must be 1-120")
		return
	}
	reader.ReadString('\n')

	//CITY
	fmt.Print("Enter your city: ")
	city, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Input ended or an error occurred:", err)
		return
	}
	city = strings.TrimSpace(city)

	fmt.Print("Enter your occupation: ")
	occupation, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Input ended or an error occurred:", err)
		return
	}
	occupation = strings.TrimSpace(occupation)

	//DISPLAY PROFILE

	fmt.Println("     ===== My Profile =====")
	fmt.Println()
	fmt.Printf("%12s %s\n", "Name: ", name)
	fmt.Printf("%12s %d\n", "Age: ", age)
	fmt.Printf("%12s %s\n", "City: ", city)
	fmt.Printf("%12s %s\n", "Occupation: ", occupation)
}
