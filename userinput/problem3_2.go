package userinput

import (
	"fmt"
	"strings"
)

func TemperatureConverter() {
	var temp float64
	var unit string
	var unitToConvert string
	var finalTemp float64

	fmt.Print("Enter temperature: ")
	_, err := fmt.Scan(&temp)
	if err != nil {
		fmt.Println("Invalid temperature")
		return
	}

	fmt.Print("Enter unit (C, F, K): ")
	_, err = fmt.Scan(&unit)
	if err != nil {
		fmt.Println("Invalid unit")
		return
	}

	fmt.Print("Enter unit to convert to (C, F, K): ")
	_, err = fmt.Scan(&unitToConvert)
	if err != nil {
		fmt.Println("Invalid target unit")
		return
	}

	// Make units uppercase so lowercase input also works.
	unit = strings.ToUpper(unit)
	unitToConvert = strings.ToUpper(unitToConvert)

	// Validate the units.
	if (unit != "C" && unit != "F" && unit != "K") ||
		(unitToConvert != "C" && unitToConvert != "F" && unitToConvert != "K") {
		fmt.Println("Invalid unit. Use C, F, or K.")
		return
	}

	// Validate absolute zero.
	if (unit == "K" && temp < 0) ||
		(unit == "C" && temp < -273.15) ||
		(unit == "F" && temp < -459.67) {
		fmt.Println("Invalid temperature: below absolute zero")
		return
	}

	// Celsius to Fahrenheit
	if unit == "C" && unitToConvert == "F" {
		finalTemp = temp*9/5 + 32
	}

	// Fahrenheit to Celsius
	if unit == "F" && unitToConvert == "C" {
		finalTemp = (temp - 32) * 5 / 9
	}

	// Celsius to Kelvin
	if unit == "C" && unitToConvert == "K" {
		finalTemp = temp + 273.15
	}

	// Kelvin to Celsius
	if unit == "K" && unitToConvert == "C" {
		finalTemp = temp - 273.15
	}

	// Fahrenheit to Kelvin
	if unit == "F" && unitToConvert == "K" {
		finalTemp = (temp-32)*5/9 + 273.15
	}

	// Kelvin to Fahrenheit
	if unit == "K" && unitToConvert == "F" {
		finalTemp = (temp-273.15)*9/5 + 32
	}

	// Same-unit conversion
	if unit == unitToConvert {
		finalTemp = temp
	}

	fmt.Printf("Converted Temperature: %.2f %s\n",
		finalTemp, unitToConvert)
}
