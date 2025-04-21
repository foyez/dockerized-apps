package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for range n {
		c := alphabet[rng.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomText generates a random text
func RandomText() string {
	return fmt.Sprintf("%s %s %s.", RandomString(6), RandomString(4), RandomString(5))
}
