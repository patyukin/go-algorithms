package main

func main() {

}

func intersection(a, b []int) []int {
	m := make(map[int]int)
	var result []int
	for _, elem := range a {
		m[elem]++
	}

	for _, elem := range b {
		if value, ok := m[elem]; ok && value > 0 {
			result = append(result, elem)
			m[elem]--
		}
	}

	return result
}
