package loops

import (
	"fmt"
)

func NumberCounter() {

	first := true
	for i := 0; i <= 100; i++ {
		if !first {
			fmt.Print(", ")
		}

		fmt.Print(i)
		first = false
	}
}
