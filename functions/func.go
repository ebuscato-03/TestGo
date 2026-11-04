package functions

import "fmt"

func Test() {
	add(5, 6)
	fmt.Println(subtract(5, 6))
}

// func add(a int, b int) {
// 	fmt.Println(a + b)
// }

func add(a, b int) {
	fmt.Println(a + b)
}

func subtract(a, b int) int {
	return a - b
}
