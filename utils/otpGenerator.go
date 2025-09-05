package utils

import (
	"math/rand"
	"time"
)

func Generate6DigitCode() int {
	rand.Seed(time.Now().UnixNano())
	return 100000 + rand.Intn(900000)
}
