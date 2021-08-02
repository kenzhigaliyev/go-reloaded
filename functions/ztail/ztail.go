package main

import (
	"fmt"
	"os"
	"student/functions"

	"github.com/01-edu/z01"
)

var errorlar = true
var filedar = true
var ci = false

func main() {
	arguments := os.Args[1:]
	arg, number, arr, minus, valid, big := IsValid(arguments)
	if !valid || big {
		return
	}
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
			if err != nil { // если конец файла
				break // выходим из цикла
			}
		}
		if len(arr) != 1 {
			if !errorlar && !filedar {
				z01.PrintRune('\n')
			}
			filedar = false
			fmt.Printf("==> %s <==", value)
			z01.PrintRune('\n')
			// fmt.Println(ci)
			if !ci {
				fmt.Printf(string(data))
				if index != len(arr)-1 {
					z01.PrintRune('\n')
				}
			} else {
				if number > length {
					if minus {
						fmt.Printf(string(data))
					} else {
						fmt.Printf("")
					}
				} else if number != uint64(0) {
					if !minus {
						fmt.Printf(string(data[number-1 : length]))
					} else {
						fmt.Printf(string(data[length-number : length]))
					}
				} else {
					if !minus {
						continue
					} else {
						fmt.Printf(string(data))
					}
				}
			}
		} else {
			if number > length {
				if minus {
					fmt.Printf(string(data))
				} else {
					fmt.Printf("")
				}
			} else if number != uint64(0) {
				if !minus {
					fmt.Printf(string(data[number-1 : length]))
				} else {
					fmt.Printf(string(data[length-number : length]))
				}
			} else {
				if !minus {
					fmt.Printf(string(data))
				} else {
					continue
				}
			}
		}
	}
}

func IsValid(arguments []string) (bool, uint64, []string, bool, bool, bool) {
	var str []string
	var result uint64
	var positive bool
	var valid = true
	var big = false
	if len(arguments) == 0 {
		return false, result, str, positive, valid, big
	}
	for index := 0; index < len(arguments); index++ {
		if arguments[index] == "-c" && index != len(arguments)-1 {
			ci = true
			if len(arguments[index+1]) < len("18446744073709551615") && len(arguments[index+1]) > 0 {
				result, positive, valid = functions.AtoiUnix(arguments[index+1])
				index = index + 1
			} else if len(arguments[index+1]) == len("18446744073709551615") && arguments[index+1] <= "18446744073709551615" {
				result, positive, valid = functions.AtoiUnix(arguments[index+1])
				index = index + 1
			} else if arguments[index+1] == "" {
				fmt.Printf("tail: invalid number of bytes: ‘’\n")
				big = true
				return false, result, str, positive, valid, big
			} else {
				fmt.Printf("tail: invalid number of bytes: ‘%s’: Value too large for defined data type\n", arguments[index+1])
				big = true
				return false, result, str, positive, valid, big
			}
			if !valid {
				fmt.Printf("tail: invalid number of bytes: ‘%s’\n", arguments[index])
				return false, result, str, positive, valid, big
			}
		} else if arguments[index] == "-c" && index == len(arguments)-1 {
			return false, result, str, positive, valid, big
		} else {
			str = append(str, arguments[index])
		}
	}
	return true, result, str, positive, valid, big
}
