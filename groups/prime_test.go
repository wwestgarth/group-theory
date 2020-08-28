package groups

import "testing"

func TestIsntPrime(t *testing.T) {

	testValues := []int{9, 15, 27, 6319}
	for _, testValue := range testValues {
		if ProbablyPrime(testValue, 10) {
			t.Errorf("Expected %d to be composite", testValue)
		}
	}

}

func TestIsPrime(t *testing.T) {

	testValues := []int{5, 11, 23, 313, 2017, 7901}
	for _, testValue := range testValues {
		if !ProbablyPrime(testValue, 10) {
			t.Errorf("Expected %d to be prime", testValue)
		}
	}
}

func TestEvens(t *testing.T) {

	testValues := []int{4, 6, 122, 1024}

	for _, testValue := range testValues {
		if ProbablyPrime(testValue, 10) {
			t.Errorf("Expected %d to be composite", testValue)
		}
	}

}
