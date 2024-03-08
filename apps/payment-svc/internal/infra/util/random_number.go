package util

import (
	"math/rand"
	"time"
)

func RandomNumber() int {
	// Adjust probabilities according to your requirements
	rand.Seed(time.Now().UnixNano())
	probability := rand.Float64() // generates a float between 0.0 and 1.0

	if probability <= 0.8 {
		return 1 // 80% chance
	} else {
		return 2 // 20% chance
	}
}
