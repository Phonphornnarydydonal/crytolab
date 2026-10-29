package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter number a: ")
	fmt.Scan(&a)

	fmt.Print("Enter number b: ")
	fmt.Scan(&b)

	bothPositive := a > 0 && b > 0

	oneGreater := a > b || b > a

	notEqual := !(a == b)

	fmt.Println("\n--- Logical Operators Results ---")
	fmt.Println("Both numbers are positive:", bothPositive)
	fmt.Println("One number is greater than the other:", oneGreater)
	fmt.Println("Numbers are not equal:", notEqual)
}
