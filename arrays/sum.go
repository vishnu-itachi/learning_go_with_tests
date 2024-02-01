package arrays

func Sum(l [5]int) int {
	sum := 0

	for i := 0; i < 5; i++ {
		sum += l[i]
	}
	return sum
}

func RefactoredSum(l [5]int) int {
	sum := 0
	for _, number := range l {
		sum += number
	}
	return sum
}

func SumSlice(l []int) int {
	sum := 0

	for _, number := range l {
		sum += number
	}

	return sum
}

func SumAll(l ...[]int) []int {
	length := len(l)
	a := make([]int, length)

	for i , numbers := range l{
		a[i] = SumSlice(numbers)
	}

	return a
}
