package main

func insertionSort(s []int) {
	for i := 1; i < len(s); i++ {
		current := s[i]
		j := i - 1
		for j >= 0 && s[j] > current {
			s[j+1] = s[j]
			j--
		}

		s[j+1] = current
	}
}
