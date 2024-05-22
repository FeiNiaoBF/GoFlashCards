package util

import (
	"math/rand"
	"strings"
	"time"
)

var randomNum *rand.Rand

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	seed := time.Now().UnixNano()
	randomNum = rand.New(rand.NewSource(seed))
}

func RandText(num int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < num; i++ {
		c := alphabet[randomNum.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
