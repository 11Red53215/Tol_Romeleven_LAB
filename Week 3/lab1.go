/*
============================================================

	Lab #1: Password Cracking (MD5)

============================================================

	Question:
	  Read a wordlist file and iterate over every line using a
	  for loop. Crack the MD5 hash below and save the verbose
	  output to a text file.

	  MD5: 6a85dfd77d9cb35770c9dc6728d73d3f

	  The cracking logic lives in utils/crack/md5.go and is
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

const targetMD5 = "6a85dfd77d9cb35770c9dc6728d73d3f"

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

	logPath := filepath.Join(outputDir, "verbose_md5.txt")
	logFile, err := os.Create(logPath)
	if err != nil {
		fmt.Println("Cannot create verbose log:", err)
		return
	}
	defer logFile.Close()

	out := io.MultiWriter(os.Stdout, logFile)

	fmt.Println("======== MD5 Password Cracking ========")
	fmt.Println("Hash:", targetMD5)
	fmt.Println("Verbose log:", logPath)
	fmt.Println()

	password, attempts, err := crack.CrackMD5(targetMD5, wordlist, out)
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
	fmt.Printf("[+] The hash %s belongs to %q\n", targetMD5, password)
}
