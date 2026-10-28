/*
============================================================
 Lab #4: Mini Calculator
============================================================
 Question:
   Write a Go program to implement a mini calculator using
   functions. The program should read two numbers and an
   operator from the user, perform the operation, and
   return/print the result. Support at least the operators:
   +, -, *, /, %.
   Handle division by zero and invalid operators gracefully.
   (Optional) Make it menu-driven and repeat until the user
   chooses Exit.

   Example:
     ===== Mini Calculator =====
     1) Add 2) Sub 3) Mul 4) Div 5) Mod 6) Exit
     Choose: 4
     Enter a: 10
     Enter b: 0
     Error: division by zero
============================================================
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(a, b float64) float64 { return a + b }
func sub(a, b float64) float64 { return a - b }
func mul(a, b float64) float64 { return a * b }

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func mod(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("modulo by zero")
	}
	return float64(int(a) % int(b)), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n===== Mini Calculator =====")
		fmt.Println("1) Add 2) Sub 3) Mul 4) Div 5) Mod 6) Exit")
		fmt.Print("Choose: ")
		line, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(line)

		if choice == "6" {
			fmt.Println("Bye!")
			return
		}

		fmt.Print("Enter a: ")
		lineA, _ := reader.ReadString('\n')
		a, errA := strconv.ParseFloat(strings.TrimSpace(lineA), 64)

		fmt.Print("Enter b: ")
		lineB, _ := reader.ReadString('\n')
		b, errB := strconv.ParseFloat(strings.TrimSpace(lineB), 64)

		if errA != nil || errB != nil {
			fmt.Println("Error: invalid number")
			continue
		}

		switch choice {
		case "1":
			fmt.Printf("Result: %v\n", add(a, b))
		case "2":
			fmt.Printf("Result: %v\n", sub(a, b))
		case "3":
			fmt.Printf("Result: %v\n", mul(a, b))
		case "4":
			r, err := div(a, b)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %v\n", r)
			}
		case "5":
			r, err := mod(a, b)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %v\n", r)
			}
		default:
			fmt.Println("Error: invalid operator")
		}
	}
}