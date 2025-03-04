package passwords

import "unicode"

// Strength evaluates the strength of a password
func Strength(password string) int {
	total := 0
	if len(password) >= 8 {
		total += 1
	}
	if !hasRepeatingChars(password) {
		total += 1
	}
	if hasNumbers(password) {
		total += 1
	}
	if hasSymbols(password) {
		total += 1
	}
	return total
}

// hasNumbers returns true if the password contains numbers
func hasNumbers(password string) bool {
	for _, char := range password {
		if unicode.IsNumber(char) {
			return true
		}
	}
	return false
}

// hasSymbols returns true if the password contains symbols
func hasSymbols(password string) bool {
	for _, char := range password {
		if unicode.IsSymbol(char) {
			return true
		}
	}
	return false
}

// hasRepeatingChars returns true if the password contains any repeating characters
func hasRepeatingChars(password string) bool {
	for i, char := range password {
		if i == 0 {
			continue
		}
		if byte(char) == password[i-1] {
			return true
		}
	}
	return false
}
