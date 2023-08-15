package main

func Factorial(n uint64) (fac uint64) {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}
