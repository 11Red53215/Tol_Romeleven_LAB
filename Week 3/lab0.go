/*
============================================================

	Lab #0: Proof the Hash Program

============================================================

	Question:
	  Write a Go program that takes two strings from the user and
	  hashes both of them with MD5, SHA-1, SHA-256, SHA-512 and
	  SHA3-256.
	  For every algorithm, compare Output A (hash of Input 1)
	  with Output B (hash of Input 2) and display whether the
	  two outputs match or do not match.

	  Pseudocode:
	    main() { proofMe(txt1, txt2) }

============================================================
*/
package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// proofMe hashes both inputs with every required algorithm and prints the
// comparison result.
func proofMe(input1, input2 string) {
	md5A := md5.Sum([]byte(input1))
	md5B := md5.Sum([]byte(input2))
	printResult("MD5",
		hex.EncodeToString(md5A[:]),
		hex.EncodeToString(md5B[:]))

	sha1A := sha1.Sum([]byte(input1))
	sha1B := sha1.Sum([]byte(input2))
	printResult("SHA1",
		hex.EncodeToString(sha1A[:]),
		hex.EncodeToString(sha1B[:]))

	sha256A := sha256.Sum256([]byte(input1))
	sha256B := sha256.Sum256([]byte(input2))
	printResult("SHA256",
		hex.EncodeToString(sha256A[:]),
		hex.EncodeToString(sha256B[:]))

	sha512A := sha512.Sum512([]byte(input1))
	sha512B := sha512.Sum512([]byte(input2))
	printResult("SHA512",
		hex.EncodeToString(sha512A[:]),
		hex.EncodeToString(sha512B[:]))

	sha3A := sha3.Sum256([]byte(input1))
	sha3B := sha3.Sum256([]byte(input2))
	printResult("SHA3-256",
		hex.EncodeToString(sha3A[:]),
		hex.EncodeToString(sha3B[:]))
}

// printResult prints Output A, Output B and the verdict for one algorithm.
func printResult(name, outputA, outputB string) {
	fmt.Printf("Hash (%s):\n", name)
	fmt.Printf(" Output A = %s\n", outputA)
	fmt.Printf(" Output B = %s\n", outputB)

	if outputA == outputB {
		fmt.Println(" => Match!")
	} else {
		fmt.Println(" => No Match!")
	}
	fmt.Println()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimRight(name, "\r\n")

	fmt.Printf("\n========%s + Hashing Program========\n", name)

	fmt.Print("Please input value 1: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimRight(input1, "\r\n")

	fmt.Print("Please input value 2: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimRight(input2, "\r\n")

	fmt.Println()
	proofMe(input1, input2)
}
