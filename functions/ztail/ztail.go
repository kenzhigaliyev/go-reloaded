package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"student/functions"
)

func main() {
	arguments := os.Args[1:]
	arg, number, arr := IsValid(arguments)
	if !arg {
		os.Stdout.ReadFrom(os.Stdin)
	} else {
		for _, value := range arr {
			info, err := ioutil.ReadFile(value)
			if err != nil {
				fmt.Println("ERROR:" + value)
			}
			length := len(info)
			fmt.Println(string(info)[length-number:])
		}
	}
}

func IsValid(arguments []string) (bool, int, []string) {
	var str []string
	var result int
	if len(arguments) == 0 {
		return false, result, str
	}
	for index := 0; index < len(arguments); index++ {
		// fmt.Println(index)
		if arguments[index] == "-c" && index != len(arguments) {
			if arguments[index+1] == "0" {
				index = index + 1
				result = 0
			} else if functions.Atoi(arguments[index+1]) != 0 {
				result = functions.Atoi(arguments[index+1])
				index = index + 1
			} else {
				return false, result, str
			}
		} else {
			str = append(str, arguments[index])
		}
	}

	// for index, val := range arguments {
	// 	fmt.Println(index)
	// 	if val == "-c" && index != len(arguments) {
	// 		if arguments[index+1] == "0" {
	// 			index = index + 2
	// 			result = 0
	// 		} else if functions.Atoi(arguments[index+1]) != 0 {
	// 			result = functions.Atoi(arguments[index+1])
	// 			index = index + 2
	// 		} else {
	// 			return false, result, str
	// 		}
	// 	} else {
	// 		str = append(str, val)
	// 	}
	// }
	// fmt.Println(result, str)
	return true, result, str
}
