package ex1

func prefixSum(s string) []int {
	ps := make([]int, len(s))
	sum := 0

	for i, char := range s {
		sum += int(char - '0')
		ps[i] = sum
	}

	return ps
}
