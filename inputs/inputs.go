package inputs

import "fmt"

func Input() {
	var name string
	fmt.Println("Enter name: ")
	fmt.Scan(&name)

	fmt.Println("Hello", name)
}
