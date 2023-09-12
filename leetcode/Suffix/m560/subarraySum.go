package m560

func subarraySum(nums []int, k int) int {
	m := map[int]int{0: 1}
	var currentSum int
	var count int
	for _, num := range nums {
		currentSum += num
		if _, ok := m[currentSum-k]; ok {
			count += m[currentSum-k]
		}
		m[currentSum]++
	}

	return count
}