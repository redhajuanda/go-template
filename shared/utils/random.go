package utils

import (
	"math/rand"
	"time"
)

// PickRandomString picks a random string
func PickRandomString(str ...string) string {

	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(str)
	return str[n]
}

// PickRandomInt picks a random string
func PickRandomInt(len int) int {

	rand.Seed(time.Now().Unix())
	n := rand.Int() % len
	return n
}
