package main

import (
	"fmt"
	"io"
	"os"
	"student/functions"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	var new int
	arg, number, arr, positive := IsValid(arguments)
	positive = positive
	if arg == false {
		fmt.Printf("tail: option requires an argument -- 'c'\nTry 'tail --help' for more information.\n")
		return
	}
	// fmt.Println(number)
	number = number
	for index, value := range arr {
		file, err := os.Open(value)
		if err != nil {
			fmt.Printf("tail: cannot open '%s' for reading: No such file or directory", value)
			z01.PrintRune('\n')
			os.Exit(1)
		}
		info, err := file.Stat()
		if err != nil {

		}
		length := uint64(info.Size())
		fmt.Println(length)
		defer file.Close()

		data := make([]byte, length)

		for {
			n, err := file.Read(data)
			if err == io.EOF { // если конец файла
				break // выходим из цикла
			}
			new = n
		}
		if len(arr) != 1 {
			fmt.Printf("==> %s <==", value)
			z01.PrintRune('\n')
			if number>length{
				fmt.Printf(string(data)
			} else {
				fmt.Print(data[:number])
			}
			if index != len(arr)-1 {
				z01.PrintRune('\n')
			}
		} else {
			fmt.Print(string(data[:new]))
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
