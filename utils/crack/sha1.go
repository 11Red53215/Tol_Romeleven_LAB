package crack

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// SHA1Hex returns the lowercase SHA-1 hash of text.
func SHA1Hex(text string) string {
	sum := sha1.Sum([]byte(text))
	return hex.EncodeToString(sum[:])
}

// CrackSHA1 reads a wordlist line by line and compares the SHA-1 hash of every
// word with target. It writes the verbose progress to out and returns the
// matching word and the number of words it tried.
func CrackSHA1(target, wordlistPath string, out io.Writer) (string, int, error) {
	target = strings.ToLower(strings.TrimSpace(target))

	file, err := os.Open(wordlistPath)
	if err != nil {
		return "", 0, err
	}
	defer file.Close()

	fmt.Fprintf(out, "Target SHA1 : %s\n", target)
	fmt.Fprintf(out, "Wordlist    : %s\n\n", wordlistPath)

	scanner := bufio.NewScanner(file)
	attempts := 0
	started := time.Now()

	for scanner.Scan() {
		word := strings.TrimRight(scanner.Text(), "\r")
		if word == "" {
			continue
		}

		attempts++
		got := SHA1Hex(word)
		fmt.Fprintf(out, "[%04d] %-28q => %s\n", attempts, word, got)

		if got == target {
			fmt.Fprintf(out, "\n[+] Cracked after %d attempt(s) in %s\n",
				attempts, time.Since(started).Round(time.Millisecond))
			fmt.Fprintf(out, "[+] SHA1(%q) = %s\n", word, got)
			return word, attempts, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", attempts, err
	}

	fmt.Fprintf(out, "\n[-] No match found after %d attempt(s) in %s\n",
		attempts, time.Since(started).Round(time.Millisecond))
	return "", attempts, nil
}
