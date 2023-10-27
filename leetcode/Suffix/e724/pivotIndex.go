package e724

func pivotIndex(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}

	var pxSum int
	for i, num := range nums {
		if pxSum == sum-num-pxSum {
			return i
		}

		pxSum += num
	}

	return -1
}
