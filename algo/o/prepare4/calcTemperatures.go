package prepare4

// Input: [30,60,90]
// Output: [1,1,0]

// Input: [73,74,75,71,69,72,76,73]
// Output: [1,1,4,2,1,1,0,0]

// https://leetcode.com/problems/daily-temperatures
func calcTemperatures(t []int) []int {
	var result []int

	for i := 0; i < len(t)-1; i++ {
		count := 0
		for j := i + 1; j < len(t); j++ {
			if t[j] > t[i] {
				count = j - i
				break
			}
		}

		result = append(result, count)
	}

	return append(result, 0)
}

func calcTemperaturesWithoutSort(t []int) []int {
	return []int{}
}
