package e268

func missingNumber(nums []int) int {
	lastNumber := len(nums)
	s := (lastNumber / 2) * (0 + lastNumber)

	var currentSum int
	for _, num := range nums {
		currentSum += num
	}

	return s - currentSum
}
