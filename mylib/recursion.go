package mylib

func Factorial(n int) int {
	// Base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}
	// Recursive case: factorial of n is n * factorial(n - 1)
	return n * Factorial(n-1)
}

func SumOfDigits(n int) int {
	// Base Case
	if n < 10 {
		return n
	}

	return n%10 + SumOfDigits(n/10)
}
