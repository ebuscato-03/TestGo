package loops


import (
    "fmt"
    "strconv"
)

func NumberPropertyAnalyzer(){

    fmt.Print("Enter number to check: ")
    var number int
    _, err := fmt.Scan(&number)
    if err != nil {
        fmt.Println("Invalid input!")
        return
    }

    if number < 0 {
        fmt.Println("Invalid input!")
        return
    }

    count:=0
    if number == 0 {
        count=1
    }
    accumulator :=0 

    
    //reverse & palindrome
    //converted the number to string
    stringNum := strconv.Itoa(number)
    first := true

    //declared a reversed string variable
    reversedNum := ""
    //used the length of the string -1 to prevent arrayoutofbounds
    for i := len(stringNum) - 1; i >= 0; i-- {
        //converted each string to int value by converting the string value of byte stringNum[i]
        digit, err := strconv.Atoi(string(stringNum[i]))

        if err != nil {}
        if !first {
            fmt.Print(", ")
        }
        //printed the digit
        fmt.Print(digit)
        first = false
        
        //concatenated each of the string digit to the reversedNum variable for comparison
        reversedNum += string(stringNum[i])
    }
    fmt.Println("")

    //compared the original number which is converted to string to the reversed string from the loop
    if (stringNum == reversedNum) {
        fmt.Println(number, " is a palindrome!")
    } else {
        fmt.Println(number, " is not a palindrome!")
    }

    fmt.Println("")
    //digit counter & sum
    for number > 0{
        digit := number % 10; //extracting the digit
        accumulator += digit // added each of the digit to accumulator for the sum
        number /= 10 // removed the last digit
        count++ // counted the digits
    }
    
    fmt.Println("Number of digits: ", count)
    fmt.Println("Sum: ", accumulator)
}