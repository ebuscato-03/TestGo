package loops

import (
	"fmt"
)

func NumberSequenceGenerator() {

	var startNum int64
	var endNum int64
	var step int64

	fmt.Print("Enter starting number: ")
	_, err := fmt.Scan(&startNum)

	if err != nil {
		fmt.Println("Invalid input!")
	}

	fmt.Print("Enter ending number: ")
	_, err1 := fmt.Scan(&endNum)

	if err1 != nil {
		fmt.Println("Invalid input!")
	}

	fmt.Print("Enter a step: ")
	_, err2 := fmt.Scan(&step)
	if err2 != nil {
		fmt.Println("Invalid step!")
	}

	//handles 0 step
	if step <= 0 {
		fmt.Println("Step cannot be zero!")
	} else {
		switch {
		case startNum < endNum:
			for i := startNum; i <= endNum; i += step {
				if i == endNum {
					fmt.Print(i)
				} else {
					fmt.Print(i, ", ")

				}
			}
		case startNum > endNum:
			for i := startNum; i >= endNum; i -= step {
				if i == endNum {
					fmt.Print(i)
				} else {
					fmt.Print(i, ", ")
				}
			}
		default:
			for i := startNum; i <= endNum; i += step {
				if i == endNum {
					fmt.Print(i)
				} else {
					fmt.Print(i, ", ")

				}
			}
		}
	}
}
