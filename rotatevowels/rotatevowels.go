package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	vowels := []rune{}
	letters := []rune{}
	for index, char1 := range arg {
		for _, char2 := range char1 {
			letters = append(letters, char2)
			if char2 == 'A' || char2 == 'a' || char2 == 'E' || char2 == 'e' || char2 == 'I' || char2 == 'i' || char2 == 'O' || char2 == 'o' || char2 == 'U' || char2 == 'u' {
				vowels = append(vowels, char2)
			}
		}
		if index != len(arg)-1 {
			letters = append(letters, ' ')
		}
	}
	length := len(vowels) - 1
	for i := 0; i < len(letters); i++ {
		if letters[i] == 'A' || letters[i] == 'a' || letters[i] == 'E' || letters[i] == 'e' || letters[i] == 'I' || letters[i] == 'i' || letters[i] == 'O' || letters[i] == 'o' || letters[i] == 'U' || letters[i] == 'u' {
			z01.PrintRune(vowels[length])
			length--
		} else {
			z01.PrintRune(letters[i])
		}
	}
	z01.PrintRune('\n')
}
