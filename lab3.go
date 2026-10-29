package main

import "fmt"

// Bitwise XOR
func myXor(a, b int) {
	fmt.Println("XOR (a ^ b):", a^b)
}

// Bitwise NOT
func myNOT(a, b int) {
	fmt.Println("NOT (~a):", ^a)
	fmt.Println("NOT (~b):", ^b)
}

// Bitwise OR
func myOR(a, b int) {
	fmt.Println("OR (a | b):", a|b)
}

// Bitwise AND
func myAND(a, b int) {
	fmt.Println("AND (a & b):", a&b)
}

// Left and right shift
func myShift(a, b int) {
	fmt.Println("Left shift (a << 1):", a<<1)
	fmt.Println("Right shift (a >> 1):", a>>1)
}

// Assignment operators
func myAssignment(a int) {
	x := a
	fmt.Println("\nAssignment Operators:")
	fmt.Println("Initial x =", x)

	x += 2
	fmt.Println("x += 2:", x)

	x -= 2
	fmt.Println("x -= 2:", x)

	x *= 2
	fmt.Println("x *= 2:", x)

	x /= 2
	fmt.Println("x /= 2:", x)

	x %= 3
	fmt.Println("x %= 3:", x)

	x &= 3
	fmt.Println("x &= 3:", x)

	x |= 4
	fmt.Println("x |= 4:", x)

	x ^= 1
	fmt.Println("x ^= 1:", x)

	x <<= 1
	fmt.Println("x <<= 1:", x)

	x >>= 1
	fmt.Println("x >>= 1:", x)
}

func main() {
	var a, b int

	fmt.Print("Enter integer a: ")
	fmt.Scan(&a)

	fmt.Print("Enter integer b: ")
	fmt.Scan(&b)

	fmt.Println("\n--- Bitwise Operators ---")
	myXor(a, b)
	myNOT(a, b)
	myOR(a, b)
	myAND(a, b)
	myShift(a, b)

	myAssignment(a)
}
