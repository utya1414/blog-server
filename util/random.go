package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(len(alphabet))]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomEmail() string {
	return RandomString(6) + "@" + RandomString(4) + ".com"
}

func RandomUsername() string {
	return RandomString(6)
}

func RandomPassword() string {
	return RandomString(6)
}
