package go_mod

func Factorial(a int) int {
	if a <= 1 {
		return 1
	} else {
		return a * Factorial(a-1)
	}
}

func Combination(a *int, b *int) int {
	if *a < *b {
		return 0
	}
	return Factorial(*a) / (Factorial(*b) * Factorial(*a-*b))
}

func Permutation(a *int, b *int) int {
	if *a < *b {
		return 0
	}
	return Factorial(*a) / Factorial(*a-*b)
}

func CirclePermutation(a *int) int {
	return Factorial(*a - 1)
}
