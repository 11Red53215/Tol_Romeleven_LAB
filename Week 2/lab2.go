/*
============================================================
 Lab #2: Logical Operators
============================================================
 Question:
   Write a Go program to demonstrate the use of logical
   operators such as &&, ||, and !.
   The program should take two integer inputs and evaluate
   logical expressions like:
     - both positive (&&)
     - one greater than the other (||)
     - not equal (!)
   Display whether each condition is true or false.
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

	bothPositive := a > 0 && b > 0
	oneGreater := a > b || b > a
	notEqual := a != b

	fmt.Printf("a=%d b=%d\n", a, b)
	fmt.Printf("Both positive (a>0 && b>0): %v\n", bothPositive)
	fmt.Printf("One greater than the other (a>b || b>a): %v\n", oneGreater)
	fmt.Printf("Not equal (!(a==b)): %v\n", notEqual)
}