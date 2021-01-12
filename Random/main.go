package Random

import (
	"math"
	"time"
)

var seed float64 = float64(time.Now().UnixNano())

// this function generates float number from range (0;1)
// values a,c,m are used in Java's java.util.Random
// from https://en.wikipedia.org/wiki/Linear_congruential_generator#Parameters_in_common_use
func RandomGenerator() float64 {
	var a float64 = 25214903917
	var c float64 = 11
	var m float64 = math.Pow(2, 48)

	// seed is recalculated each function call
	seed = math.Mod(a*seed+c, m)

	// calculated seed is divided by m to get float number in range (0;1)
	return seed / m
}
