/*
============================================================
 Lab #3: Bitwise and Assignment Operators
============================================================
 Question:
   Write a Go program to demonstrate the use of bitwise and
   assignment operators. The program should perform AND, OR,
   XOR, NOT, left shift, and right shift operations on two
   integers using functions and display the results. Also,
   show the effect of assignment operators on variable values.

   Pseudocode:
     main() {
       myXor(a,b)
       myNOT(a,b)
       myOR(a,b)
       myAND(a,b)
     }
============================================================
*/
package main

import "fmt"

func myAND(a, b int) int { return a & b }
func myOR(a, b int) int  { return a | b }
func myXOR(a, b int) int { return a ^ b }
func myNOT(a int) int    { return ^a }
func myLeftShift(a, n int) int  { return a << n }
func myRightShift(a, n int) int { return a >> n }

func main() {
	var a, b int
	fmt.Print("Enter first integer: ")
	fmt.Scan(&a)
	fmt.Print("Enter second integer: ")
	fmt.Scan(&b)

	fmt.Printf("a = %d (%b), b = %d (%b)\n", a, a, b, b)
	fmt.Printf("AND  a & b  = %d (%b)\n", myAND(a, b), myAND(a, b))
	fmt.Printf("OR   a | b  = %d (%b)\n", myOR(a, b), myOR(a, b))
	fmt.Printf("XOR  a ^ b  = %d (%b)\n", myXOR(a, b), myXOR(a, b))
	fmt.Printf("NOT  ^a     = %d\n", myNOT(a))
	fmt.Printf("Left shift  a << 1 = %d\n", myLeftShift(a, 1))
	fmt.Printf("Right shift a >> 1 = %d\n", myRightShift(a, 1))

	// Assignment operators
	c := a
	fmt.Printf("\nAssignment operators (c starts at %d):\n", c)
	c &= b
	fmt.Printf("c &= b -> %d\n", c)
	c = a
	c |= b
	fmt.Printf("c |= b -> %d\n", c)
	c = a
	c ^= b
	fmt.Printf("c ^= b -> %d\n", c)
	c = a
	c <<= 1
	fmt.Printf("c <<= 1 -> %d\n", c)
	c = a
	c >>= 1
	fmt.Printf("c >>= 1 -> %d\n", c)
}