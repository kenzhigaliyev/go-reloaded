package functions

func Atoi(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	result := 0
	positive := true
	number := []rune{}
	sign := 1
	for _, char := range s {
		number = append(number, char)
	}
	if number[0] == '-' {
		for i := 1; i < length; i++ {
			if number[i] < '0' || number[i] > '9' {
				return 0
			}
		}
		sign = -1
		positive = false
	} else if number[0] == '+' || (number[0] >= '0' && number[0] <= '9') {
		for i := 1; i < length; i++ {
			if number[i] < '0' || number[i] > '9' {
				return 0
			}
		}
	} else {
		return 0
	}
	for i := 0; i < length; i++ {
		if number[i] >= '1' && number[i] <= '9' {
			if (length - i) > 19 {
				return 0
			}
			for j := i; j < length; j++ {
				result = (result * 10) + (int(number[j]) - 48)
			}
			result = sign * result
			if result < 0 && positive {
				return 0
			} else if result > 0 && !positive {
				return 0
			} else {
				return result
			}
		}
	}
	return result
}

func AtoiUnix(s string) (uint64, bool) {
	length := len(s)
	minus := true
	result := uint64(0)

	if length == 0 {
		return result, minus
	}
	number := []rune{}
	for _, char := range s {
		number = append(number, char)
	}
	if number[0] == '+' {
		for i := 1; i < length; i++ {
			if number[i] < '0' || number[i] > '9' {
				return result, minus
			}
		}
		// sign = -1
		minus = false
	} else if number[0] == '-' || (number[0] >= '0' && number[0] <= '9') {
		for i := 1; i < length; i++ {
			if number[i] < '0' || number[i] > '9' {
				return result, minus
			}
		}
	} else {
		return result, minus
	}
	for i := 0; i < length; i++ {
		if number[i] >= '1' && number[i] <= '9' {
			result = (result * uint64(10)) + (uint64(number[i]) - 48)
			// fmt.Println(result, number[i])
		}
	}
	return result, minus
}
