package functions

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	if n < 1 || n > 9 {
		return
	}
	numbers := make([]int, n)

	for index := 0; index < n; index++ {
		numbers[index] = index
	}

	for index := 0; index < len(numbers); index++ {
		z01.PrintRune(rune(numbers[index] + 48))
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')

	for {
		if numbers[len(numbers)-1] != 9 {
			numbers[len(numbers)-1]++
			for index := 0; index < len(numbers); index++ {
				z01.PrintRune(rune(numbers[index] + 48))
			}
			if len(numbers) == 1 && numbers[len(numbers)-1] == 9 {
				break
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')
		} else if numbers[0] == 10-n {
			break
		} else {
			numbers = Comb(numbers)
			for index := 0; index < len(numbers); index++ {
				z01.PrintRune(rune(numbers[index] + 48))
			}
			if 10-n == numbers[0] && 9 == numbers[n-1] {
				break
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func Comb(numbers []int) []int {
	num1, num2 := len(numbers)-2, len(numbers)-1
	for {
		if numbers[num1] < numbers[num2] && numbers[num1] != numbers[num2]-1 {
			numbers[num1]++
			numbers[num2] = numbers[num1] + 1
			if num2 == len(numbers)-1 {
				break
			}
			result := numbers[num2] + 1
			for index := num2 + 1; index < len(numbers); index++ {
				numbers[index] = result
				result++
			}
			break
		} else if numbers[num1] == numbers[num2]-1 {
			if num1 < 1 {
				break
			}
			num1, num2 = num1-1, num2-1
		}
	}
	return numbers
}
