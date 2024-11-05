package test

import (
	"awesomeProject/functions"
	"testing"
)

func TestThatGetFactorSizeOfTwoIsEqualToTwo(t *testing.T) {
	number := 2
	numberOfFactor := functions.GetFactorSize(number)
	if numberOfFactor != 2 {
		t.Errorf("Get Factor size not meant to be more than 2 expected is %b", numberOfFactor)
	}
}

func TestThatGetFactorSizeOfThreeIsEqualToTwo(t *testing.T) {
	number := 3
	numberOfFactor := functions.GetFactorSize(number)
	if numberOfFactor != 2 {
		t.Errorf("Get Factor size not meant to be more than 2 expected is %b", numberOfFactor)
	}
}

func TestGetFactorOfTenIsGreaterThanTwo(t *testing.T) {
	number := 10
	numberOfFactor := functions.GetFactorSize(number)
	if numberOfFactor <= 2 {
		t.Errorf("Get Factor size not meant to be less than or equal 2 expected is %b", numberOfFactor)
	}
}

func TestThatTwoIsAPrime(t *testing.T) {
	number := 2
	isPrime := functions.IsPrime(number)
	if !isPrime {
		t.Errorf("Two is a prime number")
	}
}

func TestThatThreeIsAPrime(t *testing.T) {
	number := 3
	isPrime := functions.IsPrime(number)
	if !isPrime {
		t.Errorf("Three is a prime number")
	}
}

func TestThatFiveIsAPrimeNumber(t *testing.T) {
	number := 5
	isPrime := functions.IsPrime(number)
	if !isPrime {
		t.Errorf("Five is not a prime number %v", isPrime)
	}
}

func TestThatSixIsNotAPrimeNumber(t *testing.T) {
	number := 6
	isPrime := functions.IsPrime(number)
	if isPrime {
		t.Errorf("Six is not a prime number %v", isPrime)
	}
}

func TestThatOneIsNotPrimeNumber(t *testing.T) {
	number := 1
	isPrime := functions.IsPrime(number)
	if isPrime {
		t.Errorf("One is not a prime number %v", isPrime)
	}
}

func TestThatTwoIsPalindrome(t *testing.T) {
	number := 2
	isPalindrome := functions.IsPalindromic(number)
	if !isPalindrome {
		t.Errorf("Two is a palindrome number")
	}
}

func TestThatOneTwentyOneIsPalindromic(t *testing.T) {
	number := 121
	isPalindrome := functions.IsPalindromic(number)
	if !isPalindrome {
		t.Errorf("One Twenty One is a palindrome number")
	}
}

func TestThatZeroIsNotAPalindromic(t *testing.T) {
	number := 0
	isPalindrome := functions.IsPalindromic(number)
	if isPalindrome {
		t.Errorf("Zero is not a palindrome number")
	}
}

func TestThatTwoTwentyTwoIsAPalindromic(t *testing.T) {
	number := 222
	isPalindrome := functions.IsPalindromic(number)
	if !isPalindrome {
		t.Errorf("Two Twenty Two is a palindrome number")
	}
}
