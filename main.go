package main

import (
	"fmt"
	"strings"
)

func caesarDecrypt(ciphertext string, shift int32) string {
	plaintext := ""
	for _, char := range ciphertext {
		if char >= 'A' && char <= 'Z' {
			char = 'A' + (char-'A'+26-shift)%26
		} else if char >= 'a' && char <= 'z' {
			char = 'a' + (char-'a'+26-shift)%26
		}
		plaintext += string(char)
	}
	return plaintext
}

func bruteForceCaesar(ciphertext string) {
	for shift := 1; shift <= 25; shift++ {
		plaintext := caesarDecrypt(ciphertext, int32(shift))
		fmt.Printf("Shift %d: %s\n", shift, plaintext)
	}
}

func main() {
	ciphertext := "Jvtl vcly olyl Dhazvu"    // Replace with your actual ciphertext
	ciphertext = strings.ToUpper(ciphertext) // Convert to uppercase for consistency

	bruteForceCaesar(ciphertext)
}
