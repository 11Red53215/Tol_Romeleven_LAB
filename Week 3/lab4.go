/*
============================================================

	Lab #4: Mini-CTF (Cat Leak)

============================================================

	Story:
	  The victim loves cats and always searches cat-related
	  things on Google. A login URL leaked two p1crypt values,
	  and the Google search URL leaked the same value as a
	  percent-encoded query.

	  Each p1crypt parameter is only half of a 32-character
	  MD5 digest. Join them, then test cat-themed candidates
	  and the supplied wordlist until the digest matches.

	  Flag format: cryptoCTF{answer_here}

============================================================
*/
package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"lab/utils/crack"
)

const loginURL = "https://securesystem.local:8890/login.php?userId=168&p1crypt=b6ccb4ece5454dca&p1crypt=e51778b3e239ebc2"

const googleURL = "https://www.google.com/search?q=%62%36%63%63%62%34%65%63%65%35%34%35%34%64%63%61%65%35%31%37%37%38%62%33%65%32%33%39%65%62%63%32&rlz=1C5CHF"

func main() {
	labDir := "Week 3"

	_, thisFile, _, ok := runtime.Caller(0)
	if ok {
		labDir = filepath.Dir(thisFile)
	}

	fmt.Println("======== Mini-CTF: Cat Leak ========")

	// 1. Read both p1crypt values from the leaked login URL.
	parsedLogin, err := url.Parse(loginURL)
	if err != nil {
		fmt.Println("Cannot parse login URL:", err)
		return
	}
	halves := parsedLogin.Query()["p1crypt"]
	if len(halves) != 2 {
		fmt.Println("Expected two p1crypt parameters, got", len(halves))
		return
	}
	reconstructed := strings.Join(halves, "")

	// 2. Read the percent-encoded q value from the Google URL.
	parsedGoogle, err := url.Parse(googleURL)
	if err != nil {
		fmt.Println("Cannot parse Google URL:", err)
		return
	}
	googleHash := parsedGoogle.Query().Get("q")

	fmt.Println("p1crypt part 1 :", halves[0])
	fmt.Println("p1crypt part 2 :", halves[1])
	fmt.Println("Reconstructed  :", reconstructed, "(32 hex chars = 128-bit MD5)")
	fmt.Println("Google q value :", googleHash)

	if !strings.EqualFold(googleHash, reconstructed) {
		fmt.Println("[-] The two leaks do not contain the same hash.")
		return
	}
	fmt.Println("[+] Both leaks contain the same MD5 digest.")

	// The story claims bcrypt, but bcrypt hashes are text strings that start
	// with a cost prefix such as "$2a$10$". A 32-character hex value is MD5.
	fmt.Println("[!] Note: a real bcrypt hash starts with $2a$, $2b$ or $2y$,")
	fmt.Println("    so the developer actually used MD5, not bcrypt.")

	// 3. Build candidates: the wordlist plus cat-themed guesses and pairs.
	wordlist := filepath.Join(labDir, "wordlists", "nord_vpn.txt")
	candidates, err := buildCandidates(wordlist)
	if err != nil {
		fmt.Println("[!] Wordlist not loaded:", err)
	}

	fmt.Printf("\nSearching %d cat-themed candidates...\n", len(candidates))

	password := ""
	for _, candidate := range candidates {
		if crack.MD5Hex(candidate) == strings.ToLower(reconstructed) {
			password = candidate
			break
		}
	}

	if password == "" {
		fmt.Println("[-] Password not found.")
		return
	}

	fmt.Println("[+] Password cracked:", password)
	fmt.Println("[+] Flag: cryptoCTF{" + password + "}")
}

// buildCandidates loads every line of the wordlist and adds cat-related
// single words and two-word combinations (for example "meow" + "meow").
func buildCandidates(wordlistPath string) ([]string, error) {
	var candidates []string
	seen := make(map[string]bool)

	add := func(word string) {
		word = strings.TrimSpace(word)
		if word == "" || seen[word] {
			return
		}
		seen[word] = true
		candidates = append(candidates, word)
	}

	baseWords := []string{
		"cat", "cats", "kitty", "kitten", "meow", "purr", "purrs", "paw",
		"paws", "feline", "whiskers", "tabby", "tomcat", "pussycat",
		"moggie", "miau", "miao", "nyan", "garfield", "sphynx",
	}

	for _, word := range baseWords {
		add(word)
		add(strings.ToUpper(word[:1]) + word[1:])
	}

	// Cat-related two-word combinations, both lowercase and capitalised.
	for _, first := range baseWords {
		for _, second := range baseWords {
			add(first + second)
			add(strings.ToUpper(first[:1]) + first[1:] + second)
		}
	}

	file, err := os.Open(wordlistPath)
	if err != nil {
		return candidates, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		add(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return candidates, err
	}

	return candidates, nil
}
