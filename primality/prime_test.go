package primality

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProbablyPrimeComposite(t *testing.T) {

	testValues := []int{9, 15, 27, 6319, 4, 6, 122, 1024}

	for _, testValue := range testValues {
		assert.False(t, ProbablyPrime(testValue, 10))
	}
}

func TestProbablyPrime(t *testing.T) {

	testValues := []int{5, 11, 23, 313, 2017, 7901}

	for _, testValue := range testValues {
		assert.True(t, ProbablyPrime(testValue, 10))
	}
}
