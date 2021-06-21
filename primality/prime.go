package primality

import (
	"fmt"
	"math/rand"
)

var smallPrimes = map[int]bool{2: true, 3: true, 5: true, 7: true, 11: true}

func randRangeInt(min, max int) int {
	// Move interval to [0, n)
	if max <= min {
		panic(fmt.Sprintf("[%d %d) is not a valid in-zer0 interval", min, max))
	}
	return rand.Intn(max-min) + min
}

// fermatsProbablyPrimeTest a crude probably prime test for n. For a randomly chosen a < n ,
// if a^(n-1) != 1 (mod n)  => n is composite
func fermatsProbablyPrimeTest(n int) bool {

	a := randRangeInt(1, n)
	res := a
	for i := 2; i < n; i++ {
		res = (res * a) % n
	}

	return res == 1

}

// ProbablyPrime returns where the given value is, statistically, likely to be prime
func ProbablyPrime(n, samples int) bool {

	if n < 1 {
		panic(fmt.Sprintf("%v is not a > 1", n))
	}

	// Easy and quick to compute
	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	if smallPrimes[n] {
		return true
	}

	// read results from channel
	for i := 0; i < samples; i++ {
		if !fermatsProbablyPrimeTest(n) {
			return false
		}
	}

	// Could not determine if composite, probably prime
	return true
}
