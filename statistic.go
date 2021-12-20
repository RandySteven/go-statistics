package go_mod

type functions interface {
	Factorial(a int) int
	Combination(a *int, b *int) float32
	Permutation(a *int, b *int) int
	CirclePermutation(a *int) int
	Mean(arr []int) float32
	Median(arr []int) float32
	Sum(arr []int) int
	Pow(arr []int) int
}

type Statistic struct {
}

func (s Statistic) Factorial(a int) int {
	if a <= 1 {
		return 1
	} else {
		return a * s.Factorial(a-1)
	}
}

func (s Statistic) Combination(a *int, b *int) float32 {
	if *a < *b {
		return 0
	}
	return float32(s.Factorial(*a) / (s.Factorial(*b) * s.Factorial(*a-*b)))
}

func (s Statistic) Permutation(a *int, b *int) int {
	if *a < *b {
		return 0
	}
	return s.Factorial(*a) / s.Factorial(*a-*b)
}

func (s Statistic) CirclePermutation(a *int) int {
	return s.Factorial(*a - 1)
}

func (s Statistic) Mean(arr []int) float32 {
	n := len(arr)
	sum := 0

	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	return float32(sum / n)
}

func (s Statistic) Median(arr []int) float32 {
	n := len(arr)

	med := 0

	if n%2 == 0 {
		med = arr[n/2]
	} else {
		med = (arr[n/2] + arr[n/2+1]) / 2
	}

	return float32(med)
}

func (s Statistic) Sum(arr []int) int {
	n := len(arr)

	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	return sum
}

func (s Statistic) Pow(x *int, y *int) int {
	if *y == 0 {
		return 1
	}

	for i := 0; i < *y; i++ {
		*x *= *x
	}

	return *x
}
