package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	whoToGreet := "World"

	file := os.Stdin
	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat()", err)
	}
	// Check if anything in stdin
	size := fi.Size()
	if size > 0 {
		scanner := bufio.NewScanner(os.Stdin)
		// Set the split function for the scanning operation.
		scanner.Split(bufio.ScanWords)
		// Read first word (space-delim'd token)
		if scanner.Scan() {
			whoStdin := scanner.Text()
			// Ensure word read is 26 runes or less in length
			if l := len([]rune(whoStdin)); l <= 26 {
				whoToGreet = whoStdin
			} else {
				fmt.Fprintln(os.Stderr, "max length of who to greet is 26, your string length:", l)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
	fmt.Print(greet(whoToGreet))
}

func greet(who string) string {
	return "Hello" + " " + who
}