package main

func findSum(s []int, target int) []int {
	m := make(map[int]int)
	for idx, value := range s {
		result, ok := m[target-value]
		if !ok {
			m[value] = idx
			continue
		}

		return []int{result, idx}
	}

	return nil
}
