package student

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if f(a[i], a[j]) == 1 {
				// val := a[i]
				// a[i] = a[j]
				// a[j] = val
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func Compare(a, b string) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	}
	return -1
	// return strings.Compare(b, a)
}

func Compare1(b, a string) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	}
	return -1
}
