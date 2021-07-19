package functions

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n < 0 {
		return
	}
	combination := ""
	CombNum(0, n, combination)
	z01.PrintRune('\n')

}

func CombNum(first int, num int, comb string) {
	// str:=", "
	if num != 0 {
		for i := first; i <= 9; i++ {
			// if i==num{
			// 	numbers:=comb+(string(i+48))+str
			// 	CombNum(i+1,num-1,numbers)
			// } else {
			numbers := comb + (string(i + 48))
			CombNum(i+1, num-1, numbers)
			// }
		}
	} else if num == 0 {
		for _, char := range comb {
			z01.PrintRune(char)
		}
		z01.PrintRune(',')
		z01.PrintRune(' ')

	}
}
