package u205068

func happyTicket(ticket int) string {
	left := ticket % 1000
	right := ticket / 1000

	var getSum func(num int) int
	getSum = func(num int) int {
		a := num % 10
		num /= 10
		b := num % 10
		c := num / 10

		return a + b + c
	}

	if getSum(left) == getSum(right) {
		return "YES"
	}

	return "NO"
}
