package helpers

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Rand(min, max int) int {
	return rand.Intn(max-min) + min
}
