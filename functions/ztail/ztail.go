package main

import (
	"fmt"
	"io"
	"os"
	"student/functions"

	"github.com/01-edu/z01"
)

var errorlar = true
var filedar = true

func main() {
	arguments := os.Args[1:]
	arg, number, arr, minus := IsValid(arguments)
	// fmt.Println(number, minus)

	if arg == false {
		fmt.Printf("tail: option requires an argument -- 'c'\nTry 'tail --help' for more information.\n")
		return
	}
	for index, value := range arr {
		file, err := os.Open(value)
		if err != nil {
			fmt.Printf("tail: cannot open '%s' for reading: No such file or directory", value)
			z01.PrintRune('\n')
			errorlar = false
			if index == len(arr) {
				os.Exit(0)
			} else {
				continue
			}
		}
		info, err := file.Stat()
		if err != nil {
		}
		length := uint64(info.Size())
		defer file.Close()
		data := make([]byte, length)
		for {
			_, err := file.Read(data)
			if err == io.EOF { // если конец файла
				break // выходим из цикла
			}
		}
		// var counter int
		// fmt.Println(number, minus)
		if len(arr) != 1 {
			if !errorlar && !filedar {
				z01.PrintRune('\n')
			}
			filedar = false
			fmt.Printf("==> %s <==", value)
			z01.PrintRune('\n')
			if number > length {
				if minus {
					fmt.Printf(string(data))
				} else {
					fmt.Printf("")
				}
			} else {
				if !minus {
					fmt.Print(string(data[number-1 : length]))
				} else {
					fmt.Print(string(data[length-number : length]))
				}
			}
			if index != len(arr)-1 {
				// z01.PrintRune('\n')
				// if !errorlar {
				// 	// z01.PrintRune('\n')
				// }
			}
		} else {
			if number > length {
				if minus {
					fmt.Printf(string(data))
				} else {
					fmt.Printf("")
				}
			} else {
				if !minus {
					fmt.Print(string(data[number-1 : length]))
				} else {
					fmt.Print(string(data[length-number : length]))
				}
			}
			if index != len(arr)-1 {
				// z01.PrintRune('\n')
				// if !errorlar {
				// 	// z01.PrintRune('\n')
				// }
			}
		}
	}
}

func IsValid(arguments []string) (bool, uint64, []string, bool) {
	var str []string
	var result uint64
	var positive bool
	if len(arguments) == 0 {
		return false, result, str, positive
	}
	for index := 0; index < len(arguments); index++ {
		if arguments[index] == "-c" && index != len(arguments)-1 {
			// fmt.Println(arguments[index+1])
			if arguments[index+1] == "0" {
				index = index + 1
				result = 0
			} else if len(arguments[index+1]) < len("18446744073709551615") {
				result, positive = functions.AtoiUnix(arguments[index+1])
				index = index + 1
			} else if len(arguments[index+1]) == len("18446744073709551615") && arguments[index+1] <= "18446744073709551615" {
				result, positive = functions.AtoiUnix(arguments[index+1])
				index = index + 1
			} else {
				return false, result, str, positive
			}
		} else if arguments[index] == "-c" && index == len(arguments)-1 {
			return false, result, str, positive
		} else {
			str = append(str, arguments[index])
		}
	}
	return true, result, str, positive
}
