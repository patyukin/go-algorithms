package u2050689

func isLeapYear(y int) string {
	if y%4 == 0 && (y%100 != 0 || y%400 == 0) {
		return "YES"
	}

	return `NO`
}
