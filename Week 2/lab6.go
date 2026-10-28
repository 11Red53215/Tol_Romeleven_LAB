/*
============================================================
 Lab #6: XOR Encryption/Decryption
============================================================
 Question:
   Write a Go program to perform simple XOR encryption and
   decryption.
   The program should take a plaintext message and a key
   (single byte or repeating key) and produce a ciphertext
   by XORing them.
   Use the same function to decrypt the ciphertext to
   recover the original message.

   func xorEncrypt(text string, key byte) string {
       // Code here
   }
============================================================
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// XORs each byte of text with a repeating key.
func xorEncrypt(text string, key byte) string {
	out := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		out[i] = text[i] ^ key
	}
	return string(out)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter plaintext: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")

	fmt.Print("Enter single-byte key (e.g. 'K'): ")
	keyStr, _ := reader.ReadString('\n')
	keyStr = strings.TrimRight(keyStr, "\r\n")
	if keyStr == "" {
		fmt.Println("Error: empty key")
		return
	}
	key := keyStr[0]

	cipher := xorEncrypt(text, key)
	fmt.Printf("Ciphertext (hex)  : % x\n", []byte(cipher))

	plain := xorEncrypt(cipher, key)
	fmt.Println("Decrypted plaintext:", plain)
}