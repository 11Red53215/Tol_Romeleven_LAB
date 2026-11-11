/*
============================================================

	Lab #3: Password Cracking (SHA-512)

============================================================

	Question:
	  Read a wordlist file and iterate over every line using a
	  for loop. Crack the SHA-512 hash below and save the verbose
	  output to a text file.

	  SHA512: 485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38

	  The cracking logic lives in utils/crack/sha512.go and is
	  called from this file.

============================================================
*/
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"lab/utils/crack"
)

const targetSHA512 = "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"

func main() {
	labDir := "Week 3"

	// Resolve the project folder from this source file so the program also
	// works when it is started from another working directory.
	_, thisFile, _, ok := runtime.Caller(0)
	if ok {
		labDir = filepath.Dir(thisFile)
	}

	wordlist := filepath.Join(labDir, "wordlists", "nord_vpn.txt")
	outputDir := filepath.Join(labDir, "output")
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		fmt.Println("Cannot create output folder:", err)
		return
	}

	logPath := filepath.Join(outputDir, "verbose_sha512.txt")
	logFile, err := os.Create(logPath)
	if err != nil {
		fmt.Println("Cannot create verbose log:", err)
		return
	}
	defer logFile.Close()

	out := io.MultiWriter(os.Stdout, logFile)

	fmt.Println("======== SHA-512 Password Cracking ========")
	fmt.Println("Hash:", targetSHA512)
	fmt.Println("Verbose log:", logPath)
	fmt.Println()

	password, attempts, err := crack.CrackSHA512(targetSHA512, wordlist, out)
	if err != nil {
		fmt.Fprintln(out, "[-] Error:", err)
		return
	}

	fmt.Println()
	if password == "" {
		fmt.Printf("[-] Password not found after %d attempts.\n", attempts)
		return
	}

	fmt.Println("[+] PASSWORD FOUND:", password)
	fmt.Printf("[+] The hash %s belongs to %q\n", targetSHA512, password)
}
