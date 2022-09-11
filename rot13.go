package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Enter a string to encrypt: ") // Get input from user
	fmt.Scan(&input)
	fmt.Println("Encrypted string is: ", rot13(input))
}

func rot13(s string) string {
	var output string
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			output += string((c-'A'+13)%26 + 'A') //  Rot13 encryption
		} else if c >= 'a' && c <= 'z' {
			output += string((c-'a'+13)%26 + 'a')
		} else {
			output += string(c)
		}
	}
	return output // return encrypted string
}

// Language: go

// Path: rot13.go

// ROT13 cipher
