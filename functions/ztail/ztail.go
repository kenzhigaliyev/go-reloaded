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
	arg, number, arr := IsValid(arguments)
	if arg == false {
		fmt.Printf("tail: option requires an argument -- 'c'\nTry 'tail --help' for more information.\n")
		return
	}

	for index, value := range arr {
		file, err := os.Open(value)
		if err != nil {
			fmt.Printf("tail: cannot open '%s' for reading: No such file or directory", value)
			z01.PrintRune('\n')
			os.Exit(1)
		}
		defer file.Close()

		data := make([]byte, 64)

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
			fmt.Print(string(data[new-number : new]))
			if index != len(arr)-1 {
				z01.PrintRune('\n')
			}
		} else {
			fmt.Print(string(data[new-number : new]))
		}

		// for _, value := range arr {
		// 	info, err := os.Open(value)
		// 	if err != nil {
		// 		defer os.Exit(1)
		// 		fmt.Println("ztail:")
		// 		// os.Stdout.Write(err)

		// 	}
		// 	defer info.Close()
		// 	data := make([]byte, 64)
		// 	for {
		// 		values, err := info.Read(data)
		// 		if err == io.EOF {
		// 			break
		// 		}
		// 		fmt.Printf("%s", string(data[values-number:values]))
		// 		// length := len(data)
		// 		// fmt.Println(string(data)[length-number : length])
		// 	}
		// 	fmt.Printf("==> %s <==", value)
		// }
	}

	// arguments := os.Args[1:]
	// arg, number, arr := IsValid(arguments)
	// if !arg {
	// 	os.Stdout.ReadFrom(os.Stdin)
	// } else {
	// 	for _, value := range arr {
	// 		info, err := ioutil.ReadFile(value)
	// 		if err != nil {
	// 			fmt.Println("ERROR:" + value)
	// 		}
	// 		length := len(info)
	// 		fmt.Println(string(info)[length-number:])
	// 	}
	// }
}

func IsValid(arguments []string) (bool, int, []string) {
	var str []string
	var result int
	if len(arguments) == 0 {
		return false, result, str
	}
	for index := 0; index < len(arguments); index++ {
		if arguments[index] == "-c" && index != len(arguments)-1 {
			if arguments[index+1] == "0" {
				index = index + 1
				result = 0
			} else if functions.Atoi(arguments[index+1]) != 0 {
				result = functions.Atoi(arguments[index+1])
				index = index + 1
			} else {
				return false, result, str
			}
		} else if arguments[index] == "-c" && index == len(arguments)-1 {
			return false, result, str
		} else {
			str = append(str, arguments[index])
		}
	}

	return true, result, str
}
