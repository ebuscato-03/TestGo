package operations

import (
	"fmt"
	"math"
)

func ReceiptCalculator() {

	const discount = .10
	var item1, item2, item3 int64
	var q1, q2, q3 int64

	fmt.Print("Enter price of 1st item: ")
	_, err1 := fmt.Scan(&item1)
	if err1 != nil {
		fmt.Println("Invalid price")
		return
	}

	fmt.Print("Enter quantity: ")
	_, err2 := fmt.Scan(&q1)
	if err2 != nil {
		fmt.Println("Invalid quantity")
		return
	}

	fmt.Print("Enter price of 2nd item: ")
	_, err3 := fmt.Scan(&item2)
	if err3 != nil {
		fmt.Println("Invalid price")
		return
	}

	fmt.Print("Enter quantity: ")
	_, err4 := fmt.Scan(&q2)
	if err4 != nil {
		fmt.Println("Invalid quantity")
		return
	}

	fmt.Print("Enter price of third item: ")
	_, err5 := fmt.Scan(&item3)
	if err5 != nil {
		fmt.Println("Invalid quantity")
		return
	}

	fmt.Print("Enter quantity: ")
	_, err6 := fmt.Scan(&q3)
	if err6 != nil {
		fmt.Println("Invalid quantity")
		return
	}

	fmt.Print("Do you have discount card? (Y- Yes | N- No):  ")
	var hasDiscount string
	_, err7 := fmt.Scan(&hasDiscount)
	if err7 != nil {
		fmt.Println("Invalid input")
		return
	}

	TotalPrice := (item1 * q1) + (item2 * q2) + (item3 * q3)

	discountAmount := int64(math.Round(float64(TotalPrice) * discount))

	DiscountedPrice := TotalPrice - discountAmount

	switch hasDiscount {
	case "Y":
		fmt.Printf("Total Price: ₱%d.%02d\n", TotalPrice/100, TotalPrice%100)
		fmt.Printf("Discount: ₱%d.%02d\n", discountAmount/100, discountAmount%100)
		fmt.Printf("Discounted Price: ₱%d.%02d\n", DiscountedPrice/100, DiscountedPrice%100)
	case "N":
		fmt.Printf("Total Price: ₱%d.%02d\n", TotalPrice/100, TotalPrice%100)
		fmt.Printf("Discount: ₱%d.%02d\n", 0, 0)
		fmt.Printf("Discounted Price: ₱%d.%02d\n", TotalPrice/100, TotalPrice%100)
	default:
		fmt.Println("Invalid Input")
	}
}
