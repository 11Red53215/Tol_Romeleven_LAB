/*
============================================================
 Lab #1: Assignment Operators
============================================================
 Question:
   Write a Go program to demonstrate the use of assignment
   operators. The program should take two integer inputs and
   perform various assignment operations such as =, +=, -=,
   *=, /=, and %=. Display the result after each operation.
============================================================
*/
package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter first integer: ")
	fmt.Scan(&a)
	fmt.Print("Enter second integer: ")
	fmt.Scan(&b)

	x := a
	fmt.Printf("Start: x = %d\n", x)

	x += b
	fmt.Printf("x += b -> %d\n", x)

	x -= b
	fmt.Printf("x -= b -> %d\n", x)

	x *= b
	fmt.Printf("x *= b -> %d\n", x)

	if b != 0 {
		x /= b
		fmt.Printf("x /= b -> %d\n", x)
		x %= b
		fmt.Printf("x %%= b -> %d\n", x)
	} else {
		fmt.Println("Skipping /= and %= (division by zero)")
	}

	x = a
	fmt.Printf("x = a -> %d\n", x)
}