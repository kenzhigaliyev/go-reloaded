package functions

import (
	"github.com/01-edu/z01"
)

var newcounter = 0

func PrintCombN(n int) {
	if n < 0 {
		return
	}
	counter := 0
	combination := ""
	CombNum(0, n, combination, counter, newcounter)
	z01.PrintRune('\n')

}

func CombNum(first int, num int, comb string, counter int, newcounter int) {
	// newcounter := 0
	if num != 0 {
		for i := first; i <= 9; i++ {
			numbers := comb + (string(i + 48))
			counter++
			CombNum(i+1, num-1, numbers, counter, newcounter)
		}

	} else if num == 0 {
		// fmt.Println(newcounter)
		for _, char := range comb {
			z01.PrintRune(char)
			newcounter = newcounter + 1

		}
		// if newcounter < counter {
		z01.PrintRune(',')
		z01.PrintRune('!')
		// }
	}
	// fmt.Print(counter)
	// fmt.Print("---")
	// fmt.Print(newcounter)
}
