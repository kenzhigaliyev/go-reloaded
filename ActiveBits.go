package student

func ActiveBits(n int) int {
	if n == 0 {
		return 0
	}
	counter := 1
	for n > 1 {
		if n%2 == 1 {
			counter++
		}
		n = n / 2
	}
	return counter
}
