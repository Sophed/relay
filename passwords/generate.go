package passwords

import (
	"crypto/rand"
	"math/big"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890+/-!&$£^*@~#:;,."

func Generate(length int) string {
	password := ""
	for range length {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		password += string(chars[n.Int64()])
	}
	return password
}
