package main

func Season(month int) string {
	var seasons = map[string][]int{
		"Winter": {1, 2, 12},
		"Spring": {3, 4, 5},
		"Summer": {6, 7, 8},
		"Autumn": {9, 10, 11},
	}

	for k, months := range seasons {
		for _, m := range months {
			if m == month {
				return k
			}
		}
	}

	return "Season unknown"
}
