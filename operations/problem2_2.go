package operations

import "fmt"

func check() bool {
	fmt.Println("check() was called")
	return true
}

func Expressions() {

	//PART 1
	exp := 1*8 - 6 + 7/2
	fmt.Println(exp)

	//PREDICTION - 5
	//ACTUAL RESULT - 5

	exp1 := 1*(8-6) + 7/2
	fmt.Println(exp1)

	//PREDICTION - 5
	//ACTUAL RESULT - 5

	exp2 := 1 * (8 - 6 + 7) / 2
	fmt.Println(exp2)

	//PREDICTION - 4
	//ACTUTAL - 4

	//PART 2
	exp3 := 1 == 1 // TRUE because 1 is equal to 1
	fmt.Println(exp3)

	exp4 := 1 != 2 // TRUE because 1 is not equal to 2
	fmt.Println(exp4)

	exp5 := 1 > 2 // FALSE because 1 is less than 2
	fmt.Println(exp5)

	exp6 := 1 < 2 // TRUE because 2 is greater than 1
	fmt.Println(exp6)

	exp7 := 1 <= 3 // TRUE because 1 is less than to 3 - this comparison operator checks is the first num is less than or equal to second num
	fmt.Println(exp7)

	exp8 := 1 >= 3 // FALSE because 1 is less than 3 - this comparison operator checks is the first num is less than or equal to second num
	fmt.Println(exp8)

	//PART 3

	exp10 := true && true // TRUE because AND checks if both condition is true and when it does, it returns true
	fmt.Println(exp10)

	exp11 := true && false // FALSE because AND checks if both condition is true and when it does, it returns true
	fmt.Println(exp11)

	exp12 := true || true // TRUE because OR checks if one of the condition is true and when it does, it returns true
	fmt.Println(exp12)

	exp13 := true || false // TRUE because OR checks if one of the condition is true and when it does, it returns true
	fmt.Println(exp13)

	exp14 := !true // FALSE because ! inverts the the true
	fmt.Println(exp14)

	exp15 := !false // TRUE because ! inverts the false
	fmt.Println(exp15)

	//PART 4
	//SHORT CIRCUIT
	if result := false && check(); result {
		fmt.Println("First expression evaluated")
	}

	if result := false || check(); result {
		fmt.Println("Second expression evaluated")
	}

	//ONLY THE SCOND EXPRESSION WAS EVALUATED BECAUSE THE FIRST EXPRESSION WAS ALREADY FALSE AND SHORT CIRCUIT SKIPS IT COZ THE OPERATOR IS AND, AND IT NEEDS TO BE BOTH TRUE
	//IT EVALUATED THE SECOND EXPRESSION BECAUSE IT USES OR AND ONLY 1 IS NEEDED TO BE TRUE
}
