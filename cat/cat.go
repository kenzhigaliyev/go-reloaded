package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		os.Stdout.ReadFrom(os.Stdin)
	} else {
		for _, value := range arguments {
			info, err := ioutil.ReadFile(value)
			if err != nil {
				for _, val := range "cat: " + err.Error()[5:] {
					z01.PrintRune(val)
				}
				z01.PrintRune('\n')
			} else {
				for _, val := range info {
					z01.PrintRune(rune(val))
				}
			}
		}
	}
}
