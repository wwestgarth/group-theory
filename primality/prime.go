package primality

import (
	"fmt"
	"math/rand"
)

var (
	smallPrimes = map[int]bool{1: true, 2: true, 3: true, 5: true, 7: true, 11: true}
)

// getRandom returns a pseudo-random int in the range (min, max)
func getRandom(min, max int) int {
	if max <= min {
		// Really shouldn't get here
		panic(fmt.Sprintf("[%d %d) is not a valid in-zer0 interval", min, max))
	}

	// Move interval to [0, n)
	return rand.Intn(max-min) + min
}

// fermatsProbablyPrimeTest a crude probably prime test for n. For a randomly chosen a < n ,
// if a^(n-1) != 1 (mod n)  => n is composite
func fermatsProbablyPrime(n int) bool {

	a := getRandom(1, n)
	pow := a

	// Calcualte powers of a
	for i := 2; i < n; i++ {
		pow = (pow * a) % n
	}

	return pow == 1

}

// ProbablyPrime returns where the given value is, statistically, likely to be prime
func ProbablyPrime(n, nSamples int) bool {

	if n%2 == 0 {
		return false
	}
	if smallPrimes[n] {
		return true
	}

	for i := 0; i < nSamples; i++ {
		if !fermatsProbablyPrime(n) {
			return false
		}
	}

	// Could not determine if composite, probably prime
	return true
}
