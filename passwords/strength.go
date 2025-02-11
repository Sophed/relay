package passwords

import "unicode"

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

func hasNumbers(password string) bool {
	for _, char := range password {
		if unicode.IsNumber(char) {
			return true
		}
	}
	return false
}

func hasSymbols(password string) bool {
	for _, char := range password {
		if unicode.IsSymbol(char) {
			return true
		}
	}
	return false
}

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
